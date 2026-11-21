package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/** Per-user tracking of notification views and dismissals */
object PlatformNotificationInteractionTable : IdTable<String>("\"PlatformNotificationInteraction\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val notificationId =
        reference(
            "notificationId",
            PlatformNotificationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PlatformNotificationInteraction_notificationId_fkey\"",
        )
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PlatformNotificationInteraction_userId_fkey\"",
        )

    /** Set by beacon/CLI GET on first impression */
    val firstSeenAt = pgTimestamp("firstSeenAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /** Times shown (incremented per beacon/CLI view) */
    val showCount = integer("showCount").default(1)

    /** User dismissed in webapp */
    val webappDismissedAt = pgTimestamp("webappDismissedAt", 3).nullable()

    /** User clicked a link in webapp */
    val webappClickedAt = pgTimestamp("webappClickedAt", 3).nullable()

    /** Auto-dismissed or expired in CLI */
    val cliDismissedAt = pgTimestamp("cliDismissedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"PlatformNotificationInteraction_pkey\"")

    init {
        uniqueIndex("\"PlatformNotificationInteraction_notificationId_userId_key\"", notificationId, userId)
    }
}
