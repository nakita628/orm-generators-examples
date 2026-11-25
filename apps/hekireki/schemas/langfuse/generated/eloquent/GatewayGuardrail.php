<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class GatewayGuardrail extends \Illuminate\Database\Eloquent\Model
{
    use PrismaDates;

    protected $table = 'gateway_guardrails';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'organization_id',
        'name',
        'type',
        'action',
        'config',
        'encrypted_credentials',
        'created_by',
    ];

    protected $attributes = [
        'type' => 'REGEX',
        'config' => '{}',
    ];

    protected $casts = [
        'type' => GatewayGuardrailType::class,
        'action' => GatewayGuardrailAction::class,
        'config' => 'array',
    ];

    public function organization(): BelongsTo
    {
        return $this->belongsTo(Organization::class, 'organization_id');
    }

    public function createdByUser(): BelongsTo
    {
        return $this->belongsTo(User::class, 'created_by');
    }

    public function assignments(): HasMany
    {
        return $this->hasMany(GatewayGuardrailAssignment::class, 'guardrail_id');
    }
}