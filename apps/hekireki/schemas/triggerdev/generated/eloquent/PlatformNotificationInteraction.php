<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * Per-user tracking of notification views and dismissals
 *
 * @property \Carbon\CarbonInterface $firstSeenAt Set by beacon/CLI GET on first impression
 * @property int $showCount Times shown (incremented per beacon/CLI view)
 * @property \Carbon\CarbonInterface|null $webappDismissedAt User dismissed in webapp
 * @property \Carbon\CarbonInterface|null $webappClickedAt User clicked a link in webapp
 * @property \Carbon\CarbonInterface|null $cliDismissedAt Auto-dismissed or expired in CLI
 */
class PlatformNotificationInteraction extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'PlatformNotificationInteraction';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'notificationId',
        'userId',
        'firstSeenAt',
        'showCount',
        'webappDismissedAt',
        'webappClickedAt',
        'cliDismissedAt',
    ];

    protected $attributes = [
        'showCount' => 1,
    ];

    protected $casts = [
        'firstSeenAt' => 'datetime',
        'showCount' => 'integer',
        'webappDismissedAt' => 'datetime',
        'webappClickedAt' => 'datetime',
        'cliDismissedAt' => 'datetime',
    ];

    public function save(array $options = [])
    {
        if (! $this->exists) {
            if (! array_key_exists('firstSeenAt', $this->attributes)) {
                $this->setAttribute('firstSeenAt', $this->freshTimestamp());
            }
        }

        return parent::save($options);
    }

    public function notification(): BelongsTo
    {
        return $this->belongsTo(PlatformNotification::class, 'notificationId');
    }

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'userId');
    }
}