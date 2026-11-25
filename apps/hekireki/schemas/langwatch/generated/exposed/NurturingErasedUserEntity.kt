package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Each account nurturing erased, as a SHA-256 of its id and never the id: a gateway key fact
 * naming it, however late it lands, stores no creator or holder. Platform-wide, like the erasure.
 */
class NurturingErasedUserEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, NurturingErasedUserEntity>(NurturingErasedUserTable)

    var erasedAt by NurturingErasedUserTable.erasedAt
}
