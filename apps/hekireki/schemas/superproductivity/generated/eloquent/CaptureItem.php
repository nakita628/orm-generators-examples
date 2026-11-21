<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * One sealed capture waiting for a device. No relation to its grant: items
 * outlive a grant that is revoked without discarding them.
 *
 * @property string $task_id Task id the claiming device must use; replaced when another device takes over.
 */
class CaptureItem extends Model
{
    protected $table = 'capture_items';

    protected $keyType = 'string';

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'id',
        'user_id',
        'grant_id',
        'ciphertext',
        'task_id',
        'created_at',
        'lease_device_id',
        'lease_until',
    ];

    protected $casts = [
        'user_id' => 'integer',
        'created_at' => 'integer',
        'lease_until' => 'integer',
    ];

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'user_id');
    }
}