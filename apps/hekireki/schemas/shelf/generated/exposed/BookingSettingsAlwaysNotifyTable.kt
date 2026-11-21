package shelf.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.Table

object BookingSettingsAlwaysNotifyTable : Table("\"_BookingSettingsAlwaysNotify\"") {
    val a =
        reference(
            "\"A\"",
            BookingSettingsTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_BookingSettingsAlwaysNotify_A_fkey\"",
        )
    val b =
        reference(
            "\"B\"",
            TeamMemberTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_BookingSettingsAlwaysNotify_B_fkey\"",
        )
    override val primaryKey = PrimaryKey(a, b, name = "\"_BookingSettingsAlwaysNotify_AB_pkey\"")

    init {
        index("\"_BookingSettingsAlwaysNotify_B_index\"", false, b)
    }
}
