package langfuse.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class GatewaySpendLimitEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, GatewaySpendLimitEntity>(GatewaySpendLimitTable)

    var organizationId by GatewaySpendLimitTable.organizationId
    var organization by OrganizationEntity referencedOn GatewaySpendLimitTable.organizationId
    var scope by GatewaySpendLimitTable.scope
    var apiKeyId by GatewaySpendLimitTable.apiKeyId
    var apiKey by ApiKeyEntity optionalReferencedOn GatewaySpendLimitTable.apiKeyId
    var userId by GatewaySpendLimitTable.userId
    var user by UserEntity optionalReferencedOn GatewaySpendLimitTable.userId
    var interval by GatewaySpendLimitTable.interval
    var limitUsd by GatewaySpendLimitTable.limitUsd
    var alertThresholds by GatewaySpendLimitTable.alertThresholds
    var countingStartedAt by GatewaySpendLimitTable.countingStartedAt
    var ownerUserId by GatewaySpendLimitTable.ownerUserId
    var ownerUser by UserEntity optionalReferencedOn GatewaySpendLimitTable.ownerUserId
    var createdAt by GatewaySpendLimitTable.createdAt
    var updatedAt by GatewaySpendLimitTable.updatedAt
    val alerts by GatewaySpendLimitAlertEntity referrersOn GatewaySpendLimitAlertTable.spendLimitId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(GatewaySpendLimitTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
