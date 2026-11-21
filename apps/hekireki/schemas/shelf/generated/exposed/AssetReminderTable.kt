package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object AssetReminderTable : IdTable<String>("\"AssetReminder\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name")
    val message = text("message")
    val alertDateTime = pgTimestamp("alertDateTime", 3)
    val activeSchedulerReference = text("activeSchedulerReference").nullable()
    val recurrenceUnit =
        pgEnum("recurrenceUnit", "\"ReminderRecurrenceUnit\"", ReminderRecurrenceUnit.entries, ReminderRecurrenceUnit::dbName).nullable()
    val recurrenceInterval = integer("recurrenceInterval").nullable()
    val recurrenceTimezone = text("recurrenceTimezone").nullable()
    val recurrenceEndsAt = pgTimestamp("recurrenceEndsAt", 3).nullable()
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetReminder_organizationId_fkey\"",
        )
    val assetId =
        reference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetReminder_assetId_fkey\"",
        )
    val createdById =
        reference(
            "createdById",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetReminder_createdById_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AssetReminder_pkey\"")

    init {
        index("\"AssetReminder_assetId_alertDateTime_idx\"", false, assetId, alertDateTime)
        index(
            "\"AssetReminder_name_message_idx\"",
            false,
            functions = listOf(indexExpression("\"name\" gin_trgm_ops"), indexExpression("\"message\" gin_trgm_ops")),
            indexType = "gin",
        )
        index("\"AssetReminder_organizationId_alertDateTime_assetId_idx\"", false, organizationId, alertDateTime, assetId)
        index("\"AssetReminder_alertDateTime_activeSchedulerReference_idx\"", false, alertDateTime, activeSchedulerReference)
        index("\"AssetReminder_createdById_idx\"", false, createdById)
    }
}
