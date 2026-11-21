<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * *
 * * Error group state is used to track when a user has interacted with an error (ignored/resolved)
 * * The actual error data is in ClickHouse.
 *
 * @property \Carbon\CarbonInterface|null $ignoredUntil *
 * * Error is ignored until this date
 * @property int|null $ignoredUntilOccurrenceRate *
 * * Error is ignored until this occurrence rate
 * @property int|null $ignoredUntilTotalOccurrences *
 * * Error is ignored until this total occurrences
 * @property int|null $ignoredAtOccurrenceCount Total occurrence count at the time the error was ignored (from ClickHouse).
 * Used with ignoredUntilTotalOccurrences to compute occurrences since ignoring.
 * @property \Carbon\CarbonInterface|null $ignoredAt *
 * * Error was ignored at this date
 * @property string|null $ignoredReason *
 * * Reason for ignoring the error
 * @property string|null $ignoredByUserId *
 * * User who ignored the error
 * @property \Carbon\CarbonInterface|null $resolvedAt *
 * * Error was resolved at this date
 * @property string|null $resolvedInVersion *
 * * Error was resolved in this version
 * @property string|null $resolvedBy *
 * * User who resolved the error
 */
class ErrorGroupState extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'ErrorGroupState';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'organizationId',
        'projectId',
        'environmentId',
        'taskIdentifier',
        'errorFingerprint',
        'status',
        'ignoredUntil',
        'ignoredUntilOccurrenceRate',
        'ignoredUntilTotalOccurrences',
        'ignoredAtOccurrenceCount',
        'ignoredAt',
        'ignoredReason',
        'ignoredByUserId',
        'resolvedAt',
        'resolvedInVersion',
        'resolvedBy',
    ];

    protected $attributes = [
        'status' => 'UNRESOLVED',
    ];

    protected $casts = [
        'status' => ErrorGroupStatus::class,
        'ignoredUntil' => 'datetime',
        'ignoredUntilOccurrenceRate' => 'integer',
        'ignoredUntilTotalOccurrences' => 'integer',
        'ignoredAtOccurrenceCount' => 'integer',
        'ignoredAt' => 'datetime',
        'resolvedAt' => 'datetime',
    ];

    public function organization(): BelongsTo
    {
        return $this->belongsTo(Organization::class, 'organizationId');
    }

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    /**
     * *
     * * You can ignore/resolve an error across all environments, or specific ones
     */
    public function environment(): BelongsTo
    {
        return $this->belongsTo(RuntimeEnvironment::class, 'environmentId');
    }
}