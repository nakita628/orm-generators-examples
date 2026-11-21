package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class RoleChangeLogEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, RoleChangeLogEntity>(RoleChangeLogTable)

    var previousRole by RoleChangeLogTable.previousRole
    var newRole by RoleChangeLogTable.newRole

    /** How the change was made. SSO rows name the member themself as `changedBy`. */
    var source by RoleChangeLogTable.source1
    var createdAt by RoleChangeLogTable.createdAt
    var user by UserEntity referencedOn RoleChangeLogTable.userId
    var userId by RoleChangeLogTable.userId
    var changedBy by UserEntity referencedOn RoleChangeLogTable.changedById
    var changedById by RoleChangeLogTable.changedById
    var organization by OrganizationEntity referencedOn RoleChangeLogTable.organizationId
    var organizationId by RoleChangeLogTable.organizationId
}
