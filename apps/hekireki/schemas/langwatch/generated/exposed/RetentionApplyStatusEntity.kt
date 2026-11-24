package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class RetentionApplyStatusEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, RetentionApplyStatusEntity>(RetentionApplyStatusTable)

    var organizationId by RetentionApplyStatusTable.organizationId
    var scopeType by RetentionApplyStatusTable.scopeType
    var scopeId by RetentionApplyStatusTable.scopeId
    var changedAt by RetentionApplyStatusTable.changedAt
    var appliedAt by RetentionApplyStatusTable.appliedAt
    var raisedAt by RetentionApplyStatusTable.raisedAt
    var shortenedAt by RetentionApplyStatusTable.shortenedAt
    var createdAt by RetentionApplyStatusTable.createdAt
    var updatedAt by RetentionApplyStatusTable.updatedAt
    var occurredAt by RetentionApplyStatusTable.occurredAt
    var acceptedAt by RetentionApplyStatusTable.acceptedAt
    var lastEventId by RetentionApplyStatusTable.lastEventId
    var projectionVersion by RetentionApplyStatusTable.projectionVersion
}
