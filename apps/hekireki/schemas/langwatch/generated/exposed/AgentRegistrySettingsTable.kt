package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/** An organisation's agent registry settings. No row reads as the defaults. */
object AgentRegistrySettingsTable : IdTable<String>("\"AgentRegistrySettings\"") {
    override val id = text("organizationId").entityId()

    /** Days without traffic after which a listed agent reads as dormant. */
    val dormantAfterDays = integer("dormantAfterDays").default(30)
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AgentRegistrySettings_pkey\"")
}
