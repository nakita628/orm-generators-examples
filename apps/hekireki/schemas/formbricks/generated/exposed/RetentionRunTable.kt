package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One sweep of one policy for one organisation: the History tab's rows. `finishedAt` stays null
 * while the run is in progress, or if it never finished.
 */
object RetentionRunTable : IdTable<String>("\"RetentionRun\"") {
    override val id = text("id").clientDefault { CUID.randomCUID2().toString() }.entityId()
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RetentionRun_organizationId_fkey\"",
        )
    val entity = pgEnum("entity", "\"RetentionEntity\"", RetentionEntity.entries, RetentionEntity::dbName)
    val kind = pgEnum("kind", "\"RetentionRunKind\"", RetentionRunKind.entries, RetentionRunKind::dbName).default(RetentionRunKind.SWEEP)
    val startedAt = pgTimestamp("startedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val finishedAt = pgTimestamp("finishedAt", 3).nullable()
    val notifiedCount = integer("notifiedCount").default(0)

    /** Archived surveys, or deactivated members. */
    val archivedCount = integer("archivedCount").default(0)

    /** Deleted responses, or, on a purge run, deleted surveys. */
    val deletedCount = integer("deletedCount").default(0)
    val skippedCount = integer("skippedCount").default(0)

    /** False when the run changed nothing. History hides those by default. */
    val hasChanges = bool("hasChanges").default(false)

    /**
     * Where the run's scan of candidates stopped for want of time, so the next run resumes there instead
     * of rescanning the same first targets every night. Null when the scan reached the end.
     */
    val scanCursor = text("scanCursor").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"RetentionRun_pkey\"")

    init {
        index("\"RetentionRun_organizationId_startedAt_id_idx\"", false, organizationId, startedAt, id)
    }
}
