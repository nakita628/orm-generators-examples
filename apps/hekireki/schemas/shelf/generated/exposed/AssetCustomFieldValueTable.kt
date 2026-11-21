package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object AssetCustomFieldValueTable : IdTable<String>("\"AssetCustomFieldValue\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val value = jsonb<String>("value", { it }, { it })
    val assetId =
        reference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetCustomFieldValue_assetId_fkey\"",
        )
    val customFieldId =
        reference(
            "customFieldId",
            CustomFieldTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetCustomFieldValue_customFieldId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AssetCustomFieldValue_pkey\"")

    init {
        index("\"AssetCustomFieldValue_lookup_idx\"", false, assetId, customFieldId)
        index("\"AssetCustomFieldValue_customFieldId_idx\"", false, customFieldId)
    }
}
