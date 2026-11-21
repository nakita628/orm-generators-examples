package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/** Schedules can be attached to tasks to trigger them on a schedule */
class TaskScheduleEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TaskScheduleEntity>(TaskScheduleTable)

    var type by TaskScheduleTable.type

    /** users see this as `id`. They start with schedule_ */
    var friendlyId by TaskScheduleTable.friendlyId

    /** a reference to a task (not a foreign key because it's across versions) */
    var taskIdentifier by TaskScheduleTable.taskIdentifier

    /** can be provided and we won't create another with the same key */
    var deduplicationKey by TaskScheduleTable.deduplicationKey
    var userProvidedDeduplicationKey by TaskScheduleTable.userProvidedDeduplicationKey

    /** the CRON pattern */
    var generatorExpression by TaskScheduleTable.generatorExpression
    var generatorDescription by TaskScheduleTable.generatorDescription
    var generatorType by TaskScheduleTable.generatorType

    /** These are IANA format string, or the default "UTC". E.g. "America/New_York" */
    var timezone by TaskScheduleTable.timezone
    var windowDurationSeconds by TaskScheduleTable.windowDurationSeconds
    var windowPercentage by TaskScheduleTable.windowPercentage
    var defaultWindowDurationSeconds by TaskScheduleTable.defaultWindowDurationSeconds
    var minimumWindowDurationSeconds by TaskScheduleTable.minimumWindowDurationSeconds

    /** Can be provided by the user then accessed inside a run */
    var externalId by TaskScheduleTable.externalId

    /** Instances of the schedule that are active */
    val instances by TaskScheduleInstanceEntity referrersOn TaskScheduleInstanceTable.taskScheduleId

    /** @deprecated stop writing 2026-04-30; reads moved out of code (UI now derives from cron's previous slot). Drop in follow-up. */
    var lastRunTriggeredAt by TaskScheduleTable.lastRunTriggeredAt
    var project by ProjectEntity referencedOn TaskScheduleTable.projectId
    var projectId by TaskScheduleTable.projectId
    var createdAt by TaskScheduleTable.createdAt
    var updatedAt by TaskScheduleTable.updatedAt
    var active by TaskScheduleTable.active

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(TaskScheduleTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
