package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Durable, typed, bidirectional I/O primitive. Owns two S2 streams (.out / .in).
 * The row is essentially static — no status, no counters, no pointers. No
 * foreign keys: project/runtimeEnvironment/organization ids are plain
 * scalar columns (matches TaskRun pattern). List-style queries are served
 * from ClickHouse, not Postgres, so only point-lookup indexes live here.
 */
object SessionTable : IdTable<String>("\"Session\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val friendlyId = text("friendlyId").uniqueIndex("\"Session_friendlyId_key\"")

    /**
     * User-supplied identifier scoped to the environment. Used for
     * idempotent upsert and for resolving sessions via the public API.
     */
    val externalId = text("externalId").nullable()

    /** Plain string — intentionally not an enum. */
    val type = text("type")

    /** Denormalized scoping columns — no FK relations. */
    val projectId = text("projectId")
    val runtimeEnvironmentId = text("runtimeEnvironmentId")
    val environmentType =
        pgEnum("environmentType", "\"RuntimeEnvironmentType\"", RuntimeEnvironmentType.entries, RuntimeEnvironmentType::dbName)
    val organizationId = text("organizationId")

    /**
     * Task this session triggers runs against. Required — Sessions are
     * task-bound: creating a session also triggers its first run, and
     * every subsequent re-trigger uses this same identifier.
     */
    val taskIdentifier = text("taskIdentifier")

    /**
     * Trigger config used for every run this session schedules. Shape
     * (validated at the route layer, opaque to the DB):
     * { basePayload: object, machine?: string, queue?: string,
     * tags?: string[], maxAttempts?: number,
     * idleTimeoutInSeconds?: number }
     * `basePayload` carries the customer's client-data; runtime fields
     * (chatId, messages, trigger) are merged at trigger time.
     */
    val triggerConfig = jsonb<String>("triggerConfig", { it }, { it })

    /**
     * Whether this session was created from the Test/playground UI rather
     * than by a real chat.agent() trigger. Mirrors TaskRun.isTest.
     */
    val isTest = bool("isTest").default(false)
    val tags = array("tags", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }
    val metadata = jsonb<String>("metadata", { it }, { it }).nullable()

    /**
     * Live run pointer — non-FK so run deletion never cascades. Can lag
     * reality; the `.in/append` handler re-checks the snapshot status
     * before reusing it.
     */
    val currentRunId = text("currentRunId").nullable()

    /**
     * Monotonic counter used for optimistic locking on `currentRunId`
     * swaps. Bumped atomically alongside any update that changes
     * `currentRunId`.
     */
    val currentRunVersion = integer("currentRunVersion").default(0)

    /** Terminal markers — written once, never flipped back. */
    val closedAt = pgTimestamp("closedAt", 3).nullable()
    val closedReason = text("closedReason").nullable()
    val expiresAt = pgTimestamp("expiresAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /**
     * S2 basin where this session's stream pair lives. Stamped at create
     * time from `Organization.streamBasinName` so reads can resolve the
     * basin without joining org. Null when the org has no per-org basin
     * (OSS, or pre-backfill); reads fall back to the global basin.
     */
    val streamBasinName = text("streamBasinName").nullable()

    /**
     * Storage URI (with protocol prefix) for this session's chat.agent
     * snapshot blob. Set on first snapshot write. Null = pre-column session,
     * fall back to computed default path.
     */
    val chatSnapshotStoragePath = text("chatSnapshotStoragePath").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"Session_pkey\"")

    init {
        index("\"Session_expiresAt_idx\"", false, expiresAt)
        index("\"Session_currentRunId_idx\"", false, currentRunId)
        uniqueIndex("\"Session_runtimeEnvironmentId_externalId_key\"", runtimeEnvironmentId, externalId)
    }
}
