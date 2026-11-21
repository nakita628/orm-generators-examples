package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/** Schedules can be attached to tasks to trigger them on a schedule */
object TaskScheduleTable : IdTable<String>("\"TaskSchedule\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val type = pgEnum("type", "\"ScheduleType\"", ScheduleType.entries, ScheduleType::dbName).default(ScheduleType.IMPERATIVE)

    /** users see this as `id`. They start with schedule_ */
    val friendlyId = text("friendlyId").uniqueIndex("\"TaskSchedule_friendlyId_key\"")

    /** a reference to a task (not a foreign key because it's across versions) */
    val taskIdentifier = text("taskIdentifier")

    /** can be provided and we won't create another with the same key */
    val deduplicationKey = text("deduplicationKey").clientDefault { CUID.randomCUID1().toString() }
    val userProvidedDeduplicationKey = bool("userProvidedDeduplicationKey").default(false)

    /** the CRON pattern */
    val generatorExpression = text("generatorExpression")
    val generatorDescription = text("generatorDescription").default("")
    val generatorType =
        pgEnum("generatorType", "\"ScheduleGeneratorType\"", ScheduleGeneratorType.entries, ScheduleGeneratorType::dbName).default(
            ScheduleGeneratorType.CRON,
        )

    /** These are IANA format string, or the default "UTC". E.g. "America/New_York" */
    val timezone = text("timezone").default("UTC")
    val windowDurationSeconds = integer("windowDurationSeconds").nullable()
    val windowPercentage = integer("windowPercentage").nullable()
    val defaultWindowDurationSeconds = integer("defaultWindowDurationSeconds").nullable()
    val minimumWindowDurationSeconds = integer("minimumWindowDurationSeconds").nullable()

    /** Can be provided by the user then accessed inside a run */
    val externalId = text("externalId").nullable()

    /** @deprecated stop writing 2026-04-30; reads moved out of code (UI now derives from cron's previous slot). Drop in follow-up. */
    val lastRunTriggeredAt = pgTimestamp("lastRunTriggeredAt", 3).nullable()
    val projectId =
        reference(
            "projectId",
            ProjectTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TaskSchedule_projectId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val active = bool("active").default(true)
    override val primaryKey = PrimaryKey(id, name = "\"TaskSchedule_pkey\"")

    init {
        index("\"TaskSchedule_projectId_idx\"", false, projectId)
        index(
            "\"TaskSchedule_projectId_createdAt_idx\"",
            false,
            functions = listOf(indexExpression("\"projectId\""), indexExpression("\"createdAt\" DESC")),
        )
        uniqueIndex("\"TaskSchedule_projectId_deduplicationKey_key\"", projectId, deduplicationKey)
    }
}
