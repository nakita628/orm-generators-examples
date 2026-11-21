<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

/**
 * @property \Carbon\CarbonInterface $createdAt @zod.custom(z.coerce.date())
 * @property \Carbon\CarbonInterface $updatedAt @zod.custom(z.coerce.date())
 */
class GlobalProps extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'GlobalProps';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'name',
        'value',
    ];

    protected $casts = [
        'value' => 'array',
    ];
}