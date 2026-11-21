package superproductivity.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One sealed capture waiting for a device. No relation to its grant: items
 * outlive a grant that is revoked without discarding them.
 */
class CaptureItemEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, CaptureItemEntity>(CaptureItemTable)

    var userId by CaptureItemTable.userId
    var grantId by CaptureItemTable.grantId
    var ciphertext by CaptureItemTable.ciphertext

    /** Task id the claiming device must use; replaced when another device takes over. */
    var taskId by CaptureItemTable.taskId
    var createdAt by CaptureItemTable.createdAt
    var leaseDeviceId by CaptureItemTable.leaseDeviceId
    var leaseUntil by CaptureItemTable.leaseUntil
    var user by UserEntity referencedOn CaptureItemTable.userId
}
