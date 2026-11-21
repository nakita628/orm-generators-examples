<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasVersion4Uuids;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property \Carbon\CarbonInterface|null $snoozedUntil When set, `subscribed = false` is temporary rather than permanent: the snooze
 * sweep resubscribes the contact once this passes. Suppression itself is still
 * carried entirely by `subscribed`, so every send path filters snoozed contacts
 * out without knowing this column exists.
 *
 * Cleared by every other write to `subscribed` — including the bounce and
 * complaint paths, so a suppressed address can never be woken back up.
 */
class Contact extends Model
{
    use HasVersion4Uuids;
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'contacts';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'email',
        'data',
        'subscribed',
        'snoozedUntil',
        'projectId',
    ];

    protected $attributes = [
        'subscribed' => true,
    ];

    protected $casts = [
        'data' => 'array',
        'subscribed' => 'boolean',
        'snoozedUntil' => 'datetime',
    ];

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function emails(): HasMany
    {
        return $this->hasMany(Email::class, 'contactId');
    }

    public function workflowExecutions(): HasMany
    {
        return $this->hasMany(WorkflowExecution::class, 'contactId');
    }

    public function events(): HasMany
    {
        return $this->hasMany(Event::class, 'contactId');
    }

    public function segmentMemberships(): HasMany
    {
        return $this->hasMany(SegmentMembership::class, 'contactId');
    }

    public function throttleClaims(): HasMany
    {
        return $this->hasMany(ThrottleClaim::class, 'contactId');
    }
}