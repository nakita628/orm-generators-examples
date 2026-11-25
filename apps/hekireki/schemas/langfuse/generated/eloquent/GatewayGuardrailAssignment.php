<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Relations\BelongsTo;

class GatewayGuardrailAssignment extends \Illuminate\Database\Eloquent\Model
{
    use PrismaDates;

    const UPDATED_AT = null;

    protected $table = 'gateway_guardrail_assignments';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'guardrail_id',
        'organization_id',
        'scope',
        'api_key_id',
        'user_id',
    ];

    protected $casts = [
        'scope' => GatewayGuardrailAssignmentScope::class,
    ];

    public function guardrail(): BelongsTo
    {
        return $this->belongsTo(GatewayGuardrail::class, 'guardrail_id');
    }

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
}