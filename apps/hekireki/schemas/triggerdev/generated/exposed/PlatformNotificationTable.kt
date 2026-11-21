package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/** Admin-created notification definitions */
object PlatformNotificationTable : IdTable<String>("\"PlatformNotification\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val friendlyId =
        text("friendlyId").clientDefault { CUID.randomCUID1().toString() }.uniqueIndex("\"PlatformNotification_friendlyId_key\"")

    /** Admin-facing title for identification in admin tools */
    val title = text("title")

    /** Versioned JSON rendering content (see payload schema in spec) */
    val payload = jsonb<String>("payload", { it }, { it })
    val surface =
        pgEnum("surface", "\"PlatformNotificationSurface\"", PlatformNotificationSurface.entries, PlatformNotificationSurface::dbName)
    val scope = pgEnum("scope", "\"PlatformNotificationScope\"", PlatformNotificationScope.entries, PlatformNotificationScope::dbName)
    val userId =
        optReference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PlatformNotification_userId_fkey\"",
        )
    val organizationId =
        optReference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PlatformNotification_organizationId_fkey\"",
        )
    val projectId =
        optReference(
            "projectId",
            ProjectTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PlatformNotification_projectId_fkey\"",
        )

    /** When notification becomes active */
    val startsAt = pgTimestamp("startsAt", 3)

    /** When notification expires */
    val endsAt = pgTimestamp("endsAt", 3)

    /** CLI: auto-expire N days after user first saw it */
    val cliMaxDaysAfterFirstSeen = integer("cliMaxDaysAfterFirstSeen").nullable()

    /** CLI: auto-expire after shown N times to user */
    val cliMaxShowCount = integer("cliMaxShowCount").nullable()

    /** CLI: only show every N-th request (e.g. cliShowEvery=3 means show on 3rd, 6th, 9th…) */
    val cliShowEvery = integer("cliShowEvery").nullable()

    /** Ordering within same scope level (higher = more important) */
    val priority = integer("priority").default(0)

    /**
     * Staged-but-unpublished. While true the notification is hidden from all
     * user-facing reads regardless of startsAt/endsAt; publishing sets real
     * dates and flips this to false.
     */
    val isDraft = bool("isDraft").default(false)

    /** Soft delete */
    val archivedAt = pgTimestamp("archivedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"PlatformNotification_pkey\"")

    init {
        index("\"PlatformNotification_surface_scope_startsAt_idx\"", false, surface, scope, startsAt)
    }
}
