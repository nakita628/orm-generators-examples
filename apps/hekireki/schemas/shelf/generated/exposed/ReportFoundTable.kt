package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object ReportFoundTable : IdTable<String>("\"ReportFound\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val email = text("email")
    val content = text("content")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val assetId =
        optReference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ReportFound_assetId_fkey\"",
        )
    val kitId =
        optReference(
            "kitId",
            KitTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ReportFound_kitId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"ReportFound_pkey\"")

    init {
        index("\"ReportFound_assetId_idx\"", false, assetId)
        index("\"ReportFound_kitId_idx\"", false, kitId)
    }
}
