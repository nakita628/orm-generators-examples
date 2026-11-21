package shelf.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A person's push notification switches. A missing row means every switch
 * is on, so no backfill is needed. The three category columns are named
 * after `PushCategory` in apps/webapp/app/modules/push/types.ts.
 */
object NotificationPreferenceTable : IdTable<String>("\"NotificationPreference\"") {
    override val id =
        text("userId")
            .references(
                UserTable.id,
                onDelete = ReferenceOption.CASCADE,
                onUpdate = ReferenceOption.CASCADE,
                fkName = "\"NotificationPreference_userId_fkey\"",
            ).entityId()
    val pushEnabled = bool("pushEnabled").default(true)
    val bookings = bool("bookings").default(true)
    val audits = bool("audits").default(true)
    val stockAndReminders = bool("stockAndReminders").default(true)
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"NotificationPreference_pkey\"")
}
