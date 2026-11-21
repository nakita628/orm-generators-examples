package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * A durable claim on one live recovery path while activation is between its
 * guard and projected ACTIVE fact. It has no expiry: only projection of the
 * matching activation, or an explicit cancellation, may consume it.
 */
class SsoActivationRecoveryReservationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SsoActivationRecoveryReservationEntity>(SsoActivationRecoveryReservationTable)

    var organizationId by SsoActivationRecoveryReservationTable.organizationId
    var connectionId by SsoActivationRecoveryReservationTable.connectionId
    var createdAt by SsoActivationRecoveryReservationTable.createdAt
}
