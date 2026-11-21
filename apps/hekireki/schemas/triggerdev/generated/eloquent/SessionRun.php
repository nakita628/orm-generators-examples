<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * Historical record of every run a Session has owned. Append-only —
 * rows are inserted on each `ensureRunForSession` claim, never updated.
 * Lets us reconstruct the run timeline of a chat for debugging /
 * dashboard surfaces. The relation cascades on Session delete (tied to
 * the session lifecycle) but `runId` is a plain string column with no
 * FK to TaskRun so run pruning is independent.
 *
 * @property string $runId TaskRun.id (no FK — runs may be archived independently of session history)
 * @property string $reason One of: "initial" | "continuation" | "upgrade" | "manual".
 * Plain string for forward-compat with future trigger reasons.
 */
class SessionRun extends Model
{
    use PrismaDates;

    protected $table = 'SessionRun';

    protected $keyType = 'string';

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'id',
        'sessionId',
        'runId',
        'reason',
        'triggeredAt',
    ];

    protected $casts = [
        'triggeredAt' => 'datetime',
    ];

    public function save(array $options = [])
    {
        if (! $this->exists) {
            if (! array_key_exists('triggeredAt', $this->attributes)) {
                $this->setAttribute('triggeredAt', $this->freshTimestamp());
            }
        }

        return parent::save($options);
    }

    public function session(): BelongsTo
    {
        return $this->belongsTo(Session::class, 'sessionId');
    }
}