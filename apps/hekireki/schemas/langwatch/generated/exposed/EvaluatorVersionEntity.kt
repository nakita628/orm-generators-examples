package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** One immutable snapshot of an evaluator; the Evaluator row is always the latest. */
class EvaluatorVersionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, EvaluatorVersionEntity>(EvaluatorVersionTable)

    var projectId by EvaluatorVersionTable.projectId
    var evaluatorId by EvaluatorVersionTable.evaluatorId
    var version by EvaluatorVersionTable.version
    var name by EvaluatorVersionTable.name
    var type by EvaluatorVersionTable.type
    var config by EvaluatorVersionTable.config
    var workflowId by EvaluatorVersionTable.workflowId
    var commitMessage by EvaluatorVersionTable.commitMessage
    var authorId by EvaluatorVersionTable.authorId

    /** "user", "api_key" or "system": who wrote the version when authorId is null. */
    var authorKind by EvaluatorVersionTable.authorKind
    var createdAt by EvaluatorVersionTable.createdAt
}
