package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

object RetentionApplyStatusTable : IdTable<String>("\"RetentionApplyStatus\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")
    val scopeType = pgEnum("scopeType", "\"RetentionPolicyScopeType\"", RetentionPolicyScopeType.entries, RetentionPolicyScopeType::dbName)
    val scopeId = text("scopeId")
    val changedAt = double("changedAt").nullable()
    val appliedAt = double("appliedAt").nullable()
    val raisedAt = double("raisedAt").nullable()
    val shortenedAt = double("shortenedAt").nullable()
    val createdAt = double("createdAt")
    val updatedAt = double("updatedAt")
    val occurredAt = double("occurredAt")
    val acceptedAt = double("acceptedAt")
    val lastEventId = text("lastEventId")
    val projectionVersion = text("projectionVersion")
    override val primaryKey = PrimaryKey(id, name = "\"RetentionApplyStatus_pkey\"")

    init {
        uniqueIndex("\"RetentionApplyStatus_organizationId_scopeType_scopeId_key\"", organizationId, scopeType, scopeId)
    }
}
