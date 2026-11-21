<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * Records of previously-valid API keys that are still accepted for authentication
 * during a grace window after rotation. Extend or end the grace period by updating `expiresAt`.
 */
class RevokedApiKey extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = null;

    protected $table = 'RevokedApiKey';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'apiKey',
        'runtimeEnvironmentId',
        'expiresAt',
    ];

    protected $casts = [
        'expiresAt' => 'datetime',
    ];

    public function runtimeEnvironment(): BelongsTo
    {
        return $this->belongsTo(RuntimeEnvironment::class, 'runtimeEnvironmentId');
    }
}