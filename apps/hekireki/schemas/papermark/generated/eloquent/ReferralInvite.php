<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string $scopeId *
 * * Whose cap and discount this invite counts against — the fund's team, or the
 * * sender alone. One non-null key rather than a pair of nullable ones: on a
 * * nullable teamId every personal invite is distinct to Postgres, so the two
 * * keys guarded different scopes and a member could not invite a company both
 * * for their fund and for themselves.
 */
class ReferralInvite extends Model
{
    use PrismaDates;

    protected $table = 'ReferralInvite';

    protected $keyType = 'string';

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'id',
        'userId',
        'teamId',
        'scopeId',
        'email',
        'referralLink',
        'sentAt',
    ];

    protected $casts = [
        'sentAt' => 'datetime',
    ];

    public function save(array $options = [])
    {
        if (! $this->exists) {
            if (! array_key_exists('sentAt', $this->attributes)) {
                $this->setAttribute('sentAt', $this->freshTimestamp());
            }
        }

        return parent::save($options);
    }

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'userId');
    }

    public function team(): BelongsTo
    {
        return $this->belongsTo(Team::class, 'teamId');
    }
}