package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object AuditSessionTable : IdTable<String>("\"AuditSession\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name")
    val description = text("description").nullable()
    val targetId = text("targetId").nullable()
    val status = pgEnum("status", "\"AuditStatus\"", AuditStatus.entries, AuditStatus::dbName).default(AuditStatus.PENDING)
    val scopeMeta = jsonb<String>("scopeMeta", { it }, { it }).nullable()
    val expectedAssetCount = integer("expectedAssetCount").default(0)
    val foundAssetCount = integer("foundAssetCount").default(0)
    val missingAssetCount = integer("missingAssetCount").default(0)
    val unexpectedAssetCount = integer("unexpectedAssetCount").default(0)
    val startedAt = pgTimestamp("startedAt", 3).nullable()
    val dueDate = pgTimestamp("dueDate", 3).nullable()
    val completedAt = pgTimestamp("completedAt", 3).nullable()
    val cancelledAt = pgTimestamp("cancelledAt", 3).nullable()
    val activeSchedulerReference = text("activeSchedulerReference").nullable()
    val createdById =
        reference(
            "createdById",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditSession_createdById_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditSession_organizationId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AuditSession_pkey\"")

    init {
        index("\"AuditSession_organizationId_status_idx\"", false, organizationId, status)
        index("\"AuditSession_createdById_idx\"", false, createdById)
        index("\"AuditSession_status_createdAt_idx\"", false, status, createdAt)
    }
}
