package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class TierLimitEntity(
    id: EntityID<TierId>,
) : Entity<TierId>(id) {
    companion object : EntityClass<TierId, TierLimitEntity>(TierLimitTable)

    val tier by TierEntity optionalBackReferencedOn TierTable.tierLimitId
    var canImportAssets by TierLimitTable.canImportAssets
    var canExportAssets by TierLimitTable.canExportAssets
    var canImportNRM by TierLimitTable.canImportNRM
    var canHideShelfBranding by TierLimitTable.canHideShelfBranding
    var canUseReports by TierLimitTable.canUseReports
    var canUseRecurringReminders by TierLimitTable.canUseRecurringReminders
    var canUsePushNotifications by TierLimitTable.canUsePushNotifications
    var maxCustomFields by TierLimitTable.maxCustomFields
    var maxOrganizations by TierLimitTable.maxOrganizations
    var createdAt by TierLimitTable.createdAt
    var updatedAt by TierLimitTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(TierLimitTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
