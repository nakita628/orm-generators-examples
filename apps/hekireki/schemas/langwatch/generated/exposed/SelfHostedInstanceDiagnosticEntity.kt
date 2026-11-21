package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One upgrade diagnostics bundle a self-hosted install sent over Connect, on
 * LangWatch Cloud. The bytes live in object storage under
 * `upgrade-diagnostics/` and expire at 30 days by a lifecycle rule; this row
 * points at them and feeds Cloud admin and the Slack signal. Plain column
 * refs, no relation, as the usage report rows.
 */
class SelfHostedInstanceDiagnosticEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SelfHostedInstanceDiagnosticEntity>(SelfHostedInstanceDiagnosticTable)

    /** The `SelfHostedInstance.instanceId` the licence caller presented. */
    var instanceId by SelfHostedInstanceDiagnosticTable.instanceId
    var stepId by SelfHostedInstanceDiagnosticTable.stepId
    var stepDescription by SelfHostedInstanceDiagnosticTable.stepDescription
    var release by SelfHostedInstanceDiagnosticTable.release
    var attempt by SelfHostedInstanceDiagnosticTable.attempt

    /** `operator` or `automatic`. */
    var trigger by SelfHostedInstanceDiagnosticTable.trigger
    var byteLength by SelfHostedInstanceDiagnosticTable.byteLength
    var objectKey by SelfHostedInstanceDiagnosticTable.objectKey
    var receivedAt by SelfHostedInstanceDiagnosticTable.receivedAt

    /** When the lifecycle rule removes the bundle; the row is listed until then. */
    var expiresAt by SelfHostedInstanceDiagnosticTable.expiresAt

    /** Set on the report that posted this step and release's one Slack message. */
    var slackAnnouncedAt by SelfHostedInstanceDiagnosticTable.slackAnnouncedAt
}
