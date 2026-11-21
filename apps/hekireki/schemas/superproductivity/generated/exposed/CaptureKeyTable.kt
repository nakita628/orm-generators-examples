package superproductivity.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Public half of the user's capture key. The row outlives switching the
 * feature off: `disabledAt` stops a stale device re-registering that key.
 */
object CaptureKeyTable : IdTable<Int>("capture_keys") {
    override val id =
        integer("user_id")
            .references(
                UserTable.id,
                onDelete = ReferenceOption.CASCADE,
                onUpdate = ReferenceOption.CASCADE,
                fkName = "capture_keys_user_id_fkey",
            ).entityId()

    /** Base64 SPKI DER of an ECDH P-256 public key. */
    val publicKey = text("public_key")
    val fingerprint = text("fingerprint")
    val createdAt = long("created_at")
    val disabledAt = long("disabled_at").nullable()
    override val primaryKey = PrimaryKey(id, name = "capture_keys_pkey")
}
