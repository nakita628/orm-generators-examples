package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One usage report, exactly as it arrived. Kept so a number on an install's
 * row can be traced back to the report that set it, and so a question asked
 * later can be answered over the history rather than over the latest snapshot.
 */
object SelfHostedInstanceReportTable : IdTable<String>("\"SelfHostedInstanceReport\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()

    /**
     * The `SelfHostedInstance.instanceId` this report came from. A plain column
     * ref for the same reason the instance row's refs are.
     */
    val instanceId = text("instanceId")
    val receivedAt = pgTimestamp("receivedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val version = text("version").nullable()
    val reportSchemaVersion = integer("reportSchemaVersion").nullable()

    /** Fields this release had no name for. Counted, never stored by name. */
    val unknownFields = integer("unknownFields").default(0)

    /** The validated report, with every unknown field already dropped. */
    val payload = jsonb<String>("payload", { it }, { it })
    override val primaryKey = PrimaryKey(id, name = "\"SelfHostedInstanceReport_pkey\"")

    init {
        index("\"SelfHostedInstanceReport_instanceId_receivedAt_idx\"", false, instanceId, receivedAt)
        index("\"SelfHostedInstanceReport_receivedAt_idx\"", false, receivedAt)
    }
}
