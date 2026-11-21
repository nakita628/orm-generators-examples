package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A custom field value an AssetModel gives the assets that join it.
 *
 * `value` has the exact shape of `AssetCustomFieldValue.value`, so filling an
 * asset is a plain copy. Rows cascade away with their model or field.
 */
object AssetModelCustomFieldDefaultTable : IdTable<String>("\"AssetModelCustomFieldDefault\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val value = jsonb<String>("value", { it }, { it })
    val assetModelId =
        reference(
            "assetModelId",
            AssetModelTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetModelCustomFieldDefault_assetModelId_fkey\"",
        )
    val customFieldId = text("customFieldId")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AssetModelCustomFieldDefault_pkey\"")

    init {
        index("\"AssetModelCustomFieldDefault_customFieldId_idx\"", false, customFieldId)
        uniqueIndex("\"AssetModelCustomFieldDefault_assetModelId_customFieldId_key\"", assetModelId, customFieldId)
    }
}
