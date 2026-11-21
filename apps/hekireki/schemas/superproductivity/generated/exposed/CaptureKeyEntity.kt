package superproductivity.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Public half of the user's capture key. The row outlives switching the
 * feature off: `disabledAt` stops a stale device re-registering that key.
 */
class CaptureKeyEntity(
    id: EntityID<Int>,
) : Entity<Int>(id) {
    companion object : EntityClass<Int, CaptureKeyEntity>(CaptureKeyTable)

    /** Base64 SPKI DER of an ECDH P-256 public key. */
    var publicKey by CaptureKeyTable.publicKey
    var fingerprint by CaptureKeyTable.fingerprint
    var createdAt by CaptureKeyTable.createdAt
    var disabledAt by CaptureKeyTable.disabledAt
    val user by UserEntity referencedOn CaptureKeyTable.id
}
