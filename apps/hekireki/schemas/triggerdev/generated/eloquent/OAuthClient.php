<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

class OAuthClient extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'OAuthClient';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'kind',
        'clientId',
        'metadataUrl',
        'displayName',
        'selfReportedName',
        'redirectUris',
        'dedupeKey',
        'metadataExpiresAt',
        'lastUsedAt',
    ];

    protected $casts = [
        'kind' => OAuthClientKind::class,
        'redirectUris' => AsPrismaList::class,
        'metadataExpiresAt' => 'datetime',
        'lastUsedAt' => 'datetime',
    ];

    public function authorizationCodes(): HasMany
    {
        return $this->hasMany(AuthorizationCode::class, 'oauthClientId');
    }

    public function personalAccessTokens(): HasMany
    {
        return $this->hasMany(PersonalAccessToken::class, 'oauthClientId');
    }

    public function refreshTokens(): HasMany
    {
        return $this->hasMany(OAuthRefreshToken::class, 'oauthClientId');
    }
}