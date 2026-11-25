<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class FieldPeriodicRefreshState extends Model
{
    use PrismaDates;

    const CREATED_AT = null;
    const UPDATED_AT = 'last_modified_time';
    const ZONED_DATES = ['last_refreshed_at', 'pending_started_at', 'created_time', 'last_modified_time'];

    protected $table = 'field_periodic_refresh_state';

    protected $primaryKey = 'field_id';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'field_id',
        'table_id',
        'base_id',
        'enabled',
        'time_zone',
        'contains_now_precision',
        'last_refreshed_at',
        'last_refreshed_local_date',
        'last_range_hash',
        'last_status',
        'last_error',
        'pending_task_id',
        'pending_local_date',
        'pending_range_hash',
        'pending_started_at',
        'created_time',
    ];

    protected $attributes = [
        'enabled' => true,
        'time_zone' => 'UTC',
        'contains_now_precision' => false,
        'last_status' => 'idle',
    ];

    protected $casts = [
        'enabled' => 'boolean',
        'contains_now_precision' => 'boolean',
        'last_refreshed_at' => 'datetime',
        'pending_started_at' => 'datetime',
        'created_time' => 'datetime',
    ];

    public function save(array $options = [])
    {
        if (! $this->exists) {
            if (! array_key_exists('created_time', $this->attributes)) {
                $this->setAttribute('created_time', $this->freshTimestamp());
            }
        }

        return parent::save($options);
    }

    public function field(): BelongsTo
    {
        return $this->belongsTo(Field::class, 'field_id');
    }
}