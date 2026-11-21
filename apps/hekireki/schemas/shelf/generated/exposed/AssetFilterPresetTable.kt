package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object AssetFilterPresetTable : IdTable<String>("\"AssetFilterPreset\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetFilterPreset_organizationId_fkey\"",
        )
    val ownerId =
        reference(
            "ownerId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetFilterPreset_ownerId_fkey\"",
        )
    val name = text("name")
    val query = text("query")
    val starred = bool("starred").default(false)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AssetFilterPreset_pkey\"")

    init {
        index("asset_filter_presets_owner_lookup_idx", false, organizationId, ownerId)
        uniqueIndex("asset_filter_presets_owner_name_unique", organizationId, ownerId, name)
    }
}
