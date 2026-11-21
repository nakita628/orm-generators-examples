<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * Admin-created notification definitions
 *
 * @property string $title Admin-facing title for identification in admin tools
 * @property mixed $payload Versioned JSON rendering content (see payload schema in spec)
 * @property \Carbon\CarbonInterface $startsAt When notification becomes active
 * @property \Carbon\CarbonInterface $endsAt When notification expires
 * @property int|null $cliMaxDaysAfterFirstSeen CLI: auto-expire N days after user first saw it
 * @property int|null $cliMaxShowCount CLI: auto-expire after shown N times to user
 * @property int|null $cliShowEvery CLI: only show every N-th request (e.g. cliShowEvery=3 means show on 3rd, 6th, 9th…)
 * @property int $priority Ordering within same scope level (higher = more important)
 * @property bool $isDraft Staged-but-unpublished. While true the notification is hidden from all
 * user-facing reads regardless of startsAt/endsAt; publishing sets real
 * dates and flips this to false.
 * @property \Carbon\CarbonInterface|null $archivedAt Soft delete
 */
class PlatformNotification extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'PlatformNotification';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'friendlyId',
        'title',
        'payload',
        'surface',
        'scope',
        'userId',
        'organizationId',
        'projectId',
        'startsAt',
        'endsAt',
        'cliMaxDaysAfterFirstSeen',
        'cliMaxShowCount',
        'cliShowEvery',
        'priority',
        'isDraft',
        'archivedAt',
    ];

    protected $attributes = [
        'priority' => 0,
        'isDraft' => false,
    ];

    protected $casts = [
        'payload' => 'array',
        'surface' => PlatformNotificationSurface::class,
        'scope' => PlatformNotificationScope::class,
        'startsAt' => 'datetime',
        'endsAt' => 'datetime',
        'cliMaxDaysAfterFirstSeen' => 'integer',
        'cliMaxShowCount' => 'integer',
        'cliShowEvery' => 'integer',
        'priority' => 'integer',
        'isDraft' => 'boolean',
        'archivedAt' => 'datetime',
    ];

    /** Set when scope = USER */
    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'userId');
    }

    /** Set when scope = ORGANIZATION */
    public function organization(): BelongsTo
    {
        return $this->belongsTo(Organization::class, 'organizationId');
    }

    /** Set when scope = PROJECT */
    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function interactions(): HasMany
    {
        return $this->hasMany(PlatformNotificationInteraction::class, 'notificationId');
    }
}