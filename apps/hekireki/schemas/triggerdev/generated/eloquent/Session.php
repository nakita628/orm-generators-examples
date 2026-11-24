<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * Durable, typed, bidirectional I/O primitive. Owns two S2 streams (.out / .in).
 * The row is essentially static — no status, no counters, no pointers. No
 * foreign keys: project/runtimeEnvironment/organization ids are plain
 * scalar columns (matches TaskRun pattern). List-style queries are served
 * from ClickHouse, not Postgres, so only point-lookup indexes live here.
 *
 * @property string|null $externalId User-supplied identifier scoped to the environment. Used for
 * idempotent upsert and for resolving sessions via the public API.
 * @property string $type Plain string — intentionally not an enum.
 * @property string $projectId Denormalized scoping columns — no FK relations.
 * @property string $taskIdentifier Task this session triggers runs against. Required — Sessions are
 * task-bound: creating a session also triggers its first run, and
 * every subsequent re-trigger uses this same identifier.
 * @property mixed $triggerConfig Trigger config used for every run this session schedules. Shape
 * (validated at the route layer, opaque to the DB):
 * { basePayload: object, machine?: string, queue?: string,
 * tags?: string[], maxAttempts?: number,
 * idleTimeoutInSeconds?: number }
 * `basePayload` carries the customer's client-data; runtime fields
 * (chatId, messages, trigger) are merged at trigger time.
 * @property bool $isTest Whether this session was created from the Test/playground UI rather
 * than by a real chat.agent() trigger. Mirrors TaskRun.isTest.
 * @property string|null $currentRunId Live run pointer — non-FK so run deletion never cascades. Can lag
 * reality; the `.in/append` handler re-checks the snapshot status
 * before reusing it.
 * @property int $currentRunVersion Monotonic counter used for optimistic locking on `currentRunId`
 * swaps. Bumped atomically alongside any update that changes
 * `currentRunId`.
 * @property \Carbon\CarbonInterface|null $closedAt Terminal markers — written once, never flipped back.
 * @property string|null $streamBasinName S2 basin where this session's stream pair lives. Stamped at create
 * time from `Organization.streamBasinName` so reads can resolve the
 * basin without joining org. Null when the org has no per-org basin
 * (OSS, or pre-backfill); reads fall back to the global basin.
 * @property string|null $chatSnapshotStoragePath Storage URI (with protocol prefix) for this session's chat.agent
 * snapshot blob. Set on first snapshot write. Null = pre-column session,
 * fall back to computed default path.
 */
class Session extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'Session';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'friendlyId',
        'externalId',
        'type',
        'projectId',
        'runtimeEnvironmentId',
        'environmentType',
        'organizationId',
        'taskIdentifier',
        'triggerConfig',
        'isTest',
        'tags',
        'metadata',
        'currentRunId',
        'currentRunVersion',
        'closedAt',
        'closedReason',
        'expiresAt',
        'streamBasinName',
        'chatSnapshotStoragePath',
    ];

    protected $attributes = [
        'isTest' => false,
        'currentRunVersion' => 0,
    ];

    protected $casts = [
        'environmentType' => RuntimeEnvironmentType::class,
        'triggerConfig' => 'array',
        'isTest' => 'boolean',
        'tags' => AsPrismaList::class,
        'metadata' => 'array',
        'currentRunVersion' => 'integer',
        'closedAt' => 'datetime',
        'expiresAt' => 'datetime',
    ];

    public function runs(): HasMany
    {
        return $this->hasMany(SessionRun::class, 'sessionId');
    }
}