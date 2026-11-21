<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string $token Per invitee link. App minted, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet) — no schema default, so forgetting to mint fails loudly. Rotated on reactivation so a revoked (possibly leaked) token never regains access
 * @property \Carbon\CarbonInterface|null $revoked_at Revocation is always explicit: set when the host removes a pending invite, never as a side effect of close, schedule or reopen. Deleting a response instead hard deletes the participant, so the SetNull FK clears participantId and reverts the invite to pending with its token intact — the token the deleted response held is freed with it, so the next response through the same link can take it again. Cleared only by re-inviting (reactivation), which also rotates the token
 * @property string|null $participant_id Conversion join, merges invite history into the response timeline. Set only when the invitee responds via their token: a typed email is unverified, so a generic link response never claims an invite
 */
class PollInvite extends Model
{
    use PrismaDates;

    protected $table = 'poll_invites';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'poll_id',
        'email',
        'token',
        'opened_at',
        'reminded_at',
        'revoked_at',
        'participant_id',
    ];

    protected $casts = [
        'opened_at' => 'datetime',
        'reminded_at' => 'datetime',
        'revoked_at' => 'datetime',
    ];

    public function poll(): BelongsTo
    {
        return $this->belongsTo(Poll::class, 'poll_id');
    }

    public function participant(): BelongsTo
    {
        return $this->belongsTo(Participant::class, 'participant_id');
    }
}