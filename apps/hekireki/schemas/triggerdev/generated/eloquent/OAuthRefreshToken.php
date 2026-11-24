<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class OAuthRefreshToken extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = null;

    protected $table = 'OAuthRefreshToken';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'hashedToken',
        'personalAccessTokenId',
        'oauthClientId',
        'scope',
        'resource',
        'usedAt',
        'replacedById',
        'idleExpiresAt',
    ];

    protected $casts = [
        'scope' => AsPrismaList::class,
        'usedAt' => 'datetime',
        'idleExpiresAt' => 'datetime',
    ];

    public function personalAccessToken(): BelongsTo
    {
        return $this->belongsTo(PersonalAccessToken::class, 'personalAccessTokenId');
    }

    public function oauthClient(): BelongsTo
    {
        return $this->belongsTo(OAuthClient::class, 'oauthClientId');
    }
}