<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * An email waiting to be sent by the queue, one row per recipient. The
 * email is built when it is sent, from the row its subject points at.
 *
 * @property string $kind Which handler builds the email. A string rather than an enum so a new
 * kind needs no migration; the claim only takes kinds the build knows.
 * @property string $subject_id The row the email is about, interpreted by the kind. Not a foreign key:
 * kinds point at different tables, and a subject deleted before the send
 * makes its handler skip the email.
 * @property string $batch_id Groups the emails one action queued, so the action can send its own
 * first batch at once and its progress can be read back.
 * @property string $user_id The account whose action queued the email. Banning it skips whatever
 * is still pending.
 * @property \Carbon\CarbonInterface|null $send_after Not claimed before this time. Unset sends at the next run.
 */
class QueuedEmail extends Model
{
    use PrismaDates;

    protected $table = 'queued_emails';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'kind',
        'subject_id',
        'batch_id',
        'user_id',
        'status',
        'send_after',
        'attempts',
        'claimed_at',
        'sent_at',
        'last_error',
    ];

    protected $attributes = [
        'status' => 'pending',
        'attempts' => 0,
    ];

    protected $casts = [
        'status' => QueuedEmailStatus::class,
        'send_after' => 'datetime',
        'attempts' => 'integer',
        'claimed_at' => 'datetime',
        'sent_at' => 'datetime',
    ];

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'user_id');
    }
}