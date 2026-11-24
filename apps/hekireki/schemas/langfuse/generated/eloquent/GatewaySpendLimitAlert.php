<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Relations\BelongsTo;

/** One notification per limit, subject, alert type, threshold, counting start and window, so the gateway may report a limit any number of times. */
class GatewaySpendLimitAlert extends \Illuminate\Database\Eloquent\Model
{
    use PrismaDates;

    const UPDATED_AT = null;

    protected $table = 'gateway_spend_limit_alerts';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'spend_limit_id',
        'subject_id',
        'type',
        'threshold_percent',
        'counting_started_at',
        'window_start',
    ];

    protected $attributes = [
        'subject_id' => '',
        'threshold_percent' => 100,
    ];

    protected $casts = [
        'type' => GatewaySpendLimitAlertType::class,
        'threshold_percent' => 'integer',
        'counting_started_at' => 'datetime',
        'window_start' => 'datetime',
    ];

    public function spendLimit(): BelongsTo
    {
        return $this->belongsTo(GatewaySpendLimit::class, 'spend_limit_id');
    }
}