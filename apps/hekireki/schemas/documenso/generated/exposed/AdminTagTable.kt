package documenso.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/** @zod.import(["import { ZAdminTagTypeSchema } from '@documenso/lib/types/admin-tags';"]) */
object AdminTagTable : IdTable<String>("\"AdminTag\"") {
    override val id = text("id").entityId()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /** [AdminTagType] @zod.custom.use(ZAdminTagTypeSchema) */
    val type = text("type")
    val description = text("description")
    val userId =
        optReference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AdminTag_userId_fkey\"",
        )
    val organisationId =
        optReference(
            "organisationId",
            OrganisationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AdminTag_organisationId_fkey\"",
        )
    val createdByUserId =
        optReference(
            "createdByUserId",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AdminTag_createdByUserId_fkey\"",
        )
    val adminRuleId =
        optReference(
            "adminRuleId",
            AdminRuleTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AdminTag_adminRuleId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"AdminTag_pkey\"")

    init {
        index("\"AdminTag_userId_idx\"", false, userId)
        index("\"AdminTag_organisationId_idx\"", false, organisationId)
    }
}
