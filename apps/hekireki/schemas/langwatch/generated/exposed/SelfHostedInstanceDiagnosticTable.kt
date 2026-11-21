package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One upgrade diagnostics bundle a self-hosted install sent over Connect, on
 * LangWatch Cloud. The bytes live in object storage under
 * `upgrade-diagnostics/` and expire at 30 days by a lifecycle rule; this row
 * points at them and feeds Cloud admin and the Slack signal. Plain column
 * refs, no relation, as the usage report rows.
 */
object SelfHostedInstanceDiagnosticTable : IdTable<String>("\"SelfHostedInstanceDiagnostic\"") {
    override val id = text("id").entityId()

    /** The `SelfHostedInstance.instanceId` the licence caller presented. */
    val instanceId = text("instanceId")
    val stepId = text("stepId")
    val stepDescription = text("stepDescription").nullable()
    val release = text("release").nullable()
    val attempt = integer("attempt")

    /** `operator` or `automatic`. */
    val trigger = text("trigger")
    val byteLength = integer("byteLength")
    val objectKey = text("objectKey")
    val receivedAt = pgTimestamp("receivedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /** When the lifecycle rule removes the bundle; the row is listed until then. */
    val expiresAt = pgTimestamp("expiresAt", 3)

    /** Set on the report that posted this step and release's one Slack message. */
    val slackAnnouncedAt = pgTimestamp("slackAnnouncedAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"SelfHostedInstanceDiagnostic_pkey\"")

    init {
        index("\"SelfHostedInstanceDiagnostic_instanceId_receivedAt_idx\"", false, instanceId, receivedAt)
    }
}
