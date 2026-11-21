<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string $type Validated by a zod union in the feature, not a DB enum: the activity vocabulary grows without migrations
 * @property string|null $user_id Subject ids are soft references (no FKs): activity must outlive its subjects — including account deletion — so refs persist as identity and the payload carries the display snapshot; whether a ref still resolves is the renderer's question
 */
class PollActivity extends Model
{
    use PrismaDates;

    const UPDATED_AT = null;

    protected $table = 'poll_activities';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'poll_id',
        'type',
        'user_id',
        'participant_id',
        'invite_id',
        'option_id',
        'payload',
    ];

    protected $attributes = [
        'payload' => '{}',
    ];

    protected $casts = [
        'payload' => 'array',
    ];

    public function poll(): BelongsTo
    {
        return $this->belongsTo(Poll::class, 'poll_id');
    }
}