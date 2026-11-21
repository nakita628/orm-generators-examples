package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/** Admin-created notification definitions */
class PlatformNotificationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, PlatformNotificationEntity>(PlatformNotificationTable)

    var friendlyId by PlatformNotificationTable.friendlyId

    /** Admin-facing title for identification in admin tools */
    var title by PlatformNotificationTable.title

    /** Versioned JSON rendering content (see payload schema in spec) */
    var payload by PlatformNotificationTable.payload
    var surface by PlatformNotificationTable.surface
    var scope by PlatformNotificationTable.scope

    /** Set when scope = USER */
    var user by UserEntity optionalReferencedOn PlatformNotificationTable.userId
    var userId by PlatformNotificationTable.userId

    /** Set when scope = ORGANIZATION */
    var organization by OrganizationEntity optionalReferencedOn PlatformNotificationTable.organizationId
    var organizationId by PlatformNotificationTable.organizationId

    /** Set when scope = PROJECT */
    var project by ProjectEntity optionalReferencedOn PlatformNotificationTable.projectId
    var projectId by PlatformNotificationTable.projectId

    /** When notification becomes active */
    var startsAt by PlatformNotificationTable.startsAt

    /** When notification expires */
    var endsAt by PlatformNotificationTable.endsAt

    /** CLI: auto-expire N days after user first saw it */
    var cliMaxDaysAfterFirstSeen by PlatformNotificationTable.cliMaxDaysAfterFirstSeen

    /** CLI: auto-expire after shown N times to user */
    var cliMaxShowCount by PlatformNotificationTable.cliMaxShowCount

    /** CLI: only show every N-th request (e.g. cliShowEvery=3 means show on 3rd, 6th, 9th…) */
    var cliShowEvery by PlatformNotificationTable.cliShowEvery

    /** Ordering within same scope level (higher = more important) */
    var priority by PlatformNotificationTable.priority

    /**
     * Staged-but-unpublished. While true the notification is hidden from all
     * user-facing reads regardless of startsAt/endsAt; publishing sets real
     * dates and flips this to false.
     */
    var isDraft by PlatformNotificationTable.isDraft

    /** Soft delete */
    var archivedAt by PlatformNotificationTable.archivedAt
    var createdAt by PlatformNotificationTable.createdAt
    var updatedAt by PlatformNotificationTable.updatedAt
    val interactions by PlatformNotificationInteractionEntity referrersOn PlatformNotificationInteractionTable.notificationId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(PlatformNotificationTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
