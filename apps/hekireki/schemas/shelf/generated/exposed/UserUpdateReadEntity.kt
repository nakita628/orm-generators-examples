package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class UserUpdateReadEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, UserUpdateReadEntity>(UserUpdateReadTable)

    var user by UserEntity referencedOn UserUpdateReadTable.userId
    var userId by UserUpdateReadTable.userId
    var update by UpdateEntity referencedOn UserUpdateReadTable.updateId
    var updateId by UserUpdateReadTable.updateId
    var readAt by UserUpdateReadTable.readAt
}
