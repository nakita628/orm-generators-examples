package teable.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

class DomainEventInboxEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<DomainEventInboxEntity>(DomainEventInboxTable)

    val consumerId by DomainEventInboxTable.consumerId
    val eventId by DomainEventInboxTable.eventId
    var createdAt by DomainEventInboxTable.createdAt
}
