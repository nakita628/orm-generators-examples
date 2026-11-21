package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One row per erased account, holding only its creation time: no id, name or address. It keeps
 * the install's first-created account knowable after an erasure deletes that row.
 * Spec: modules/user/specs/user.feature
 */
class UserErasureEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, UserErasureEntity>(UserErasureTable)

    var accountCreatedAt by UserErasureTable.accountCreatedAt
    var erasedAt by UserErasureTable.erasedAt
}
