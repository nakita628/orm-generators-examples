<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;

/**
 * Used by the RunEngine during TaskRun execution
 * It has the required information to transactionally progress a run through states,
 * and prevent side effects like heartbeats failing a run that has progressed.
 * It is optimised for performance and is designed to be cleared at some point,
 * so there are no cascading relationships to other models.
 *
 * @property RunEngineVersion $engine This should always be 2+ (V1 didn't use the run engine or snapshots)
 * @property TaskRunExecutionStatus $executionStatus The execution status
 * @property string $description For debugging
 * @property bool $isValid We store invalid snapshots as a record of the run state when we tried to move
 * @property string|null $previousSnapshotId The previous snapshot ID
 * @property string $runId Run
 * @property int|null $attemptNumber This is the current run attempt number. Users can define how many attempts they want for a run.
 * @property string $environmentId Environment
 * @property mixed $completedWaitpointOrder An array of waitpoint IDs in the correct order, used for batches
 * @property string|null $checkpointId Checkpoint
 * @property string|null $workerId Worker
 * @property mixed $metadata Metadata used by various systems in the run engine
 */
class TaskRunExecutionSnapshot extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'TaskRunExecutionSnapshot';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'engine',
        'executionStatus',
        'description',
        'isValid',
        'error',
        'previousSnapshotId',
        'runId',
        'runStatus',
        'batchId',
        'attemptNumber',
        'environmentId',
        'environmentType',
        'projectId',
        'organizationId',
        'completedWaitpointOrder',
        'checkpointId',
        'workerId',
        'runnerId',
        'lastHeartbeatAt',
        'metadata',
    ];

    protected $attributes = [
        'engine' => 'V2',
        'isValid' => true,
    ];

    protected $casts = [
        'engine' => RunEngineVersion::class,
        'executionStatus' => TaskRunExecutionStatus::class,
        'isValid' => 'boolean',
        'runStatus' => TaskRunStatus::class,
        'attemptNumber' => 'integer',
        'environmentType' => RuntimeEnvironmentType::class,
        'lastHeartbeatAt' => 'datetime',
        'metadata' => 'array',
    ];

    public function run(): BelongsTo
    {
        return $this->belongsTo(TaskRun::class, 'runId');
    }

    public function checkpoint(): BelongsTo
    {
        return $this->belongsTo(TaskRunCheckpoint::class, 'checkpointId');
    }

    /** Waitpoints that have been completed for this execution */
    public function completedWaitpoints(): BelongsToMany
    {
        return $this->belongsToMany(Waitpoint::class, '_completedWaitpoints', 'A', 'B');
    }
}