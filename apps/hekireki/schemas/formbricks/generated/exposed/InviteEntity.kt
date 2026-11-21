package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Represents pending invitations to join an organization.
 * Used to manage the process of adding new users to an organization.
 * Once accepted, invites are converted into memberships.
 *
 * @property id - Unique identifier for the invite
 * @property email - Email address of the invited user
 * @property name - Optional display name for the invited user
 * @property organization - The organization sending the invite
 * @property creator - The user who created the invite
 * @property acceptor - The user who accepted the invite (if accepted)
 * @property expiresAt - When the invite becomes invalid
 * @property role - Intended role for the invited user
 * @property teamIds - Teams the user will be added to upon acceptance
 */
class InviteEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, InviteEntity>(InviteTable)

    var email by InviteTable.email
    var name by InviteTable.name
    var organization by OrganizationEntity referencedOn InviteTable.organizationId
    var organizationId by InviteTable.organizationId
    var creator by UserEntity referencedOn InviteTable.creatorId
    var creatorId by InviteTable.creatorId
    var acceptor by UserEntity optionalReferencedOn InviteTable.acceptorId
    var acceptorId by InviteTable.acceptorId
    var createdAt by InviteTable.createdAt
    var expiresAt by InviteTable.expiresAt
    var role by InviteTable.role
    var teamIds by InviteTable.teamIds
}
