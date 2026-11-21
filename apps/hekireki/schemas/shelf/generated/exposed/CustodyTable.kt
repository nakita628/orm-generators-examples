package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object CustodyTable : IdTable<String>("\"Custody\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val teamMemberId =
        reference(
            "teamMemberId",
            TeamMemberTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Custody_teamMemberId_fkey\"",
        )
    val assetId =
        reference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Custody_assetId_fkey\"",
        )
    val kitCustodyId =
        optReference(
            "kitCustodyId",
            KitCustodyTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Custody_kitCustodyId_fkey\"",
        )
    val quantity = integer("quantity").default(1)
    val locationId =
        optReference(
            "locationId",
            LocationTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Custody_locationId_fkey\"",
        )
    val sourceUnknown = bool("sourceUnknown").default(false)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Custody_pkey\"")

    init {
        index("\"Custody_assetId_teamMemberId_idx\"", false, assetId, teamMemberId)
        index("\"Custody_teamMemberId_idx\"", false, teamMemberId)
        index("\"Custody_kitCustodyId_idx\"", false, kitCustodyId)
        index("\"Custody_locationId_idx\"", false, locationId)
    }
}
