package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object GatewayA2ATargetTable : IdTable<String>("\"GatewayA2ATarget\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")
    val projectId = text("projectId").nullable()
    val name = text("name")
    val agentName = text("agentName")
    val cardUrl = text("cardUrl")
    val upstreamUrl = text("upstreamUrl")
    val encryptedUpstreamHeaders = text("encryptedUpstreamHeaders").nullable()
    val restrictToAllowList = bool("restrictToAllowList").default(false)
    val allowedVirtualKeyIds = array("allowedVirtualKeyIds", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }
    val blocked = bool("blocked").default(false)
    val blockReason = text("blockReason").nullable()
    val createdById = text("createdById")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val archivedAt = pgTimestamp("archivedAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"GatewayA2ATarget_pkey\"")

    init {
        index("\"GatewayA2ATarget_organizationId_projectId_idx\"", false, organizationId, projectId)
    }
}
