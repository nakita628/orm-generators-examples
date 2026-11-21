package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class InviteEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, InviteEntity>(InviteTable)

    var inviter by UserEntity referencedOn InviteTable.inviterId
    var inviterId by InviteTable.inviterId
    var organization by OrganizationEntity referencedOn InviteTable.organizationId
    var organizationId by InviteTable.organizationId
    var inviteeUser by UserEntity optionalReferencedOn InviteTable.inviteeUserId
    var inviteeUserId by InviteTable.inviteeUserId
    var inviteeTeamMember by TeamMemberEntity referencedOn InviteTable.teamMemberId
    var teamMemberId by InviteTable.teamMemberId
    var inviteeEmail by InviteTable.inviteeEmail
    var status by InviteTable.status
    var inviteCode by InviteTable.inviteCode
    var roles by InviteTable.roles
    var inviteMessage by InviteTable.inviteMessage
    var expiresAt by InviteTable.expiresAt
    var createdAt by InviteTable.createdAt
    var updatedAt by InviteTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(InviteTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
