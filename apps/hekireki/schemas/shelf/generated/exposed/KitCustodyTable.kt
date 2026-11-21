package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object KitCustodyTable : IdTable<String>("\"KitCustody\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val custodianId =
        reference(
            "custodianId",
            TeamMemberTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"KitCustody_custodianId_fkey\"",
        )
    val kitId =
        reference(
            "kitId",
            KitTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"KitCustody_kitId_fkey\"",
        ).uniqueIndex("\"KitCustody_kitId_key\"")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"KitCustody_pkey\"")

    init {
        index("\"KitCustody_custodianId_idx\"", false, custodianId)
    }
}
