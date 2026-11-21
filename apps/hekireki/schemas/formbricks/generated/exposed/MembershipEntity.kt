package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * Links users to organizations with specific roles.
 * Manages organization membership and permissions.
 * Core model for managing user access within organizations.
 *
 * @property organization - The organization the user belongs to
 * @property user - The member user
 * @property accepted - Whether the user has accepted the membership
 * @property role - User's role within the organization (owner, manager, member, billing)
 */
class MembershipEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<MembershipEntity>(MembershipTable)

    val organization by OrganizationEntity referencedOn MembershipTable.organizationId
    val organizationId by MembershipTable.organizationId
    val user by UserEntity referencedOn MembershipTable.userId
    val userId by MembershipTable.userId
    var accepted by MembershipTable.accepted
    var role by MembershipTable.role
}
