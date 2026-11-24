<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * A Waitpoint blocks a run from continuing until it's completed
 * If there's a waitpoint blocking a run, it shouldn't be in the queue
 *
 * @property string $idempotencyKey If it's an Event type waitpoint, this is the event. It can also be provided for the DATETIME type
 * @property bool $userProvidedIdempotencyKey If this is true then we can show it in the dashboard/return it from the SDK
 * @property \Carbon\CarbonInterface|null $idempotencyKeyExpiresAt If there's a user provided idempotency key, this is the time it expires at
 * @property string|null $inactiveIdempotencyKey If an idempotencyKey is no longer active, we store it here and generate a new one for the idempotencyKey field.
 * Clearing an idempotencyKey is useful for debounce or cancelling child runs.
 * This is a workaround because Prisma doesn't support partial indexes.
 * @property string|null $completedByTaskRunId If it's a RUN type waitpoint, this is the associated run
 * @property \Carbon\CarbonInterface|null $completedAfter If it's a DATETIME type waitpoint, this is the date.
 * If it's a MANUAL waitpoint, this can be set as the `timeout`.
 * @property string|null $completedByBatchId If it's a BATCH type waitpoint, this is the associated batch
 * @property string|null $output When completed, an output can be stored here
 * @property list<string> $tags Denormized column that holds the raw tags
 * Denormalized column that holds the raw tags
 */
class Waitpoint extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'Waitpoint';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'friendlyId',
        'type',
        'status',
        'completedAt',
        'idempotencyKey',
        'userProvidedIdempotencyKey',
        'idempotencyKeyExpiresAt',
        'inactiveIdempotencyKey',
        'completedByTaskRunId',
        'completedAfter',
        'completedByBatchId',
        'output',
        'outputType',
        'outputIsError',
        'projectId',
        'environmentId',
        'tags',
    ];

    protected $attributes = [
        'status' => 'PENDING',
        'outputType' => 'application/json',
        'outputIsError' => false,
    ];

    protected $casts = [
        'type' => WaitpointType::class,
        'status' => WaitpointStatus::class,
        'completedAt' => 'datetime',
        'userProvidedIdempotencyKey' => 'boolean',
        'idempotencyKeyExpiresAt' => 'datetime',
        'completedAfter' => 'datetime',
        'outputIsError' => 'boolean',
        'tags' => AsPrismaList::class,
    ];

    public function completedByTaskRun(): BelongsTo
    {
        return $this->belongsTo(TaskRun::class, 'completedByTaskRunId');
    }

    public function completedByBatch(): BelongsTo
    {
        return $this->belongsTo(BatchTaskRun::class, 'completedByBatchId');
    }

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function environment(): BelongsTo
    {
        return $this->belongsTo(RuntimeEnvironment::class, 'environmentId');
    }

    /** The runs this waitpoint is blocking */
    public function blockingTaskRuns(): HasMany
    {
        return $this->hasMany(TaskRunWaitpoint::class, 'waitpointId');
    }

    /** All runs that have ever been blocked by this waitpoint, used for display purposes */
    public function connectedRuns(): BelongsToMany
    {
        return $this->belongsToMany(TaskRun::class, '_WaitpointRunConnections', 'B', 'A');
    }

    /** When a waitpoint is complete */
    public function completedExecutionSnapshots(): BelongsToMany
    {
        return $this->belongsToMany(TaskRunExecutionSnapshot::class, '_completedWaitpoints', 'B', 'A');
    }
}