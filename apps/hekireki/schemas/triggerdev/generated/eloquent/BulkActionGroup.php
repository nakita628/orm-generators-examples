<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * Bulk actions, like canceling and replaying runs
 *
 * @property BulkActionStatus $status When the group is created it's pending. After we've processed all the items it's completed. This does not mean the associated runs are completed.
 * @property string|null $queryName The query that will be executed to get the runs to process (if null, we are passing run ids directly)
 * @property mixed $params The params that will be passed to the query
 * @property string|null $dedupeKey Billing-limit dedupe key (grace hitAt or resolve job key). Null for dashboard bulk actions.
 * @property mixed $cursor The cursor that will be passed to the query (null for the first page)
 * @property int $successCount The number of runs that have been processed successfully
 * @property int $failureCount The number of runs that have been failed
 * @property int $totalCount The total number of runs that will be processed
 * @property string|null $name The name of the bulk action
 * @property \Carbon\CarbonInterface|null $completedAt The time the bulk action was completed
 */
class BulkActionGroup extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'BulkActionGroup';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'friendlyId',
        'projectId',
        'environmentId',
        'type',
        'status',
        'queryName',
        'params',
        'dedupeKey',
        'cursor',
        'successCount',
        'failureCount',
        'totalCount',
        'completionNotification',
        'userId',
        'name',
        'completedAt',
    ];

    protected $attributes = [
        'status' => 'PENDING',
        'successCount' => 0,
        'failureCount' => 0,
        'totalCount' => 0,
        'completionNotification' => 'NONE',
    ];

    protected $casts = [
        'type' => BulkActionType::class,
        'status' => BulkActionStatus::class,
        'params' => 'array',
        'cursor' => 'array',
        'successCount' => 'integer',
        'failureCount' => 'integer',
        'totalCount' => 'integer',
        'completionNotification' => BulkActionNotificationType::class,
        'completedAt' => 'datetime',
    ];

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    /** If this is set then it's a V2 Bulk Action that supports queries */
    public function environment(): BelongsTo
    {
        return $this->belongsTo(RuntimeEnvironment::class, 'environmentId');
    }

    /** The userId who did the bulk action */
    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'userId');
    }
}