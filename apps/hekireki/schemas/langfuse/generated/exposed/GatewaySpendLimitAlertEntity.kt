package langfuse.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** One notification per limit, subject, alert type, threshold, counting start and window, so the gateway may report a limit any number of times. */
class GatewaySpendLimitAlertEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, GatewaySpendLimitAlertEntity>(GatewaySpendLimitAlertTable)

    var spendLimitId by GatewaySpendLimitAlertTable.spendLimitId
    var spendLimit by GatewaySpendLimitEntity referencedOn GatewaySpendLimitAlertTable.spendLimitId
    var subjectId by GatewaySpendLimitAlertTable.subjectId
    var type by GatewaySpendLimitAlertTable.type
    var thresholdPercent by GatewaySpendLimitAlertTable.thresholdPercent
    var countingStartedAt by GatewaySpendLimitAlertTable.countingStartedAt
    var windowStart by GatewaySpendLimitAlertTable.windowStart
    var createdAt by GatewaySpendLimitAlertTable.createdAt
}
