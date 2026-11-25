<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;

class CodeMigration extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = '__code_migrations';

    protected $keyType = 'string';

    protected $fillable = [
        'id',
        'name',
    ];

    protected function insertAndSetId(Builder $query, $attributes)
    {
        $base = $query->getQuery();
        $base->applyBeforeQueryCallbacks();
        $connection = $base->getConnection();
        $connection->recordsHaveBeenModified();
        $keyName = $this->getKeyName();
        $row = $connection->selectFromWriteConnection(
            $base->getGrammar()->compileInsertGetId($base, $attributes, $keyName),
            $base->cleanBindings($attributes),
        )[0];
        $this->setAttribute($keyName, is_object($row) ? $row->{$keyName} : $row[$keyName]);
    }
}