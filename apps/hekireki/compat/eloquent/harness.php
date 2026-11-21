<?php

// Runs the plan through the Eloquent models hekireki wrote, against the database Prisma created.
//
//   php harness.php crud <plan.json> <models dir> <url>
//
// Eloquent reads whatever columns a row has, so every column "maps"; what is held to the plan is
// each model's table, its key, its attribute defaults and casts (what a value reads back as),
// and its relations.

// Warnings go to stderr: stdout is the report.
ini_set('display_errors', 'stderr');

require '/harness/vendor/autoload.php';

use Illuminate\Database\Capsule\Manager as Capsule;
use Illuminate\Database\Eloquent\Collection;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\Relation;

[, $mode, $planPath, $modelsDir, $url] = $argv;
if ($mode !== 'crud') {
    throw new RuntimeException('crud only');
}
$plan = json_decode(file_get_contents($planPath), true);

$capsule = new Capsule();
$capsule->addConnection(connection($plan['dialect'], $url));
$capsule->setAsGlobal();
$capsule->bootEloquent();
date_default_timezone_set('UTC');

// App\Models\X lives in X.php, as Composer's PSR-4 autoload would find it in an application.
spl_autoload_register(function (string $class) use ($modelsDir) {
    if (str_starts_with($class, 'App\\Models\\')) {
        $file = $modelsDir . '/' . substr($class, strlen('App\\Models\\')) . '.php';
        if (is_file($file)) {
            require_once $file;
        }
    }
});
$files = glob($modelsDir . '/*.php');
$classes = [];
foreach ($files as $file) {
    require_once $file;
}
foreach (get_declared_classes() as $class) {
    if (str_starts_with($class, 'App\\Models\\') && is_subclass_of($class, Model::class)) {
        $classes[(new $class())->getTable()] = $class;
    }
}

$out = ['orm' => 'eloquent', 'models' => []];
$rows = [];
$state = [];

foreach ($plan['models'] as $m) {
    $r = ['model' => $m['name']];
    $class = $classes[$m['table']] ?? null;
    if ($class === null) {
        $r['missing'] = true;
        $out['models'][] = $r;
        continue;
    }
    $r['unmapped'] = [];
    if (! $m['insertable']) {
        $r['skipped'] = 'a required key closes a cycle';
        $out['models'][] = $r;
        continue;
    }
    $model = new $class();
    try {
        foreach ($m['fields'] as $f) {
            if ($f['write'] === 'value') {
                $model->setAttribute($f['column'], toPhp($f, $f['value']));
            } elseif ($f['write'] === 'parent') {
                $parent = $f['fk']['available'] ? ($rows[$f['fk']['model']] ?? null) : null;
                $model->setAttribute($f['column'], $parent ? toPhp($f, $parent[$f['fk']['column']]) : null);
            }
        }
        $model->save();
        $r['insert'] = ['ok' => true];
    } catch (Throwable $e) {
        $r['insert'] = ['ok' => false, 'error' => message($e)];
        $out['models'][] = $r;
        continue;
    }
    $key = [];
    foreach ($m['pk'] as $col) {
        $key[$col] = $model->getAttributes()[$col] ?? $model->getAttribute($col);
    }
    $r['pk'] = array_map(fn ($v) => normal($v), $key);
    try {
        $loaded = find($class, $key);
        if ($loaded === null) {
            throw new RuntimeException('no row came back for the key the insert returned');
        }
        $values = [];
        foreach ($m['fields'] as $f) {
            $values[$f['column']] = normal($loaded->getAttribute($f['column']), $f);
        }
        $r['read'] = ['ok' => true, 'values' => $values];
        $rows[$m['name']] = $values;
    } catch (Throwable $e) {
        $r['read'] = ['ok' => false, 'error' => message($e)];
        $out['models'][] = $r;
        continue;
    }
    if ($m['update']) {
        try {
            $col = $m['update']['column'];
            $row = find($class, $key);
            $row->setAttribute($col, $m['update']['value']);
            $row->save();
            $f = current(array_filter($m['fields'], fn ($x) => $x['column'] === $col));
            $r['update'] = ['ok' => true, 'value' => normal(find($class, $key)->getAttribute($col), $f)];
        } catch (Throwable $e) {
            $r['update'] = ['ok' => false, 'error' => message($e)];
        }
    }
    $state[$m['name']] = [$class, $key];
    $out['models'][] = $r;
}

$index = [];
foreach ($out['models'] as $i => $r) {
    $index[$r['model']] = $i;
}

foreach ($plan['models'] as $m) {
    if (! isset($state[$m['name']])) {
        continue;
    }
    [$class, $key] = $state[$m['name']];
    $relations = [];
    foreach ($m['relations'] as $rel) {
        $result = ['name' => $rel['name'], 'kind' => $rel['kind']];
        if (! isset($state[$rel['model']])) {
            $relations[] = $result + ['skipped' => 'no row on the other side'];
            continue;
        }
        $method = findRelation($class, $rel['name']);
        if ($method === null) {
            $relations[] = $result + ['found' => false];
            continue;
        }
        try {
            $value = find($class, $key)->{$method};
            $count = $value instanceof Collection ? $value->count() : ($value === null ? 0 : 1);
            $relations[] = $result + ['found' => true, 'loaded' => $count];
        } catch (Throwable $e) {
            $relations[] = $result + ['found' => true, 'error' => message($e)];
        }
    }
    $out['models'][$index[$m['name']]]['relations'] = $relations;
}

foreach (array_reverse($plan['models']) as $m) {
    if (! isset($state[$m['name']])) {
        continue;
    }
    [$class, $key] = $state[$m['name']];
    try {
        find($class, $key)->delete();
        $gone = find($class, $key) === null;
        $out['models'][$index[$m['name']]]['delete'] = $gone ? ['ok' => true] : ['ok' => false, 'error' => 'the row is still there'];
    } catch (Throwable $e) {
        $out['models'][$index[$m['name']]]['delete'] = ['ok' => false, 'error' => message($e)];
    }
}

echo json_encode($out, JSON_UNESCAPED_SLASHES | JSON_PRESERVE_ZERO_FRACTION);

function find(string $class, array $key): ?Model
{
    $query = $class::query();
    foreach ($key as $col => $value) {
        $query->where($col, $value);
    }
    return $query->first();
}

function findRelation(string $class, string $prismaName): ?string
{
    $squash = fn ($s) => strtolower(str_replace('_', '', $s));
    foreach (get_class_methods($class) as $method) {
        if ($squash($method) !== $squash($prismaName)) {
            continue;
        }
        $returns = (new ReflectionMethod($class, $method))->getReturnType();
        if ($returns instanceof ReflectionNamedType && is_subclass_of($returns->getName(), Relation::class)) {
            return $method;
        }
    }
    return null;
}

function toPhp(array $f, mixed $value): mixed
{
    if ($value === null) {
        return null;
    }
    if ($f['list']) {
        return array_map(fn ($v) => toPhp(['list' => false] + $f, $v), $value);
    }
    return match ($f['type']) {
        'DateTime' => new DateTimeImmutable($value),
        'Bytes' => base64_decode($value),
        default => $value,
    };
}

function normal(mixed $value, ?array $f = null): mixed
{
    if ($value === null) {
        return null;
    }
    if ($value instanceof DateTimeInterface) {
        return DateTimeImmutable::createFromInterface($value)->setTimezone(new DateTimeZone('UTC'))->format('Y-m-d\TH:i:s.v\Z');
    }
    if ($value instanceof BackedEnum) {
        return $value->value;
    }
    if ($value instanceof Illuminate\Support\Collection || $value instanceof ArrayObject) {
        $value = $value->toArray();
    }
    if (is_array($value)) {
        return array_map(fn ($v) => normal($v), $value);
    }
    if (is_string($value) && ($f['type'] ?? null) === 'Bytes') {
        return base64_encode($value);
    }
    if (is_resource($value)) {
        $value = stream_get_contents($value);
        return ($f['type'] ?? null) === 'Bytes' ? base64_encode($value) : $value;
    }
    return $value;
}

function message(Throwable $e): string
{
    return substr(get_class($e) . ': ' . $e->getMessage(), 0, 600);
}

function connection(string $dialect, string $url): array
{
    if ($dialect === 'sqlite') {
        return ['driver' => 'sqlite', 'database' => substr($url, strlen('file:')), 'foreign_key_constraints' => true];
    }
    $u = parse_url($url);
    return [
        'driver' => $dialect === 'postgresql' ? 'pgsql' : 'mysql',
        'host' => $u['host'],
        'port' => $u['port'],
        'database' => ltrim($u['path'], '/'),
        'username' => urldecode($u['user']),
        'password' => urldecode($u['pass']),
        'charset' => $dialect === 'postgresql' ? 'utf8' : 'utf8mb4',
        'collation' => $dialect === 'postgresql' ? null : 'utf8mb4_unicode_ci',
        'timezone' => '+00:00',
    ];
}
