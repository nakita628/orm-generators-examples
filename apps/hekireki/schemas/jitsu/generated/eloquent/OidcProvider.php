<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property \Carbon\CarbonInterface $createdAt @zod.custom(z.coerce.date())
 * @property \Carbon\CarbonInterface $updatedAt @zod.custom(z.coerce.date())
 */
class OidcProvider extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'OidcProvider';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'name',
        'issuer',
        'clientId',
        'clientSecret',
        'authorizationEndpoint',
        'tokenEndpoint',
        'userinfoEndpoint',
        'jwksUri',
        'introspectionEndpoint',
        'scopes',
        'audience',
        'prompt',
        'usernameClaim',
        'emailClaim',
        'nameClaim',
        'groupClaim',
        'enabled',
        'autoDiscovery',
    ];

    protected $attributes = [
        'usernameClaim' => 'email',
        'emailClaim' => 'email',
        'nameClaim' => 'name',
        'groupClaim' => 'groups',
        'enabled' => true,
        'autoDiscovery' => true,
    ];

    protected $casts = [
        'enabled' => 'boolean',
        'autoDiscovery' => 'boolean',
    ];

    public function oidcLoginGroups(): HasMany
    {
        return $this->hasMany(OidcLoginGroup::class, 'oidcProviderId');
    }
}