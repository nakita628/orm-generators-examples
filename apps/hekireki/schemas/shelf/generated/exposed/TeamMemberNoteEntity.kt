package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class TeamMemberNoteEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TeamMemberNoteEntity>(TeamMemberNoteTable)

    var content by TeamMemberNoteTable.content
    var type by TeamMemberNoteTable.type
    var createdAt by TeamMemberNoteTable.createdAt
    var updatedAt by TeamMemberNoteTable.updatedAt
    var user by UserEntity optionalReferencedOn TeamMemberNoteTable.userId
    var userId by TeamMemberNoteTable.userId
    var teamMember by TeamMemberEntity referencedOn TeamMemberNoteTable.teamMemberId
    var teamMemberId by TeamMemberNoteTable.teamMemberId
    var organization by OrganizationEntity referencedOn TeamMemberNoteTable.organizationId
    var organizationId by TeamMemberNoteTable.organizationId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(TeamMemberNoteTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
