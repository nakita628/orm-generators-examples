<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class GatewaySpendLimit extends \Illuminate\Database\Eloquent\Model
{
    use PrismaDates;

    protected $table = 'gateway_spend_limits';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'organization_id',
        'scope',
        'api_key_id',
        'user_id',
        'interval',
        'limit_usd',
        'alert_thresholds',
        'counting_started_at',
        'owner_user_id',
    ];

    protected $casts = [
        'scope' => GatewaySpendLimitScope::class,
        'interval' => GatewaySpendLimitInterval::class,
        'limit_usd' => AsDecimal::class,
        'alert_thresholds' => AsPrismaList::class . ':int',
        'counting_started_at' => 'datetime',
    ];

    public function organization(): BelongsTo
    {
        return $this->belongsTo(Organization::class, 'organization_id');
    }

    public function apiKey(): BelongsTo
    {
        return $this->belongsTo(ApiKey::class, 'api_key_id');
    }

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'user_id');
    }

    public function ownerUser(): BelongsTo
    {
        return $this->belongsTo(User::class, 'owner_user_id');
    }

    public function alerts(): HasMany
    {
        return $this->hasMany(GatewaySpendLimitAlert::class, 'spend_limit_id');
    }
}