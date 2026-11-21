package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class BulkActionItemEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, BulkActionItemEntity>(BulkActionItemTable)

    /** @deprecated not used in new BulkActions */
    var friendlyId by BulkActionItemTable.friendlyId
    var groupId by BulkActionItemTable.groupId
    var type by BulkActionItemTable.type

    /** When the item is created it's pending. After we've processed the item it's completed. This does not mean the associated runs are completed. */
    var status by BulkActionItemTable.status

    /** The run that is the source of the action, e.g. when replaying this is the original run */
    var sourceRun by TaskRunEntity referencedOn BulkActionItemTable.sourceRunId
    var sourceRunId by BulkActionItemTable.sourceRunId

    /** The run that's a result of the action, this will be set when the run has been created */
    var destinationRun by TaskRunEntity optionalReferencedOn BulkActionItemTable.destinationRunId
    var destinationRunId by BulkActionItemTable.destinationRunId
    var error by BulkActionItemTable.error
    var createdAt by BulkActionItemTable.createdAt
    var updatedAt by BulkActionItemTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(BulkActionItemTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
