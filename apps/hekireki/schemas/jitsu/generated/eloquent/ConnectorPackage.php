<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

/**
 * @property \Carbon\CarbonInterface $createdAt @zod.custom(z.coerce.date())
 * @property \Carbon\CarbonInterface $updatedAt @zod.custom(z.coerce.date())
 */
class ConnectorPackage extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'ConnectorPackage';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'packageId',
        'packageType',
        'meta',
        'logoSvg',
    ];

    protected $attributes = [
        'packageType' => 'airbyte',
    ];

    protected $casts = [
        'meta' => 'array',
        'logoSvg' => AsBytes::class,
    ];
}