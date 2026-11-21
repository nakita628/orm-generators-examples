<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string|null $description Optional one-line note the creator adds to say what the key is for.
 * @property mixed $allowed_ip_ranges CIDR ranges this key may be used from. Empty means any IP.
 */
class ApiKey extends Model
{
    use PrismaDates;

    protected $table = 'api_keys';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'name',
        'key_hash',
        'last_four',
        'description',
        'runtime_environment_id',
        'created_by_user_id',
        'issued_by',
        'preset_id',
        'scopes',
        'allowed_ip_ranges',
        'last_used_at',
        'last_used_ip',
        'revoked_at',
        'expires_at',
    ];

    protected $attributes = [
        'issued_by' => 'USER',
    ];

    protected $casts = [
        'issued_by' => ApiKeyIssuer::class,
        'last_used_at' => 'datetime',
        'revoked_at' => 'datetime',
        'expires_at' => 'datetime',
    ];

    public function runtimeEnvironment(): BelongsTo
    {
        return $this->belongsTo(RuntimeEnvironment::class, 'runtime_environment_id');
    }

    public function createdBy(): BelongsTo
    {
        return $this->belongsTo(User::class, 'created_by_user_id');
    }
}