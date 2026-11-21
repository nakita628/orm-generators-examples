package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Durable, typed, bidirectional I/O primitive. Owns two S2 streams (.out / .in).
 * The row is essentially static — no status, no counters, no pointers. No
 * foreign keys: project/runtimeEnvironment/organization ids are plain
 * scalar columns (matches TaskRun pattern). List-style queries are served
 * from ClickHouse, not Postgres, so only point-lookup indexes live here.
 */
class SessionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SessionEntity>(SessionTable)

    var friendlyId by SessionTable.friendlyId

    /**
     * User-supplied identifier scoped to the environment. Used for
     * idempotent upsert and for resolving sessions via the public API.
     */
    var externalId by SessionTable.externalId

    /** Plain string — intentionally not an enum. */
    var type by SessionTable.type

    /** Denormalized scoping columns — no FK relations. */
    var projectId by SessionTable.projectId
    var runtimeEnvironmentId by SessionTable.runtimeEnvironmentId
    var environmentType by SessionTable.environmentType
    var organizationId by SessionTable.organizationId

    /**
     * Task this session triggers runs against. Required — Sessions are
     * task-bound: creating a session also triggers its first run, and
     * every subsequent re-trigger uses this same identifier.
     */
    var taskIdentifier by SessionTable.taskIdentifier

    /**
     * Trigger config used for every run this session schedules. Shape
     * (validated at the route layer, opaque to the DB):
     * { basePayload: object, machine?: string, queue?: string,
     * tags?: string[], maxAttempts?: number,
     * idleTimeoutInSeconds?: number }
     * `basePayload` carries the customer's client-data; runtime fields
     * (chatId, messages, trigger) are merged at trigger time.
     */
    var triggerConfig by SessionTable.triggerConfig

    /**
     * Whether this session was created from the Test/playground UI rather
     * than by a real chat.agent() trigger. Mirrors TaskRun.isTest.
     */
    var isTest by SessionTable.isTest
    var tags by SessionTable.tags
    var metadata by SessionTable.metadata

    /**
     * Live run pointer — non-FK so run deletion never cascades. Can lag
     * reality; the `.in/append` handler re-checks the snapshot status
     * before reusing it.
     */
    var currentRunId by SessionTable.currentRunId

    /**
     * Monotonic counter used for optimistic locking on `currentRunId`
     * swaps. Bumped atomically alongside any update that changes
     * `currentRunId`.
     */
    var currentRunVersion by SessionTable.currentRunVersion

    /** Terminal markers — written once, never flipped back. */
    var closedAt by SessionTable.closedAt
    var closedReason by SessionTable.closedReason
    var expiresAt by SessionTable.expiresAt
    var createdAt by SessionTable.createdAt
    var updatedAt by SessionTable.updatedAt

    /**
     * S2 basin where this session's stream pair lives. Stamped at create
     * time from `Organization.streamBasinName` so reads can resolve the
     * basin without joining org. Null when the org has no per-org basin
     * (OSS, or pre-backfill); reads fall back to the global basin.
     */
    var streamBasinName by SessionTable.streamBasinName

    /**
     * Storage URI (with protocol prefix) for this session's chat.agent
     * snapshot blob. Set on first snapshot write. Null = pre-column session,
     * fall back to computed default path.
     */
    var chatSnapshotStoragePath by SessionTable.chatSnapshotStoragePath
    val runs by SessionRunEntity referrersOn SessionRunTable.sessionId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SessionTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
