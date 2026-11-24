<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;
use Illuminate\Database\Eloquent\Relations\HasMany;
use Illuminate\Database\Eloquent\Relations\HasOne;

/**
 * @property mixed $idempotencyKeyOptions Stores the user-provided key and scope: { key: string, scope: "run" | "attempt" | "global" }
 * @property mixed $debounce Debounce options: { key: string, delay: string, createdAt: Date }
 * @property string $masterQueue The main queue that this run is part of
 * @property string|null $region User-facing geo region, stamped at trigger; workerQueue is where it actually ran.
 * @property string|null $secondaryMasterQueue @deprecated
 * @property int|null $attemptNumber From engine v2+ this will be defined after a run has been dequeued (starting at 1)
 * @property list<string> $runTags Denormized column that holds the raw tags
 * @property string|null $taskVersion Denormalized version of the background worker task
 * @property \Carbon\CarbonInterface|null $startedAt startedAt marks the point at which a run is dequeued from MarQS
 * @property \Carbon\CarbonInterface|null $executedAt executedAt is set when the first attempt is about to execute
 * @property int $priorityMs The "priority" of the run. This is just a negative offset in ms for the queue timestamp
 * E.g. a value of 60_000 would put the run into the queue 60s ago.
 * @property mixed $gates Gates for this run: other queues it must also hold a concurrency slot in while
 * executing. Shape: [{ queue: string, concurrencyKey?: string }]
 * @property string|null $oneTimeUseToken optional token that can be used to authenticate the task run
 * @property string $taskEventStore Where the logs are stored
 * @property bool $resumeParentOnCompletion whether or not the task run was created because of a triggerAndWait for batchTriggerAndWait
 * @property int $depth The depth of this task run in the task run hierarchy
 * @property string|null $parentSpanId The span ID of the "trigger" span in the parent task run
 * @property mixed $runChainState Holds the state of the run chain for deadlock detection
 * @property string|null $seedMetadata seed run metadata
 * @property string|null $metadata Run metadata
 * @property mixed $annotations Structured annotations: triggerSource, triggerAction, rootTriggerSource, rootScheduleId
 * @property bool|null $isWarmStart Whether the latest attempt was a warm start. Null until first attempt starts.
 * @property string|null $output Run output
 * @property mixed $error Run error
 * @property string|null $planType Organization's billing plan type (cached for fallback when billing API fails)
 * @property string $realtimeStreamsVersion The version of the realtime streams implementation used by the run
 * @property list<string> $realtimeStreams Store the stream keys that are being used by the run
 * @property string|null $streamBasinName S2 basin where this run's realtime streams live. Stamped at create
 * time from `Organization.streamBasinName` so reads can resolve the
 * basin without joining org. Null when the org has no per-org basin
 * (OSS, or pre-backfill); reads fall back to the global basin.
 */
class TaskRun extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'TaskRun';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'number',
        'friendlyId',
        'engine',
        'status',
        'statusReason',
        'idempotencyKey',
        'idempotencyKeyExpiresAt',
        'idempotencyKeyOptions',
        'debounce',
        'taskIdentifier',
        'isTest',
        'payload',
        'payloadType',
        'context',
        'traceContext',
        'traceId',
        'spanId',
        'runtimeEnvironmentId',
        'environmentType',
        'projectId',
        'organizationId',
        'queue',
        'lockedQueueId',
        'masterQueue',
        'region',
        'secondaryMasterQueue',
        'attemptNumber',
        'runTags',
        'taskVersion',
        'sdkVersion',
        'cliVersion',
        'startedAt',
        'executedAt',
        'completedAt',
        'machinePreset',
        'usageDurationMs',
        'costInCents',
        'baseCostInCents',
        'lockedAt',
        'lockedById',
        'lockedToVersionId',
        'priorityMs',
        'concurrencyKey',
        'gates',
        'delayUntil',
        'queuedAt',
        'ttl',
        'expiredAt',
        'maxAttempts',
        'lockedRetryConfig',
        'oneTimeUseToken',
        'taskEventStore',
        'queueTimestamp',
        'scheduleInstanceId',
        'scheduleId',
        'bulkActionGroupIds',
        'logsDeletedAt',
        'replayedFromTaskRunFriendlyId',
        'rootTaskRunId',
        'parentTaskRunId',
        'parentTaskRunAttemptId',
        'batchId',
        'resumeParentOnCompletion',
        'depth',
        'parentSpanId',
        'runChainState',
        'seedMetadata',
        'seedMetadataType',
        'metadata',
        'metadataType',
        'metadataVersion',
        'annotations',
        'isWarmStart',
        'output',
        'outputType',
        'error',
        'planType',
        'maxDurationInSeconds',
        'realtimeStreamsVersion',
        'realtimeStreams',
        'streamBasinName',
    ];

    protected $attributes = [
        'number' => 0,
        'engine' => 'V1',
        'status' => 'PENDING',
        'isTest' => false,
        'payloadType' => 'application/json',
        'masterQueue' => 'main',
        'usageDurationMs' => 0,
        'costInCents' => 0.0,
        'baseCostInCents' => 0.0,
        'priorityMs' => 0,
        'taskEventStore' => 'taskEvent',
        'resumeParentOnCompletion' => false,
        'depth' => 0,
        'seedMetadataType' => 'application/json',
        'metadataType' => 'application/json',
        'metadataVersion' => 1,
        'outputType' => 'application/json',
        'realtimeStreamsVersion' => 'v1',
    ];

    protected $casts = [
        'number' => 'integer',
        'engine' => RunEngineVersion::class,
        'status' => TaskRunStatus::class,
        'idempotencyKeyExpiresAt' => 'datetime',
        'idempotencyKeyOptions' => 'array',
        'debounce' => 'array',
        'isTest' => 'boolean',
        'context' => 'array',
        'traceContext' => 'array',
        'environmentType' => RuntimeEnvironmentType::class,
        'attemptNumber' => 'integer',
        'runTags' => AsPrismaList::class,
        'startedAt' => 'datetime',
        'executedAt' => 'datetime',
        'completedAt' => 'datetime',
        'usageDurationMs' => 'integer',
        'costInCents' => 'float',
        'baseCostInCents' => 'float',
        'lockedAt' => 'datetime',
        'priorityMs' => 'integer',
        'gates' => 'array',
        'delayUntil' => 'datetime',
        'queuedAt' => 'datetime',
        'expiredAt' => 'datetime',
        'maxAttempts' => 'integer',
        'lockedRetryConfig' => 'array',
        'queueTimestamp' => 'datetime',
        'bulkActionGroupIds' => AsPrismaList::class,
        'logsDeletedAt' => 'datetime',
        'resumeParentOnCompletion' => 'boolean',
        'depth' => 'integer',
        'runChainState' => 'array',
        'metadataVersion' => 'integer',
        'annotations' => 'array',
        'isWarmStart' => 'boolean',
        'error' => 'array',
        'maxDurationInSeconds' => 'integer',
        'realtimeStreams' => AsPrismaList::class,
    ];

    public function runtimeEnvironment(): BelongsTo
    {
        return $this->belongsTo(RuntimeEnvironment::class, 'runtimeEnvironmentId');
    }

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function lockedBy(): BelongsTo
    {
        return $this->belongsTo(BackgroundWorkerTask::class, 'lockedById');
    }

    public function lockedToVersion(): BelongsTo
    {
        return $this->belongsTo(BackgroundWorker::class, 'lockedToVersionId');
    }

    /** This represents the original task that that was triggered outside of a Trigger.dev task */
    public function rootTaskRun(): BelongsTo
    {
        return $this->belongsTo(TaskRun::class, 'rootTaskRunId');
    }

    /** The immediate parent run of this task run */
    public function parentTaskRun(): BelongsTo
    {
        return $this->belongsTo(TaskRun::class, 'parentTaskRunId');
    }

    /** The immediate parent attempt of this task run */
    public function parentTaskRunAttempt(): BelongsTo
    {
        return $this->belongsTo(TaskRunAttempt::class, 'parentTaskRunAttemptId');
    }

    /** The batch run that this task run is a part of */
    public function batch(): BelongsTo
    {
        return $this->belongsTo(BatchTaskRun::class, 'batchId');
    }

    /** When this run is finished, the waitpoint will be marked as completed */
    public function associatedWaitpoint(): HasOne
    {
        return $this->hasOne(Waitpoint::class, 'completedByTaskRunId');
    }

    public function dependency(): HasOne
    {
        return $this->hasOne(TaskRunDependency::class, 'taskRunId');
    }

    public function attempts(): HasMany
    {
        return $this->hasMany(TaskRunAttempt::class, 'taskRunId');
    }

    public function checkpoints(): HasMany
    {
        return $this->hasMany(Checkpoint::class, 'runId');
    }

    /** If there are any blocked waitpoints, the run won't be executed */
    public function blockedByWaitpoints(): HasMany
    {
        return $this->hasMany(TaskRunWaitpoint::class, 'taskRunId');
    }

    public function batchItems(): HasMany
    {
        return $this->hasMany(BatchTaskRunItem::class, 'taskRunId');
    }

    public function CheckpointRestoreEvent(): HasMany
    {
        return $this->hasMany(CheckpointRestoreEvent::class, 'runId');
    }

    public function executionSnapshots(): HasMany
    {
        return $this->hasMany(TaskRunExecutionSnapshot::class, 'runId');
    }

    public function sourceBulkActionItems(): HasMany
    {
        return $this->hasMany(BulkActionItem::class, 'sourceRunId');
    }

    public function destinationBulkActionItems(): HasMany
    {
        return $this->hasMany(BulkActionItem::class, 'destinationRunId');
    }

    /** The root run will have a list of all the descendant runs, children, grand children, etc. */
    public function descendantRuns(): HasMany
    {
        return $this->hasMany(TaskRun::class, 'rootTaskRunId');
    }

    /** The immediate child runs of this task run */
    public function childRuns(): HasMany
    {
        return $this->hasMany(TaskRun::class, 'parentTaskRunId');
    }

    public function playgroundConversations(): HasMany
    {
        return $this->hasMany(PlaygroundConversation::class, 'runId');
    }

    /** All waitpoints that blocked this run at some point, used for display purposes */
    public function connectedWaitpoints(): BelongsToMany
    {
        return $this->belongsToMany(Waitpoint::class, '_WaitpointRunConnections', 'A', 'B');
    }
}