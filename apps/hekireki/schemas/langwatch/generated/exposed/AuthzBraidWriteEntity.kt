package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

class AuthzBraidWriteEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<AuthzBraidWriteEntity>(AuthzBraidWriteTable)

    val organizationId by AuthzBraidWriteTable.organizationId
    val commandId by AuthzBraidWriteTable.commandId
    var grantId by AuthzBraidWriteTable.grantId
    var readerProjectId by AuthzBraidWriteTable.readerProjectId
    var ownerProjectId by AuthzBraidWriteTable.ownerProjectId
    var termsKey by AuthzBraidWriteTable.termsKey
    var condition by AuthzBraidWriteTable.condition
    var expiresAt by AuthzBraidWriteTable.expiresAt
    var replacesGrantId by AuthzBraidWriteTable.replacesGrantId
    var createdByUserId by AuthzBraidWriteTable.createdByUserId
    var createdAt by AuthzBraidWriteTable.createdAt
}
