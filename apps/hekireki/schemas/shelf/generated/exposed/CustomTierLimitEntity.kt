package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class CustomTierLimitEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, CustomTierLimitEntity>(CustomTierLimitTable)

    var user by UserEntity optionalReferencedOn CustomTierLimitTable.userId
    var userId by CustomTierLimitTable.userId
    var canImportAssets by CustomTierLimitTable.canImportAssets
    var canExportAssets by CustomTierLimitTable.canExportAssets
    var canImportNRM by CustomTierLimitTable.canImportNRM
    var canHideShelfBranding by CustomTierLimitTable.canHideShelfBranding
    var canUseReports by CustomTierLimitTable.canUseReports
    var canUseRecurringReminders by CustomTierLimitTable.canUseRecurringReminders
    var canUsePushNotifications by CustomTierLimitTable.canUsePushNotifications
    var maxCustomFields by CustomTierLimitTable.maxCustomFields
    var maxOrganizations by CustomTierLimitTable.maxOrganizations
    var isEnterprise by CustomTierLimitTable.isEnterprise
    var createdAt by CustomTierLimitTable.createdAt
    var updatedAt by CustomTierLimitTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(CustomTierLimitTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
