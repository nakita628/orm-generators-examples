package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** One immutable snapshot of an agent; the Agent row is always the latest. */
class AgentVersionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AgentVersionEntity>(AgentVersionTable)

    var projectId by AgentVersionTable.projectId
    var agentId by AgentVersionTable.agentId
    var version by AgentVersionTable.version
    var name by AgentVersionTable.name
    var type by AgentVersionTable.type
    var config by AgentVersionTable.config
    var workflowId by AgentVersionTable.workflowId
    var commitMessage by AgentVersionTable.commitMessage
    var authorId by AgentVersionTable.authorId

    /** "user", "api_key" or "system": who wrote the version when authorId is null. */
    var authorKind by AgentVersionTable.authorKind
    var createdAt by AgentVersionTable.createdAt
}
