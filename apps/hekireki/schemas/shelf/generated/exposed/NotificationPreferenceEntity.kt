package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A person's push notification switches. A missing row means every switch
 * is on, so no backfill is needed. The three category columns are named
 * after `PushCategory` in apps/webapp/app/modules/push/types.ts.
 */
class NotificationPreferenceEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, NotificationPreferenceEntity>(NotificationPreferenceTable)

    val user by UserEntity referencedOn NotificationPreferenceTable.id
    var pushEnabled by NotificationPreferenceTable.pushEnabled
    var bookings by NotificationPreferenceTable.bookings
    var audits by NotificationPreferenceTable.audits
    var stockAndReminders by NotificationPreferenceTable.stockAndReminders
    var updatedAt by NotificationPreferenceTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(NotificationPreferenceTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
