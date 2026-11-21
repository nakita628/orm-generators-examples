package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/**
 * D08 — one connection's directory sync, as the `scim-sync` pipeline's fold
 * writes it.
 *
 * Pure event truth: every column is fold-written, rows are never deleted,
 * and REVOKED is a tombstone rather than an absence — a dead letter that
 * vanished when a connection was torn down would be a removal nobody could
 * check afterwards.
 *
 * No relation to Organization or SsoConnection, matching `SsoConnection`
 * itself: a projection is rebuilt from events, and a foreign key would make
 * a replay depend on relational state the log does not carry.
 */
object ScimSyncStateTable : IdTable<String>("\"ScimSyncState\"") {
    /** The aggregate: the connection's id, which IS its sync's id. */
    override val id = text("id").entityId()
    val connectionId = text("connectionId")
    val organizationId = text("organizationId")

    /** TOKEN_ISSUED | SYNCING | ERROR | REVOKED */
    val state = text("state")
    val lastPushedAt = pgTimestamp("lastPushedAt", 3).nullable()

    /**
     * The failure standing right now: { op, errorCode, attempts, retiredAtMs,
     * userId, occurredAtMs }. A reason CODE and ids — never a token, never a
     * secret, never an internal hostname.
     */
    val lastFailure = jsonb<String>("lastFailure", { it }, { it }).nullable()

    /**
     * Every failure retired without ever being applied, newest last. The dead
     * letters an administrator has to see; retiring never drops one.
     */
    val deadLetters = jsonb<String>("deadLetters", { it }, { it }).default("[]")

    /** revoke | teardown, once the sync has ended. */
    val revokedCause = text("revokedCause").nullable()
    val occurredAt = pgTimestamp("occurredAt", 3)
    val lastEventId = text("lastEventId")
    val acceptedAt = pgTimestamp("acceptedAt", 3)
    val projectionVersion = text("projectionVersion")

    /**
     * Event-derived, not database-managed: `@default(now())` / `@updatedAt`
     * would make a replayed row differ from the row it rebuilds, which is
     * exactly the parity this projection has to keep.
     */
    val createdAt = pgTimestamp("createdAt", 3)
    val updatedAt = pgTimestamp("updatedAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"ScimSyncState_pkey\"")

    init {
        index("\"ScimSyncState_organizationId_idx\"", false, organizationId)
        index("\"ScimSyncState_state_idx\"", false, state)
        index("\"ScimSyncState_connectionId_idx\"", false, connectionId)
    }
}
