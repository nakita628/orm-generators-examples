package shelf.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.Table

object BookingNotificationRecipientsTable : Table("\"_BookingNotificationRecipients\"") {
    val a =
        reference(
            "\"A\"",
            BookingTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_BookingNotificationRecipients_A_fkey\"",
        )
    val b =
        reference(
            "\"B\"",
            TeamMemberTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"_BookingNotificationRecipients_B_fkey\"",
        )
    override val primaryKey = PrimaryKey(a, b, name = "\"_BookingNotificationRecipients_AB_pkey\"")

    init {
        index("\"_BookingNotificationRecipients_B_index\"", false, b)
    }
}
