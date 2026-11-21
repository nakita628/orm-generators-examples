package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * A durable claim on one live recovery path while activation is between its
 * guard and projected ACTIVE fact. It has no expiry: only projection of the
 * matching activation, or an explicit cancellation, may consume it.
 */
object SsoActivationRecoveryReservationTable : IdTable<String>("\"SsoActivationRecoveryReservation\"") {
    override val id = text("commandId").entityId()
    val organizationId = text("organizationId")
    val connectionId = text("connectionId")
    val createdAt = pgTimestamp("createdAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"SsoActivationRecoveryReservation_pkey\"")

    init {
        index("\"SsoActivationRecoveryReservation_organizationId_idx\"", false, organizationId)
        uniqueIndex("SsoActivationRecoveryReservation_organizationId_connectionI_key", organizationId, connectionId)
    }
}
