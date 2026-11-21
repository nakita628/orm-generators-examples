package shelf.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.Table

object AssetReminderToTeamMemberTable : Table("\"_AssetReminderToTeamMember\"") {
    val a =
        reference(
            "\"A\"",
            AssetReminderTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_AssetReminderToTeamMember_A_fkey\"",
        )
    val b =
        reference(
            "\"B\"",
            TeamMemberTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_AssetReminderToTeamMember_B_fkey\"",
        )
    override val primaryKey = PrimaryKey(a, b, name = "\"_AssetReminderToTeamMember_AB_pkey\"")

    init {
        index("\"_AssetReminderToTeamMember_B_index\"", false, b)
    }
}
