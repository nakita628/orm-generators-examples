<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;
use Illuminate\Database\Eloquent\Relations\HasOne;

/** @property string $token Per response edit link: grants edit access to this response without a session. App minted at creation, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet — url safe with no linkifier edge cases, unlike base64url); never a DB default so a code path that forgets to mint fails loudly. A response created through an emailed invite takes the invite's token. Rows that predate minting were backfilled with gen_random_uuid() hex */
class Participant extends Model
{
    use PrismaDates;

    protected $table = 'participants';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'name',
        'email',
        'note',
        'token',
        'user_id',
        'poll_id',
        'locale',
        'time_zone',
    ];

    public function poll(): BelongsTo
    {
        return $this->belongsTo(Poll::class, 'poll_id');
    }

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'user_id');
    }

    public function invite(): HasOne
    {
        return $this->hasOne(PollInvite::class, 'participant_id');
    }

    public function votes(): HasMany
    {
        return $this->hasMany(Vote::class, 'participant_id');
    }
}