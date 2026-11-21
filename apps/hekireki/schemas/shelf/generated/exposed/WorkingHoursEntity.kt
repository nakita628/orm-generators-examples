package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class WorkingHoursEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WorkingHoursEntity>(WorkingHoursTable)

    var enabled by WorkingHoursTable.enabled
    var weeklySchedule by WorkingHoursTable.weeklySchedule
    var organization by OrganizationEntity referencedOn WorkingHoursTable.organizationId
    var organizationId by WorkingHoursTable.organizationId
    val overrides by WorkingHoursOverrideEntity referrersOn WorkingHoursOverrideTable.workingHoursId
    var createdAt by WorkingHoursTable.createdAt
    var updatedAt by WorkingHoursTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(WorkingHoursTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
