package superproductivity.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

class SyncDeviceEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<SyncDeviceEntity>(SyncDeviceTable)

    val clientId by SyncDeviceTable.clientId
    val userId by SyncDeviceTable.userId
    var deviceName by SyncDeviceTable.deviceName
    var userAgent by SyncDeviceTable.userAgent

    /**
     * Bare semver the client last reported on download (`18.22.0`), NULL until
     * it does. Read by the checkpoint gate (#9962); never shown to users.
     */
    var appVersion by SyncDeviceTable.appVersion
    var lastSeenAt by SyncDeviceTable.lastSeenAt

    /** Reserved: written once at device registration, never read or updated. */
    var lastAckedSeq by SyncDeviceTable.lastAckedSeq
    var createdAt by SyncDeviceTable.createdAt
    val user by UserEntity referencedOn SyncDeviceTable.userId
}
