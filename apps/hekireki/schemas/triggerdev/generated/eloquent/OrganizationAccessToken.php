<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string $name User-provided name for the token
 * @property OrganizationAccessTokenType $type Used to differentiate between user-generated and system-generated tokens
 * @property string $hashedToken This is used to find the token in the database
 * @property \Carbon\CarbonInterface|null $expiresAt Optional expiration date for the token
 */
class OrganizationAccessToken extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'OrganizationAccessToken';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'name',
        'type',
        'hashedToken',
        'organizationId',
        'expiresAt',
        'revokedAt',
        'lastAccessedAt',
    ];

    protected $attributes = [
        'type' => 'USER',
    ];

    protected $casts = [
        'type' => OrganizationAccessTokenType::class,
        'expiresAt' => 'datetime',
        'revokedAt' => 'datetime',
        'lastAccessedAt' => 'datetime',
    ];

    public function organization(): BelongsTo
    {
        return $this->belongsTo(Organization::class, 'organizationId');
    }
}