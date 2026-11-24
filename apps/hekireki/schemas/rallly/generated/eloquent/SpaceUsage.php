<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * Monthly counters for metered actions charged to a space, independent of
 * the records they produce, so deleting a poll does not give back what it
 * spent.
 */
class SpaceUsage extends Model
{
    use PrismaDates;

    const CREATED_AT = null;
    const KEY_COLUMNS = ['space_id', 'resource', 'period'];

    protected $table = 'space_usage';

    protected $primaryKey = null;

    public $incrementing = false;

    protected $fillable = [
        'space_id',
        'resource',
        'period',
        'count',
        'near_limit_warned_at',
        'at_limit_warned_at',
    ];

    protected $attributes = [
        'count' => 0,
    ];

    protected $casts = [
        'resource' => SpaceUsageResource::class,
        'period' => AsPrismaDate::class,
        'count' => 'integer',
        'near_limit_warned_at' => 'datetime',
        'at_limit_warned_at' => 'datetime',
    ];

    public function getKey()
    {
        $key = [];
        foreach (static::KEY_COLUMNS as $column) {
            $key[$column] = $this->getAttribute($column);
        }

        return $key;
    }

    public function newEloquentBuilder($query)
    {
        return new CompositeKeyBuilder($query);
    }

    public function newCollection(array $models = [])
    {
        return new CompositeKeyCollection($models);
    }

    public static function destroy($ids)
    {
        $count = 0;
        foreach ((new static())->newQuery()->whereKey(func_num_args() > 1 ? func_get_args() : $ids)->get() as $model) {
            if ($model->delete()) {
                $count++;
            }
        }

        return $count;
    }

    protected function setKeysForSaveQuery($query)
    {
        foreach (static::KEY_COLUMNS as $column) {
            $query->where($column, '=', $this->original[$column] ?? $this->getAttribute($column));
        }

        return $query;
    }

    protected function setKeysForSelectQuery($query)
    {
        foreach (static::KEY_COLUMNS as $column) {
            $query->where($column, '=', $this->original[$column] ?? $this->getAttribute($column));
        }

        return $query;
    }

    public function delete()
    {
        $this->mergeAttributesFromCachedCasts();

        if (! $this->exists) {
            return null;
        }

        if ($this->fireModelEvent('deleting') === false) {
            return false;
        }

        $this->touchOwners();
        $this->performDeleteOnModel();
        $this->fireModelEvent('deleted', false);

        return true;
    }

    public function space(): BelongsTo
    {
        return $this->belongsTo(Space::class, 'space_id');
    }
}