package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One sweep of one policy for one organisation: the History tab's rows. `finishedAt` stays null
 * while the run is in progress, or if it never finished.
 */
class RetentionRunEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, RetentionRunEntity>(RetentionRunTable)

    var organization by OrganizationEntity referencedOn RetentionRunTable.organizationId
    var organizationId by RetentionRunTable.organizationId
    var entity by RetentionRunTable.entity
    var kind by RetentionRunTable.kind
    var startedAt by RetentionRunTable.startedAt
    var finishedAt by RetentionRunTable.finishedAt
    var notifiedCount by RetentionRunTable.notifiedCount

    /** Archived surveys, or deactivated members. */
    var archivedCount by RetentionRunTable.archivedCount

    /** Deleted responses, or, on a purge run, deleted surveys. */
    var deletedCount by RetentionRunTable.deletedCount
    var skippedCount by RetentionRunTable.skippedCount

    /** False when the run changed nothing. History hides those by default. */
    var hasChanges by RetentionRunTable.hasChanges

    /**
     * Where the run's scan of candidates stopped for want of time, so the next run resumes there instead
     * of rescanning the same first targets every night. Null when the scan reached the end.
     */
    var scanCursor by RetentionRunTable.scanCursor
    val items by RetentionRunItemEntity referrersOn RetentionRunItemTable.runId
}
