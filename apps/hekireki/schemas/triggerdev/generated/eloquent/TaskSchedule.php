<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * Schedules can be attached to tasks to trigger them on a schedule
 *
 * @property string $friendlyId users see this as `id`. They start with schedule_
 * @property string $taskIdentifier a reference to a task (not a foreign key because it's across versions)
 * @property string $deduplicationKey can be provided and we won't create another with the same key
 * @property string $generatorExpression the CRON pattern
 * @property string $timezone These are IANA format string, or the default "UTC". E.g. "America/New_York"
 * @property string|null $externalId Can be provided by the user then accessed inside a run
 * @property \Carbon\CarbonInterface|null $lastRunTriggeredAt @deprecated stop writing 2026-04-30; reads moved out of code (UI now derives from cron's previous slot). Drop in follow-up.
 */
class TaskSchedule extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'TaskSchedule';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'type',
        'friendlyId',
        'taskIdentifier',
        'deduplicationKey',
        'userProvidedDeduplicationKey',
        'generatorExpression',
        'generatorDescription',
        'generatorType',
        'timezone',
        'windowDurationSeconds',
        'windowPercentage',
        'defaultWindowDurationSeconds',
        'minimumWindowDurationSeconds',
        'externalId',
        'lastRunTriggeredAt',
        'projectId',
        'active',
    ];

    protected $attributes = [
        'type' => 'IMPERATIVE',
        'userProvidedDeduplicationKey' => false,
        'generatorDescription' => '',
        'generatorType' => 'CRON',
        'timezone' => 'UTC',
        'active' => true,
    ];

    protected $casts = [
        'type' => ScheduleType::class,
        'userProvidedDeduplicationKey' => 'boolean',
        'generatorType' => ScheduleGeneratorType::class,
        'windowDurationSeconds' => 'integer',
        'windowPercentage' => 'integer',
        'defaultWindowDurationSeconds' => 'integer',
        'minimumWindowDurationSeconds' => 'integer',
        'lastRunTriggeredAt' => 'datetime',
        'active' => 'boolean',
    ];

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    /** Instances of the schedule that are active */
    public function instances(): HasMany
    {
        return $this->hasMany(TaskScheduleInstance::class, 'taskScheduleId');
    }
}