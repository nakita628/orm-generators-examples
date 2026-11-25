package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** One person's digest for one ISO week, accepted by our own mailer. The id is `userId:isoWeek`. */
class DigestSendEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DigestSendEntity>(DigestSendTable)

    var userId by DigestSendTable.userId
    var organizationId by DigestSendTable.organizationId
    var isoWeek by DigestSendTable.isoWeek
    var templateId by DigestSendTable.templateId
    var sentAt by DigestSendTable.sentAt
    var createdAt by DigestSendTable.createdAt
}
