package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A reason the environment variables "Needs attention" page shows a key in one environment, which
 * the project's team chose to ignore. An ignored EMPTY goes when the value is next saved.
 */
object EnvironmentVariableIgnoreTable : IdTable<String>("\"EnvironmentVariableIgnore\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val projectId =
        reference(
            "projectId",
            ProjectTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"EnvironmentVariableIgnore_projectId_fkey\"",
        )
    val environmentId =
        reference(
            "environmentId",
            RuntimeEnvironmentTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"EnvironmentVariableIgnore_environmentId_fkey\"",
        )
    val variableId =
        reference(
            "variableId",
            EnvironmentVariableTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"EnvironmentVariableIgnore_variableId_fkey\"",
        )
    val reason =
        pgEnum(
            "reason",
            "\"EnvironmentVariableIgnoreReason\"",
            EnvironmentVariableIgnoreReason.entries,
            EnvironmentVariableIgnoreReason::dbName,
        )
    val createdById =
        optReference(
            "createdById",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"EnvironmentVariableIgnore_createdById_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"EnvironmentVariableIgnore_pkey\"")

    init {
        index("\"EnvironmentVariableIgnore_variableId_idx\"", false, variableId)
        index("\"EnvironmentVariableIgnore_projectId_idx\"", false, projectId)
        index("\"EnvironmentVariableIgnore_createdById_idx\"", false, createdById)
        uniqueIndex("\"EnvironmentVariableIgnore_environmentId_variableId_reason_key\"", environmentId, variableId, reason)
    }
}
