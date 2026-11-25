package teable.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class FieldPeriodicRefreshStateEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, FieldPeriodicRefreshStateEntity>(FieldPeriodicRefreshStateTable)

    var tableId by FieldPeriodicRefreshStateTable.tableId
    var baseId by FieldPeriodicRefreshStateTable.baseId
    var enabled by FieldPeriodicRefreshStateTable.enabled
    var timeZone by FieldPeriodicRefreshStateTable.timeZone
    var containsNowPrecision by FieldPeriodicRefreshStateTable.containsNowPrecision
    var lastRefreshedAt by FieldPeriodicRefreshStateTable.lastRefreshedAt
    var lastRefreshedLocalDate by FieldPeriodicRefreshStateTable.lastRefreshedLocalDate
    var lastRangeHash by FieldPeriodicRefreshStateTable.lastRangeHash
    var lastStatus by FieldPeriodicRefreshStateTable.lastStatus
    var lastError by FieldPeriodicRefreshStateTable.lastError
    var pendingTaskId by FieldPeriodicRefreshStateTable.pendingTaskId
    var pendingLocalDate by FieldPeriodicRefreshStateTable.pendingLocalDate
    var pendingRangeHash by FieldPeriodicRefreshStateTable.pendingRangeHash
    var pendingStartedAt by FieldPeriodicRefreshStateTable.pendingStartedAt
    var createdTime by FieldPeriodicRefreshStateTable.createdTime
    var lastModifiedTime by FieldPeriodicRefreshStateTable.lastModifiedTime
    val field by FieldEntity referencedOn FieldPeriodicRefreshStateTable.id

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(FieldPeriodicRefreshStateTable.lastModifiedTime)) {
                lastModifiedTime = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
