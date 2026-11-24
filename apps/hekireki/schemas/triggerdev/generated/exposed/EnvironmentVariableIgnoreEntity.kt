package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * A reason the environment variables "Needs attention" page shows a key in one environment, which
 * the project's team chose to ignore. An ignored EMPTY goes when the value is next saved.
 */
class EnvironmentVariableIgnoreEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, EnvironmentVariableIgnoreEntity>(EnvironmentVariableIgnoreTable)

    var project by ProjectEntity referencedOn EnvironmentVariableIgnoreTable.projectId
    var projectId by EnvironmentVariableIgnoreTable.projectId
    var environment by RuntimeEnvironmentEntity referencedOn EnvironmentVariableIgnoreTable.environmentId
    var environmentId by EnvironmentVariableIgnoreTable.environmentId

    /** For MISSING, a variable of the project with no value in this environment. */
    var variable by EnvironmentVariableEntity referencedOn EnvironmentVariableIgnoreTable.variableId
    var variableId by EnvironmentVariableIgnoreTable.variableId
    var reason by EnvironmentVariableIgnoreTable.reason
    var createdBy by UserEntity optionalReferencedOn EnvironmentVariableIgnoreTable.createdById
    var createdById by EnvironmentVariableIgnoreTable.createdById
    var createdAt by EnvironmentVariableIgnoreTable.createdAt
}
