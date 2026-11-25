package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/** An organisation's agent registry settings. No row reads as the defaults. */
class AgentRegistrySettingsEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AgentRegistrySettingsEntity>(AgentRegistrySettingsTable)

    /** Days without traffic after which a listed agent reads as dormant. */
    var dormantAfterDays by AgentRegistrySettingsTable.dormantAfterDays
    var updatedAt by AgentRegistrySettingsTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AgentRegistrySettingsTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
