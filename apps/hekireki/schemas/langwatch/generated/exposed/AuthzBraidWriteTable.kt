package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object AuthzBraidWriteTable : CompositeIdTable("\"AuthzBraidWrite\"") {
    val organizationId = text("organizationId").entityId()
    val commandId = text("commandId").entityId()
    val grantId = text("grantId")
    val readerProjectId = text("readerProjectId")
    val ownerProjectId = text("ownerProjectId")
    val termsKey = text("termsKey")
    val condition = jsonb<String>("condition", { it }, { it })
    val expiresAt = pgTimestamp("expiresAt", 3).nullable()
    val replacesGrantId = text("replacesGrantId").nullable()
    val createdByUserId = text("createdByUserId").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(organizationId, commandId, name = "\"AuthzBraidWrite_pkey\"")

    init {
        index("\"AuthzBraidWrite_organizationId_readerProjectId_createdAt_idx\"", false, organizationId, readerProjectId, createdAt)
    }
}
