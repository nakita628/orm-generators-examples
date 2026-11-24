package rallly.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Monthly counters for metered actions charged to a space, independent of
 * the records they produce, so deleting a poll does not give back what it
 * spent.
 */
class SpaceUsageEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<SpaceUsageEntity>(SpaceUsageTable)

    val spaceId by SpaceUsageTable.spaceId
    val resource by SpaceUsageTable.resource
    val period by SpaceUsageTable.period
    var count by SpaceUsageTable.count
    var nearLimitWarnedAt by SpaceUsageTable.nearLimitWarnedAt
    var atLimitWarnedAt by SpaceUsageTable.atLimitWarnedAt
    var updatedAt by SpaceUsageTable.updatedAt
    val space by SpaceEntity referencedOn SpaceUsageTable.spaceId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SpaceUsageTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
