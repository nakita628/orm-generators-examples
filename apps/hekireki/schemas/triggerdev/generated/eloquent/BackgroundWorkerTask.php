<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property mixed $gates Gates declared on the task: other queues its runs must also hold a concurrency
 * slot in while executing. Shape: [{ queue: string, concurrencyKey?: string }]
 * @property mixed $config Extra task configuration JSON. Shape depends on triggerSource.
 * AGENT: { type: "ai-sdk-chat" }
 * @property mixed $executionPolicy RunExecutionPolicy from @trigger.dev/core. Null means the standard policy.
 */
class BackgroundWorkerTask extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'BackgroundWorkerTask';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'slug',
        'description',
        'friendlyId',
        'filePath',
        'exportName',
        'workerId',
        'projectId',
        'fileId',
        'runtimeEnvironmentId',
        'queueConfig',
        'retryConfig',
        'machineConfig',
        'gates',
        'queueId',
        'maxDurationInSeconds',
        'ttl',
        'triggerSource',
        'config',
        'executionMode',
        'executionPolicy',
        'payloadSchema',
    ];

    protected $attributes = [
        'triggerSource' => 'STANDARD',
        'executionMode' => 'STANDARD',
    ];

    protected $casts = [
        'queueConfig' => 'array',
        'retryConfig' => 'array',
        'machineConfig' => 'array',
        'gates' => 'array',
        'maxDurationInSeconds' => 'integer',
        'triggerSource' => TaskTriggerSource::class,
        'config' => 'array',
        'executionMode' => TaskExecutionMode::class,
        'executionPolicy' => 'array',
        'payloadSchema' => 'array',
    ];

    public function worker(): BelongsTo
    {
        return $this->belongsTo(BackgroundWorker::class, 'workerId');
    }

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function file(): BelongsTo
    {
        return $this->belongsTo(BackgroundWorkerFile::class, 'fileId');
    }

    public function runtimeEnvironment(): BelongsTo
    {
        return $this->belongsTo(RuntimeEnvironment::class, 'runtimeEnvironmentId');
    }

    public function queue(): BelongsTo
    {
        return $this->belongsTo(TaskQueue::class, 'queueId');
    }

    public function attempts(): HasMany
    {
        return $this->hasMany(TaskRunAttempt::class, 'backgroundWorkerTaskId');
    }

    public function runs(): HasMany
    {
        return $this->hasMany(TaskRun::class, 'lockedById');
    }
}