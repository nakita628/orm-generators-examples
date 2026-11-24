<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class DomainEventDelivery extends Model
{
    use PrismaDates;

    const UPDATED_AT = null;
    const ZONED_DATES = ['lease_expires_at', 'next_attempt_at'];

    protected $table = 'domain_event_delivery';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'event_id',
        'consumer_id',
        'status',
        'attempts',
        'max_attempts',
        'lease_token',
        'lease_expires_at',
        'next_attempt_at',
        'last_error',
    ];

    protected $attributes = [
        'attempts' => 0,
        'max_attempts' => 12,
    ];

    protected $casts = [
        'attempts' => 'integer',
        'max_attempts' => 'integer',
        'lease_expires_at' => 'datetime',
        'next_attempt_at' => 'datetime',
    ];

    public function save(array $options = [])
    {
        if (! $this->exists) {
            if (! array_key_exists('next_attempt_at', $this->attributes)) {
                $this->setAttribute('next_attempt_at', $this->freshTimestamp());
            }
        }

        return parent::save($options);
    }
}