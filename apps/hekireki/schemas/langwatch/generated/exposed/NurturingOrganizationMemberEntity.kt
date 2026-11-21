package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * Nurturing's own copy of each organization's members, from the owners' member facts (§9):
 * organization's signed_up and invite_accepted and auth's sso_auto_added join, member_removed
 * and member_disabled leave, member_enabled returns. Each column keeps its latest time.
 */
class NurturingOrganizationMemberEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<NurturingOrganizationMemberEntity>(NurturingOrganizationMemberTable)

    val organizationId by NurturingOrganizationMemberTable.organizationId
    val userId by NurturingOrganizationMemberTable.userId
    var joinedAt by NurturingOrganizationMemberTable.joinedAt
    var leftAt by NurturingOrganizationMemberTable.leftAt
    var returnedAt by NurturingOrganizationMemberTable.returnedAt
}
