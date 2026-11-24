package teable.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class DomainEventOutboxEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DomainEventOutboxEntity>(DomainEventOutboxTable)

    var baseId by DomainEventOutboxTable.baseId
    var tableId by DomainEventOutboxTable.tableId
    var messageName by DomainEventOutboxTable.messageName
    var schemaVersion by DomainEventOutboxTable.schemaVersion
    var aggregateId by DomainEventOutboxTable.aggregateId
    var payload by DomainEventOutboxTable.payload
    var payloadBytes by DomainEventOutboxTable.payloadBytes
    var catalogGeneration by DomainEventOutboxTable.catalogGeneration
    var requiredConsumers by DomainEventOutboxTable.requiredConsumers
    var bindingId by DomainEventOutboxTable.bindingId
    var storageEpoch by DomainEventOutboxTable.storageEpoch
    var unpublished by DomainEventOutboxTable.unpublished
    var settled by DomainEventOutboxTable.settled
    var settledAt by DomainEventOutboxTable.settledAt
    var createdAt by DomainEventOutboxTable.createdAt
}
