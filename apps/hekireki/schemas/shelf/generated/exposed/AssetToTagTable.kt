package shelf.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.Table

object AssetToTagTable : Table("\"_AssetToTag\"") {
    val a =
        reference(
            "\"A\"",
            AssetTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_AssetToTag_A_fkey\"",
        )
    val b =
        reference(
            "\"B\"",
            TagTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_AssetToTag_B_fkey\"",
        )
    override val primaryKey = PrimaryKey(a, b, name = "\"_AssetToTag_AB_pkey\"")

    init {
        index("\"_AssetToTag_B_index\"", false, b)
    }
}
