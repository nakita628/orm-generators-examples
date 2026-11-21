<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string|null $friendlyId @deprecated not used in new BulkActions
 * @property BulkActionItemStatus $status When the item is created it's pending. After we've processed the item it's completed. This does not mean the associated runs are completed.
 */
class BulkActionItem extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'BulkActionItem';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'friendlyId',
        'groupId',
        'type',
        'status',
        'sourceRunId',
        'destinationRunId',
        'error',
    ];

    protected $attributes = [
        'status' => 'PENDING',
    ];

    protected $casts = [
        'type' => BulkActionType::class,
        'status' => BulkActionItemStatus::class,
    ];

    /** The run that is the source of the action, e.g. when replaying this is the original run */
    public function sourceRun(): BelongsTo
    {
        return $this->belongsTo(TaskRun::class, 'sourceRunId');
    }

    /** The run that's a result of the action, this will be set when the run has been created */
    public function destinationRun(): BelongsTo
    {
        return $this->belongsTo(TaskRun::class, 'destinationRunId');
    }
}