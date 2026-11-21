package shelf.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.Table

object BookingToTagTable : Table("\"_BookingToTag\"") {
    val a =
        reference(
            "\"A\"",
            BookingTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_BookingToTag_A_fkey\"",
        )
    val b =
        reference(
            "\"B\"",
            TagTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_BookingToTag_B_fkey\"",
        )
    override val primaryKey = PrimaryKey(a, b, name = "\"_BookingToTag_AB_pkey\"")

    init {
        index("\"_BookingToTag_B_index\"", false, b)
    }
}
