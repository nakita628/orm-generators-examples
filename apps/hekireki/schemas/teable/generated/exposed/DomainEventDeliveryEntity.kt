package teable.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class DomainEventDeliveryEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DomainEventDeliveryEntity>(DomainEventDeliveryTable)

    var eventId by DomainEventDeliveryTable.eventId
    var consumerId by DomainEventDeliveryTable.consumerId
    var status by DomainEventDeliveryTable.status
    var attempts by DomainEventDeliveryTable.attempts
    var maxAttempts by DomainEventDeliveryTable.maxAttempts
    var leaseToken by DomainEventDeliveryTable.leaseToken
    var leaseExpiresAt by DomainEventDeliveryTable.leaseExpiresAt
    var nextAttemptAt by DomainEventDeliveryTable.nextAttemptAt
    var lastError by DomainEventDeliveryTable.lastError
    var createdAt by DomainEventDeliveryTable.createdAt
}
