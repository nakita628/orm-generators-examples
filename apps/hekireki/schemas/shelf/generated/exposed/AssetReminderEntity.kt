package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AssetReminderEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AssetReminderEntity>(AssetReminderTable)

    var name by AssetReminderTable.name
    var message by AssetReminderTable.message
    var alertDateTime by AssetReminderTable.alertDateTime
    var activeSchedulerReference by AssetReminderTable.activeSchedulerReference
    var recurrenceUnit by AssetReminderTable.recurrenceUnit
    var recurrenceInterval by AssetReminderTable.recurrenceInterval
    var recurrenceTimezone by AssetReminderTable.recurrenceTimezone
    var recurrenceEndsAt by AssetReminderTable.recurrenceEndsAt
    var organization by OrganizationEntity referencedOn AssetReminderTable.organizationId
    var organizationId by AssetReminderTable.organizationId
    var asset by AssetEntity referencedOn AssetReminderTable.assetId
    var assetId by AssetReminderTable.assetId
    var teamMembers by TeamMemberEntity via AssetReminderToTeamMemberTable
    var createdBy by UserEntity referencedOn AssetReminderTable.createdById
    var createdById by AssetReminderTable.createdById
    var createdAt by AssetReminderTable.createdAt
    var updatedAt by AssetReminderTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AssetReminderTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
