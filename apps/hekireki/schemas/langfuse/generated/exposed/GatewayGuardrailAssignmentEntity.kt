package langfuse.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class GatewayGuardrailAssignmentEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, GatewayGuardrailAssignmentEntity>(GatewayGuardrailAssignmentTable)

    var guardrailId by GatewayGuardrailAssignmentTable.guardrailId
    var guardrail by GatewayGuardrailEntity referencedOn GatewayGuardrailAssignmentTable.guardrailId
    var organizationId by GatewayGuardrailAssignmentTable.organizationId
    var organization by OrganizationEntity referencedOn GatewayGuardrailAssignmentTable.organizationId
    var scope by GatewayGuardrailAssignmentTable.scope
    var apiKeyId by GatewayGuardrailAssignmentTable.apiKeyId
    var apiKey by ApiKeyEntity optionalReferencedOn GatewayGuardrailAssignmentTable.apiKeyId
    var userId by GatewayGuardrailAssignmentTable.userId
    var user by UserEntity optionalReferencedOn GatewayGuardrailAssignmentTable.userId
    var createdAt by GatewayGuardrailAssignmentTable.createdAt
}
