package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

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
class ScimSyncStateEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ScimSyncStateEntity>(ScimSyncStateTable)

    var connectionId by ScimSyncStateTable.connectionId
    var organizationId by ScimSyncStateTable.organizationId

    /** TOKEN_ISSUED | SYNCING | ERROR | REVOKED */
    var state by ScimSyncStateTable.state
    var lastPushedAt by ScimSyncStateTable.lastPushedAt

    /**
     * The failure standing right now: { op, errorCode, attempts, retiredAtMs,
     * userId, occurredAtMs }. A reason CODE and ids — never a token, never a
     * secret, never an internal hostname.
     */
    var lastFailure by ScimSyncStateTable.lastFailure

    /**
     * Every failure retired without ever being applied, newest last. The dead
     * letters an administrator has to see; retiring never drops one.
     */
    var deadLetters by ScimSyncStateTable.deadLetters

    /** revoke | teardown, once the sync has ended. */
    var revokedCause by ScimSyncStateTable.revokedCause
    var occurredAt by ScimSyncStateTable.occurredAt
    var lastEventId by ScimSyncStateTable.lastEventId
    var acceptedAt by ScimSyncStateTable.acceptedAt
    var projectionVersion by ScimSyncStateTable.projectionVersion

    /**
     * Event-derived, not database-managed: `@default(now())` / `@updatedAt`
     * would make a replayed row differ from the row it rebuilds, which is
     * exactly the parity this projection has to keep.
     */
    var createdAt by ScimSyncStateTable.createdAt
    var updatedAt by ScimSyncStateTable.updatedAt
}
