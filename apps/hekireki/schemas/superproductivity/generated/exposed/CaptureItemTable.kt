package superproductivity.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * One sealed capture waiting for a device. No relation to its grant: items
 * outlive a grant that is revoked without discarding them.
 */
object CaptureItemTable : IdTable<String>("capture_items") {
    override val id = text("id").entityId()
    val userId =
        reference(
            "user_id",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "capture_items_user_id_fkey",
        )
    val grantId = text("grant_id")
    val ciphertext = text("ciphertext")

    /** Task id the claiming device must use; replaced when another device takes over. */
    val taskId = text("task_id")
    val createdAt = long("created_at")
    val leaseDeviceId = text("lease_device_id").nullable()
    val leaseUntil = long("lease_until").nullable()
    override val primaryKey = PrimaryKey(id, name = "capture_items_pkey")

    init {
        index("capture_items_user_id_created_at_idx", false, userId, createdAt)
        index("capture_items_grant_id_idx", false, grantId)
    }
}
