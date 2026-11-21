package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object TaskQueueTable : IdTable<String>("\"TaskQueue\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val friendlyId = text("friendlyId").uniqueIndex("\"TaskQueue_friendlyId_key\"")
    val name = text("name")
    val type = pgEnum("type", "\"TaskQueueType\"", TaskQueueType.entries, TaskQueueType::dbName).default(TaskQueueType.VIRTUAL)
    val version = pgEnum("version", "\"TaskQueueVersion\"", TaskQueueVersion.entries, TaskQueueVersion::dbName).default(TaskQueueVersion.V1)
    val orderableName = text("orderableName").nullable()

    /**
     * How the queue's concurrency was declared: V1 = legacy concurrencyLimit semantics
     * (per key when keyed, whole queue otherwise), V2 = the explicit perKey/total shape.
     */
    val concurrencyVersion =
        pgEnum(
            "concurrencyVersion",
            "\"TaskQueueConcurrencyVersion\"",
            TaskQueueConcurrencyVersion.entries,
            TaskQueueConcurrencyVersion::dbName,
        ).default(TaskQueueConcurrencyVersion.V1)

    /**
     * QUEUE rows are real queues; LIMIT rows back named concurrency limits and are
     * excluded from every queue-facing surface.
     */
    val role = pgEnum("role", "\"TaskQueueRole\"", TaskQueueRole.entries, TaskQueueRole::dbName).default(TaskQueueRole.QUEUE)
    val projectId =
        reference(
            "projectId",
            ProjectTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TaskQueue_projectId_fkey\"",
        )
    val runtimeEnvironmentId =
        reference(
            "runtimeEnvironmentId",
            RuntimeEnvironmentTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TaskQueue_runtimeEnvironmentId_fkey\"",
        )

    /** Represents the current concurrency limit for the queue */
    val concurrencyLimit = integer("concurrencyLimit").nullable()

    /** When the concurrency limit was overridden */
    val concurrencyLimitOverriddenAt = pgTimestamp("concurrencyLimitOverriddenAt", 3).nullable()

    /** Who overrode the concurrency limit (will be null if overridden via the API) */
    val concurrencyLimitOverriddenBy = text("concurrencyLimitOverriddenBy").nullable()

    /** If concurrencyLimit is overridden, this is the overridden value */
    val concurrencyLimitBase = integer("concurrencyLimitBase").nullable()

    /**
     * If the override was expressed as a percentage of the environment limit, this stores that
     * percentage (the source of truth). The absolute concurrencyLimit is materialized from it.
     * Decimal(5,2) allows fractional percentages like 12.50% (0.01–100.00).
     */
    val concurrencyLimitOverridePercent = decimal("concurrencyLimitOverridePercent", 5, 2).nullable()

    /**
     * Caps total concurrent runs across ALL concurrencyKey values of this queue
     * (concurrencyLimit applies per key value). Null = no total cap.
     */
    val totalConcurrencyLimit = integer("totalConcurrencyLimit").nullable()

    /** When the total concurrency limit was overridden */
    val totalConcurrencyLimitOverriddenAt = pgTimestamp("totalConcurrencyLimitOverriddenAt", 3).nullable()

    /** Who overrode the total concurrency limit (null when overridden via the API) */
    val totalConcurrencyLimitOverriddenBy = text("totalConcurrencyLimitOverriddenBy").nullable()

    /** If totalConcurrencyLimit is overridden, the declared value it reverts to on reset */
    val totalConcurrencyLimitBase = integer("totalConcurrencyLimitBase").nullable()
    val rateLimit = jsonb<String>("rateLimit", { it }, { it }).nullable()
    val paused = bool("paused").default(false)

    /** Hidden from the dashboard queue list. Display-only: never read on trigger/dequeue paths. */
    val archivedAt = pgTimestamp("archivedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"TaskQueue_pkey\"")

    init {
        uniqueIndex("\"TaskQueue_runtimeEnvironmentId_name_key\"", runtimeEnvironmentId, name)
    }
}
