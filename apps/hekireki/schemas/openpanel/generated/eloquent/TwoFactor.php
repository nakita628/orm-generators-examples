<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class TwoFactor extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'two_factor';

    protected $keyType = 'string';

    protected $fillable = [
        'id',
        'userId',
        'secret',
        'backupCodes',
        'verified',
        'failedVerificationCount',
        'lockedUntil',
    ];

    protected $attributes = [
        'verified' => true,
        'failedVerificationCount' => 0,
    ];

    protected $casts = [
        'verified' => 'boolean',
        'failedVerificationCount' => 'integer',
        'lockedUntil' => 'datetime',
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

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'userId');
    }
}