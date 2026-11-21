<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property \Carbon\CarbonInterface $createdAt @zod.custom(z.coerce.date())
 * @property \Carbon\CarbonInterface $updatedAt @zod.custom(z.coerce.date())
 * @property \Carbon\CarbonInterface|null $lastUsed @zod.custom(z.coerce.date())
 * @property \Carbon\CarbonInterface|null $expiresAt @zod.custom(z.coerce.date())
 */
class UserApiToken extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'UserApiToken';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'lastUsed',
        'hint',
        'hash',
        'userId',
        'type',
        'name',
        'expiresAt',
        'workspaceId',
        'role',
        'oauthClientId',
    ];

    protected $casts = [
        'lastUsed' => 'datetime',
        'expiresAt' => 'datetime',
    ];

    public function user(): BelongsTo
    {
        return $this->belongsTo(UserProfile::class, 'userId');
    }

    public function oauthClient(): BelongsTo
    {
        return $this->belongsTo(OAuthClient::class, 'oauthClientId');
    }

    public function accessTokens(): HasMany
    {
        return $this->hasMany(OAuthAccessToken::class, 'refreshTokenId');
    }
}