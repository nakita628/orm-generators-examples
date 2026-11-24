<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * Used to generate PersonalAccessTokens, they're one-time use
 *
 * @property string|null $requesterIp Client IP of the request that minted the code, compared with the approving browser's IP. Deleted with the code once it expires.
 * @property string|null $codeChallenge S256 PKCE challenge from the client that minted the code. Null for codes minted by CLIs that predate it.
 * @property int|null $callbackPort Loopback port the approving browser is redirected to with the approval code. Null means the user pastes it.
 * @property string|null $approvalCodeHash sha256 of the approval code, which only the approving browser ever sees.
 * @property string|null $requestedByUserId The user who started the OAuth authorization. Only they can approve it.
 */
class AuthorizationCode extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'AuthorizationCode';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'code',
        'personalAccessTokenId',
        'requesterIp',
        'codeChallenge',
        'callbackPort',
        'approvalCodeHash',
        'oauthClientId',
        'redirectUri',
        'scope',
        'resource',
        'approvedByUserId',
        'requestedByUserId',
    ];

    protected $casts = [
        'callbackPort' => 'integer',
        'scope' => AsPrismaList::class,
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