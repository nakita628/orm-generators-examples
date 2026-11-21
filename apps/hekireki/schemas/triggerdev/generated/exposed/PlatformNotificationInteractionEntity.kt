package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/** Per-user tracking of notification views and dismissals */
class PlatformNotificationInteractionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, PlatformNotificationInteractionEntity>(PlatformNotificationInteractionTable)

    var notification by PlatformNotificationEntity referencedOn PlatformNotificationInteractionTable.notificationId
    var notificationId by PlatformNotificationInteractionTable.notificationId
    var user by UserEntity referencedOn PlatformNotificationInteractionTable.userId
    var userId by PlatformNotificationInteractionTable.userId

    /** Set by beacon/CLI GET on first impression */
    var firstSeenAt by PlatformNotificationInteractionTable.firstSeenAt

    /** Times shown (incremented per beacon/CLI view) */
    var showCount by PlatformNotificationInteractionTable.showCount

    /** User dismissed in webapp */
    var webappDismissedAt by PlatformNotificationInteractionTable.webappDismissedAt

    /** User clicked a link in webapp */
    var webappClickedAt by PlatformNotificationInteractionTable.webappClickedAt

    /** Auto-dismissed or expired in CLI */
    var cliDismissedAt by PlatformNotificationInteractionTable.cliDismissedAt
    var createdAt by PlatformNotificationInteractionTable.createdAt
    var updatedAt by PlatformNotificationInteractionTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(PlatformNotificationInteractionTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
