package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One usage report, exactly as it arrived. Kept so a number on an install's
 * row can be traced back to the report that set it, and so a question asked
 * later can be answered over the history rather than over the latest snapshot.
 */
class SelfHostedInstanceReportEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SelfHostedInstanceReportEntity>(SelfHostedInstanceReportTable)

    /**
     * The `SelfHostedInstance.instanceId` this report came from. A plain column
     * ref for the same reason the instance row's refs are.
     */
    var instanceId by SelfHostedInstanceReportTable.instanceId
    var receivedAt by SelfHostedInstanceReportTable.receivedAt
    var version by SelfHostedInstanceReportTable.version
    var reportSchemaVersion by SelfHostedInstanceReportTable.reportSchemaVersion

    /** Fields this release had no name for. Counted, never stored by name. */
    var unknownFields by SelfHostedInstanceReportTable.unknownFields

    /** The validated report, with every unknown field already dropped. */
    var payload by SelfHostedInstanceReportTable.payload
}
