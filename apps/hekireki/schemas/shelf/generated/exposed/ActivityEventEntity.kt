package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class ActivityEventEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ActivityEventEntity>(ActivityEventTable)

    var organizationId by ActivityEventTable.organizationId
    var occurredAt by ActivityEventTable.occurredAt
    var actorUserId by ActivityEventTable.actorUserId
    var actorSnapshot by ActivityEventTable.actorSnapshot
    var action by ActivityEventTable.action
    var entityType by ActivityEventTable.entityType
    var entityId by ActivityEventTable.entityId
    var assetId by ActivityEventTable.assetId
    var bookingId by ActivityEventTable.bookingId
    var auditSessionId by ActivityEventTable.auditSessionId
    var auditAssetId by ActivityEventTable.auditAssetId
    var kitId by ActivityEventTable.kitId
    var locationId by ActivityEventTable.locationId
    var teamMemberId by ActivityEventTable.teamMemberId
    var targetUserId by ActivityEventTable.targetUserId
    var field by ActivityEventTable.field
    var fromValue by ActivityEventTable.fromValue
    var toValue by ActivityEventTable.toValue
    var meta by ActivityEventTable.meta
    var organization by OrganizationEntity referencedOn ActivityEventTable.organizationId
}
