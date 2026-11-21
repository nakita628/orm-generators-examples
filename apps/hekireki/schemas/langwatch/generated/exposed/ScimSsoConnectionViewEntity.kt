package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * SCIM's own copy of an organization's SSO connections: a peer fold over
 * identity's connection facts (ARCHITECTURE §9), folded with identity-contract's
 * reducer, one row per connection. SCIM reads this instead of asking identity.
 * Rebuilt from identity's event log by a projection replay, so no relation:
 * a foreign key would make a replay depend on state the log does not carry.
 */
class ScimSsoConnectionViewEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ScimSsoConnectionViewEntity>(ScimSsoConnectionViewTable)

    var organizationId by ScimSsoConnectionViewTable.organizationId

    /** The connection as the reducer folded it (SsoConnectionState + LastEventOccurredAt). */
    var folded by ScimSsoConnectionViewTable.folded

    /** The event ids of the last delivery folded in, so a redelivery folds once. */
    var appliedEventIds by ScimSsoConnectionViewTable.appliedEventIds

    /** A row of another version reads as undecodable and is re-folded from the log. */
    var projectionVersion by ScimSsoConnectionViewTable.projectionVersion

    /** Event-derived, so a replayed row equals the row it rebuilds. */
    var createdAt by ScimSsoConnectionViewTable.createdAt
    var updatedAt by ScimSsoConnectionViewTable.updatedAt
}
