package shelf.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.Table

object CategoryToCustomFieldTable : Table("\"_CategoryToCustomField\"") {
    val a =
        reference(
            "\"A\"",
            CategoryTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_CategoryToCustomField_A_fkey\"",
        )
    val b =
        reference(
            "\"B\"",
            CustomFieldTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_CategoryToCustomField_B_fkey\"",
        )
    override val primaryKey = PrimaryKey(a, b, name = "\"_CategoryToCustomField_AB_pkey\"")

    init {
        index("\"_CategoryToCustomField_B_index\"", false, b)
    }
}
