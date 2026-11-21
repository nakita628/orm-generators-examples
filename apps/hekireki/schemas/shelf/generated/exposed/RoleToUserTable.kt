package shelf.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.Table

object RoleToUserTable : Table("\"_RoleToUser\"") {
    val a =
        reference(
            "\"A\"",
            RoleTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_RoleToUser_A_fkey\"",
        )
    val b =
        reference(
            "\"B\"",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_RoleToUser_B_fkey\"",
        )
    override val primaryKey = PrimaryKey(a, b, name = "\"_RoleToUser_AB_pkey\"")

    init {
        index("\"_RoleToUser_B_index\"", false, b)
    }
}
