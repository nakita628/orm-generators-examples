package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * Nurturing's own copy of each organization's members, from the owners' member facts (§9):
 * organization's signed_up and invite_accepted and auth's sso_auto_added join, member_removed
 * and member_disabled leave, member_enabled returns. Each column keeps its latest time.
 */
object NurturingOrganizationMemberTable : CompositeIdTable("\"NurturingOrganizationMember\"") {
    val organizationId = text("organizationId").entityId()
    val userId = text("userId").entityId()
    val joinedAt = pgTimestamp("joinedAt", 3).nullable()
    val leftAt = pgTimestamp("leftAt", 3).nullable()
    val returnedAt = pgTimestamp("returnedAt", 3).nullable()
    override val primaryKey = PrimaryKey(organizationId, userId, name = "\"NurturingOrganizationMember_pkey\"")

    init {
        index("\"NurturingOrganizationMember_userId_idx\"", false, userId)
    }
}
