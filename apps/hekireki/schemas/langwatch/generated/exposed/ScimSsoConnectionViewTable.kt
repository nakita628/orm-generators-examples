package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/**
 * SCIM's own copy of an organization's SSO connections: a peer fold over
 * identity's connection facts (ARCHITECTURE §9), folded with identity-contract's
 * reducer, one row per connection. SCIM reads this instead of asking identity.
 * Rebuilt from identity's event log by a projection replay, so no relation:
 * a foreign key would make a replay depend on state the log does not carry.
 */
object ScimSsoConnectionViewTable : IdTable<String>("\"ScimSsoConnectionView\"") {
    /** The aggregate: identity's connection id. */
    override val id = text("id").entityId()
    val organizationId = text("organizationId")

    /** The connection as the reducer folded it (SsoConnectionState + LastEventOccurredAt). */
    val folded = jsonb<String>("folded", { it }, { it })

    /** The event ids of the last delivery folded in, so a redelivery folds once. */
    val appliedEventIds = array("appliedEventIds", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /** A row of another version reads as undecodable and is re-folded from the log. */
    val projectionVersion = text("projectionVersion")

    /** Event-derived, so a replayed row equals the row it rebuilds. */
    val createdAt = pgTimestamp("createdAt", 3)
    val updatedAt = pgTimestamp("updatedAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"ScimSsoConnectionView_pkey\"")

    init {
        index("\"ScimSsoConnectionView_organizationId_createdAt_idx\"", false, organizationId, createdAt)
    }
}
