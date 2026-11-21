<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

/**
 * This is the unified otel span/log model that will eventually be replaced by clickhouse
 *
 * @property string $message This matches the span name for a trace event, or the log body for a log event
 * @property bool $isDebug deprecated: don't use this, moving this to properties, this now uses TaskEventKind.LOG
 * @property int $startTime This is the time the event started in nanoseconds since the epoch
 * @property int $duration This is the duration of the event in nanoseconds
 * @property mixed $properties This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering
 * @property mixed $metadata This represents all span attributes in the $metadata namespace, like $metadata.payload
 * @property mixed $style This represents all span attributes in the $style namespace, like $style
 * @property mixed $output This represents all span attributes in the $output namespace, like $output
 * @property string|null $outputType This represents the mimetype of the output, such as application/json or application/super+json
 */
class TaskEvent extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = null;

    protected $table = 'TaskEvent';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'message',
        'traceId',
        'spanId',
        'parentId',
        'tracestate',
        'isError',
        'isPartial',
        'isCancelled',
        'isDebug',
        'serviceName',
        'serviceNamespace',
        'level',
        'kind',
        'status',
        'links',
        'events',
        'startTime',
        'duration',
        'attemptId',
        'attemptNumber',
        'environmentId',
        'environmentType',
        'organizationId',
        'projectId',
        'projectRef',
        'runId',
        'runIsTest',
        'idempotencyKey',
        'taskSlug',
        'taskPath',
        'taskExportName',
        'workerId',
        'workerVersion',
        'queueId',
        'queueName',
        'batchId',
        'properties',
        'metadata',
        'style',
        'output',
        'outputType',
        'payload',
        'payloadType',
        'usageDurationMs',
        'usageCostInCents',
        'machinePreset',
        'machinePresetCpu',
        'machinePresetMemory',
        'machinePresetCentsPerMs',
    ];

    protected $attributes = [
        'isError' => false,
        'isPartial' => false,
        'isCancelled' => false,
        'isDebug' => false,
        'level' => 'TRACE',
        'kind' => 'INTERNAL',
        'status' => 'UNSET',
        'duration' => 0,
        'runIsTest' => false,
        'usageDurationMs' => 0,
        'usageCostInCents' => 0.0,
    ];

    protected $casts = [
        'isError' => 'boolean',
        'isPartial' => 'boolean',
        'isCancelled' => 'boolean',
        'isDebug' => 'boolean',
        'level' => TaskEventLevel::class,
        'kind' => TaskEventKind::class,
        'status' => TaskEventStatus::class,
        'links' => 'array',
        'events' => 'array',
        'startTime' => 'integer',
        'duration' => 'integer',
        'attemptNumber' => 'integer',
        'environmentType' => RuntimeEnvironmentType::class,
        'runIsTest' => 'boolean',
        'properties' => 'array',
        'metadata' => 'array',
        'style' => 'array',
        'output' => 'array',
        'payload' => 'array',
        'usageDurationMs' => 'integer',
        'usageCostInCents' => 'float',
        'machinePresetCpu' => 'float',
        'machinePresetMemory' => 'float',
        'machinePresetCentsPerMs' => 'float',
    ];
}