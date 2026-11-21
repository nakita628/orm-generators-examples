package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/** A user-defined metrics dashboard */
class MetricsDashboardEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, MetricsDashboardEntity>(MetricsDashboardTable)

    var friendlyId by MetricsDashboardTable.friendlyId
    var title by MetricsDashboardTable.title
    var description by MetricsDashboardTable.description
    var organization by OrganizationEntity referencedOn MetricsDashboardTable.organizationId
    var organizationId by MetricsDashboardTable.organizationId
    var project by ProjectEntity optionalReferencedOn MetricsDashboardTable.projectId
    var projectId by MetricsDashboardTable.projectId

    /** Who created the dashboard */
    var owner by UserEntity optionalReferencedOn MetricsDashboardTable.ownerId
    var ownerId by MetricsDashboardTable.ownerId
    var createdAt by MetricsDashboardTable.createdAt
    var updatedAt by MetricsDashboardTable.updatedAt

    /**
     * JSON that defines the config, queries, layout and config of all widgets.
     * There will be a version field for the format.
     */
    var layout by MetricsDashboardTable.layout

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(MetricsDashboardTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
