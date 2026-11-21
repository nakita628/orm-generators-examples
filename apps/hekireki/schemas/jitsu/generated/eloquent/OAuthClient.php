<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property \Carbon\CarbonInterface $createdAt @zod.custom(z.coerce.date())
 * @property \Carbon\CarbonInterface $updatedAt @zod.custom(z.coerce.date())
 */
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
        'clientSecretHash',
        'name',
        'redirectUris',
    ];

    public function tokens(): HasMany
    {
        return $this->hasMany(UserApiToken::class, 'oauthClientId');
    }
}