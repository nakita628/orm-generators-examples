package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class GatewayA2ATargetEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, GatewayA2ATargetEntity>(GatewayA2ATargetTable)

    var organizationId by GatewayA2ATargetTable.organizationId
    var projectId by GatewayA2ATargetTable.projectId
    var name by GatewayA2ATargetTable.name
    var agentName by GatewayA2ATargetTable.agentName
    var cardUrl by GatewayA2ATargetTable.cardUrl
    var upstreamUrl by GatewayA2ATargetTable.upstreamUrl
    var encryptedUpstreamHeaders by GatewayA2ATargetTable.encryptedUpstreamHeaders
    var restrictToAllowList by GatewayA2ATargetTable.restrictToAllowList
    var allowedVirtualKeyIds by GatewayA2ATargetTable.allowedVirtualKeyIds
    var blocked by GatewayA2ATargetTable.blocked
    var blockReason by GatewayA2ATargetTable.blockReason
    var createdById by GatewayA2ATargetTable.createdById
    var createdAt by GatewayA2ATargetTable.createdAt
    var updatedAt by GatewayA2ATargetTable.updatedAt
    var archivedAt by GatewayA2ATargetTable.archivedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(GatewayA2ATargetTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
