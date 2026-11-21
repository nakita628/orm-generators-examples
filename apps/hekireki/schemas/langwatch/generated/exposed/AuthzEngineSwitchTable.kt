package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Whether an organization's permission checks run on the authorization engine: authz's own
 * switch, set when the engine migration proves the organization and pinned off by an operator's
 * rollback. No row is the legacy path. Spec: modules/authz/specs/engine-switch.feature.
 */
object AuthzEngineSwitchTable : IdTable<String>("\"AuthzEngineSwitch\"") {
    override val id = text("organizationId").entityId()
    val onEngine = bool("onEngine")
    val pinned = bool("pinned").default(false)
    val switchedAt = pgTimestamp("switchedAt", 3)
    val lastAppliedEventId = text("lastAppliedEventId").nullable()
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AuthzEngineSwitch_pkey\"")
}
