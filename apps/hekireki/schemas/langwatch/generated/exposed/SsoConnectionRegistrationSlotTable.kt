package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The one-statement pre-event registration lock. `SsoConnection` remains
 * event truth; this row only decides which command may append that truth.
 * A slot whose projected connection is terminal may be atomically replaced.
 */
object SsoConnectionRegistrationSlotTable : CompositeIdTable("\"SsoConnectionRegistrationSlot\"") {
    val organizationId = text("organizationId").entityId()

    /** legacy | direct */
    val kind = text("kind").entityId()
    val connectionId = text("connectionId").uniqueIndex("\"SsoConnectionRegistrationSlot_connectionId_key\"")
    val replacesConnectionId = text("replacesConnectionId").nullable()
    val commandId = text("commandId")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(organizationId, kind, name = "\"SsoConnectionRegistrationSlot_pkey\"")
}
