package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object ScimTokenTable : IdTable<String>("\"ScimToken\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val tokenHash = text("tokenHash").uniqueIndex("\"ScimToken_tokenHash_key\"")
    val label = text("label")
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ScimToken_organizationId_fkey\"",
        )
    val createdById = text("createdById")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val lastUsedAt = pgTimestamp("lastUsedAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"ScimToken_pkey\"")

    init {
        index("\"ScimToken_organizationId_idx\"", false, organizationId)
    }
}
