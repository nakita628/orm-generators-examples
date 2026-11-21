<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property \Carbon\CarbonInterface $createdAt @zod.custom(z.coerce.date())
 * @property \Carbon\CarbonInterface $updatedAt @zod.custom(z.coerce.date())
 */
class ProfileBuilder extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'ProfileBuilder';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'name',
        'deleted',
        'workspaceId',
        'intermediateStorageCredentials',
        'destinationId',
        'connectionOptions',
        'version',
    ];

    protected $attributes = [
        'deleted' => false,
        'version' => 0,
    ];

    protected $casts = [
        'deleted' => 'boolean',
        'intermediateStorageCredentials' => 'array',
        'connectionOptions' => 'array',
        'version' => 'integer',
    ];

    public function workspace(): BelongsTo
    {
        return $this->belongsTo(Workspace::class, 'workspaceId');
    }

    public function destination(): BelongsTo
    {
        return $this->belongsTo(ConfigurationObject::class, 'destinationId');
    }

    public function functions(): HasMany
    {
        return $this->hasMany(ProfileBuilderFunction::class, 'profileBuilderId');
    }
}