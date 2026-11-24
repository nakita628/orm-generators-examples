<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property string $secret AES-256-GCM under SECRET_PASSWORD (see @rallly/utils/encryption); the plaintext is shown once at creation
 * @property list<string> $events Subscribed event names (`poll.closed`, ...); validated by the feature schema, not a DB enum, so the vocabulary grows without migrations
 * @property string $version The payload contract this endpoint was built against (WEBHOOK_VERSION at creation). No default: a writer that forgets it fails instead of recording a stale contract. A breaking payload change branches on this so existing endpoints keep their shape
 * @property \Carbon\CarbonInterface $cursor Fan-out position in the poll activity log: activities created after this instant are still to be delivered. Starts at creation so a new endpoint never replays history
 * @property int $consecutive_failures Exhausted deliveries in a row; reset by any success. The dispatcher disables the endpoint once it reaches its cap
 */
class SpaceWebhook extends Model
{
    use PrismaDates;

    protected $table = 'space_webhooks';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'space_id',
        'url',
        'secret',
        'events',
        'enabled',
        'version',
        'cursor',
        'last_delivered_at',
        'last_failed_at',
        'consecutive_failures',
    ];

    protected $attributes = [
        'enabled' => true,
        'consecutive_failures' => 0,
    ];

    protected $casts = [
        'events' => AsPrismaList::class,
        'enabled' => 'boolean',
        'cursor' => 'datetime',
        'last_delivered_at' => 'datetime',
        'last_failed_at' => 'datetime',
        'consecutive_failures' => 'integer',
    ];

    public function save(array $options = [])
    {
        if (! $this->exists) {
            if (! array_key_exists('cursor', $this->attributes)) {
                $this->setAttribute('cursor', $this->freshTimestamp());
            }
        }

        return parent::save($options);
    }

    public function space(): BelongsTo
    {
        return $this->belongsTo(Space::class, 'space_id');
    }

    public function deliveries(): HasMany
    {
        return $this->hasMany(WebhookDelivery::class, 'webhook_id');
    }
}