<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * Public half of the user's capture key. The row outlives switching the
 * feature off: `disabledAt` stops a stale device re-registering that key.
 *
 * @property string $public_key Base64 SPKI DER of an ECDH P-256 public key.
 */
class CaptureKey extends Model
{
    protected $table = 'capture_keys';

    protected $primaryKey = 'user_id';

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'user_id',
        'public_key',
        'fingerprint',
        'created_at',
        'disabled_at',
    ];

    protected $casts = [
        'created_at' => 'integer',
        'disabled_at' => 'integer',
    ];

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'user_id');
    }
}