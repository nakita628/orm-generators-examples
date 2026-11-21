<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property \Carbon\CarbonInterface $createdAt @zod.custom(z.coerce.date())
 * @property \Carbon\CarbonInterface $updatedAt @zod.custom(z.coerce.date())
 */
class WorkspaceOptions extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'WorkspaceOptions';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'workspaceId',
        'namespace',
        'value',
    ];

    protected $casts = [
        'value' => 'array',
    ];

    public function workspace(): BelongsTo
    {
        return $this->belongsTo(Workspace::class, 'workspaceId');
    }
}