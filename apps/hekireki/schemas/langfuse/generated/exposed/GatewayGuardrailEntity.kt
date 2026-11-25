package langfuse.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class GatewayGuardrailEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, GatewayGuardrailEntity>(GatewayGuardrailTable)

    var organizationId by GatewayGuardrailTable.organizationId
    var organization by OrganizationEntity referencedOn GatewayGuardrailTable.organizationId
    var name by GatewayGuardrailTable.name
    var type by GatewayGuardrailTable.type
    var action by GatewayGuardrailTable.action
    var config by GatewayGuardrailTable.config
    var encryptedCredential by GatewayGuardrailTable.encryptedCredential
    var createdBy by GatewayGuardrailTable.createdBy
    var createdByUser by UserEntity optionalReferencedOn GatewayGuardrailTable.createdBy
    var createdAt by GatewayGuardrailTable.createdAt
    var updatedAt by GatewayGuardrailTable.updatedAt
    val assignments by GatewayGuardrailAssignmentEntity referrersOn GatewayGuardrailAssignmentTable.guardrailId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(GatewayGuardrailTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
