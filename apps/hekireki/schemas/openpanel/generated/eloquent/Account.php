<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Account extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'accounts';

    protected $keyType = 'string';

    protected $fillable = [
        'id',
        'userId',
        'email',
        'provider',
        'providerId',
        'accountId',
        'idToken',
        'accessToken',
        'refreshToken',
        'accessTokenExpiresAt',
        'refreshTokenExpiresAt',
        'scope',
        'password',
    ];

    protected $casts = [
        'accessTokenExpiresAt' => 'datetime',
        'refreshTokenExpiresAt' => 'datetime',
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

    public function resetPasswords(): HasMany
    {
        return $this->hasMany(ResetPassword::class, 'accountId');
    }
}