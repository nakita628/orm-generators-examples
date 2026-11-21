package inboxzero.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One condition row in the split builder. A split may hold several of the same
 * kind (two senders, two labels), so kind is deliberately not unique per split.
 */
class MailSplitFilterEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, MailSplitFilterEntity>(MailSplitFilterTable)

    var kind by MailSplitFilterTable.kind
    var value by MailSplitFilterTable.value
    var order by MailSplitFilterTable.order
    var mailSplitId by MailSplitFilterTable.mailSplitId
    var mailSplit by MailSplitEntity referencedOn MailSplitFilterTable.mailSplitId
}
