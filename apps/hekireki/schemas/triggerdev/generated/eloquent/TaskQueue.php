<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property TaskQueueConcurrencyVersion $concurrencyVersion How the queue's concurrency was declared: V1 = legacy concurrencyLimit semantics
 * (per key when keyed, whole queue otherwise), V2 = the explicit perKey/total shape.
 * @property TaskQueueRole $role QUEUE rows are real queues; LIMIT rows back named concurrency limits and are
 * excluded from every queue-facing surface.
 * @property int|null $concurrencyLimit Represents the current concurrency limit for the queue
 * @property \Carbon\CarbonInterface|null $concurrencyLimitOverriddenAt When the concurrency limit was overridden
 * @property string|null $concurrencyLimitOverriddenBy Who overrode the concurrency limit (will be null if overridden via the API)
 * @property int|null $concurrencyLimitBase If concurrencyLimit is overridden, this is the overridden value
 * @property string|null $concurrencyLimitOverridePercent If the override was expressed as a percentage of the environment limit, this stores that
 * percentage (the source of truth). The absolute concurrencyLimit is materialized from it.
 * Decimal(5,2) allows fractional percentages like 12.50% (0.01–100.00).
 * @property int|null $totalConcurrencyLimit Caps total concurrent runs across ALL concurrencyKey values of this queue
 * (concurrencyLimit applies per key value). Null = no total cap.
 * @property \Carbon\CarbonInterface|null $totalConcurrencyLimitOverriddenAt When the total concurrency limit was overridden
 * @property string|null $totalConcurrencyLimitOverriddenBy Who overrode the total concurrency limit (null when overridden via the API)
 * @property int|null $totalConcurrencyLimitBase If totalConcurrencyLimit is overridden, the declared value it reverts to on reset
 * @property \Carbon\CarbonInterface|null $archivedAt Hidden from the dashboard queue list. Display-only: never read on trigger/dequeue paths.
 */
class TaskQueue extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'TaskQueue';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'friendlyId',
        'name',
        'type',
        'version',
        'orderableName',
        'concurrencyVersion',
        'role',
        'projectId',
        'runtimeEnvironmentId',
        'concurrencyLimit',
        'concurrencyLimitOverriddenAt',
        'concurrencyLimitOverriddenBy',
        'concurrencyLimitBase',
        'concurrencyLimitOverridePercent',
        'totalConcurrencyLimit',
        'totalConcurrencyLimitOverriddenAt',
        'totalConcurrencyLimitOverriddenBy',
        'totalConcurrencyLimitBase',
        'rateLimit',
        'paused',
        'archivedAt',
    ];

    protected $attributes = [
        'type' => 'VIRTUAL',
        'version' => 'V1',
        'concurrencyVersion' => 'V1',
        'role' => 'QUEUE',
        'paused' => false,
    ];

    protected $casts = [
        'type' => TaskQueueType::class,
        'version' => TaskQueueVersion::class,
        'concurrencyVersion' => TaskQueueConcurrencyVersion::class,
        'role' => TaskQueueRole::class,
        'concurrencyLimit' => 'integer',
        'concurrencyLimitOverriddenAt' => 'datetime',
        'concurrencyLimitBase' => 'integer',
        'concurrencyLimitOverridePercent' => AsDecimal::class,
        'totalConcurrencyLimit' => 'integer',
        'totalConcurrencyLimitOverriddenAt' => 'datetime',
        'totalConcurrencyLimitBase' => 'integer',
        'rateLimit' => 'array',
        'paused' => 'boolean',
        'archivedAt' => 'datetime',
    ];

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function runtimeEnvironment(): BelongsTo
    {
        return $this->belongsTo(RuntimeEnvironment::class, 'runtimeEnvironmentId');
    }

    public function attempts(): HasMany
    {
        return $this->hasMany(TaskRunAttempt::class, 'queueId');
    }

    public function tasks(): HasMany
    {
        return $this->hasMany(BackgroundWorkerTask::class, 'queueId');
    }

    public function workers(): BelongsToMany
    {
        return $this->belongsToMany(BackgroundWorker::class, '_BackgroundWorkerToTaskQueue', 'B', 'A');
    }
}