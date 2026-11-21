package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class WorkingHoursOverrideEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WorkingHoursOverrideEntity>(WorkingHoursOverrideTable)

    var date by WorkingHoursOverrideTable.date
    var isOpen by WorkingHoursOverrideTable.isOpen
    var openTime by WorkingHoursOverrideTable.openTime
    var closeTime by WorkingHoursOverrideTable.closeTime
    var reason by WorkingHoursOverrideTable.reason
    var workingHours by WorkingHoursEntity referencedOn WorkingHoursOverrideTable.workingHoursId
    var workingHoursId by WorkingHoursOverrideTable.workingHoursId
    var createdAt by WorkingHoursOverrideTable.createdAt
    var updatedAt by WorkingHoursOverrideTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(WorkingHoursOverrideTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
