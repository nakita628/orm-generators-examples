package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/** Bulk actions, like canceling and replaying runs */
class BulkActionGroupEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, BulkActionGroupEntity>(BulkActionGroupTable)

    var friendlyId by BulkActionGroupTable.friendlyId
    var project by ProjectEntity referencedOn BulkActionGroupTable.projectId
    var projectId by BulkActionGroupTable.projectId

    /** If this is set then it's a V2 Bulk Action that supports queries */
    var environment by RuntimeEnvironmentEntity optionalReferencedOn BulkActionGroupTable.environmentId
    var environmentId by BulkActionGroupTable.environmentId
    var type by BulkActionGroupTable.type

    /** When the group is created it's pending. After we've processed all the items it's completed. This does not mean the associated runs are completed. */
    var status by BulkActionGroupTable.status

    /** The query that will be executed to get the runs to process (if null, we are passing run ids directly) */
    var queryName by BulkActionGroupTable.queryName

    /** The params that will be passed to the query */
    var params by BulkActionGroupTable.params

    /** Billing-limit dedupe key (grace hitAt or resolve job key). Null for dashboard bulk actions. */
    var dedupeKey by BulkActionGroupTable.dedupeKey

    /** The cursor that will be passed to the query (null for the first page) */
    var cursor by BulkActionGroupTable.cursor

    /** The number of runs that have been processed successfully */
    var successCount by BulkActionGroupTable.successCount

    /** The number of runs that have been failed */
    var failureCount by BulkActionGroupTable.failureCount

    /** The total number of runs that will be processed */
    var totalCount by BulkActionGroupTable.totalCount
    var completionNotification by BulkActionGroupTable.completionNotification

    /** The userId who did the bulk action */
    var user by UserEntity optionalReferencedOn BulkActionGroupTable.userId
    var userId by BulkActionGroupTable.userId

    /** The name of the bulk action */
    var name by BulkActionGroupTable.name

    /** The time the bulk action was completed */
    var completedAt by BulkActionGroupTable.completedAt
    var createdAt by BulkActionGroupTable.createdAt
    var updatedAt by BulkActionGroupTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(BulkActionGroupTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
