package formbricks.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

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
object MembershipTable : CompositeIdTable("\"Membership\"") {
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Membership_organizationId_fkey\"",
        )
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Membership_userId_fkey\"",
        )
    val accepted = bool("accepted").default(false)
    val role = pgEnum("role", "\"OrganizationRole\"", OrganizationRole.entries, OrganizationRole::dbName).default(OrganizationRole.MEMBER)
    override val primaryKey = PrimaryKey(userId, organizationId, name = "\"Membership_pkey\"")

    init {
        addIdColumn(userId)
        addIdColumn(organizationId)
        index("\"Membership_organizationId_idx\"", false, organizationId)
    }
}
