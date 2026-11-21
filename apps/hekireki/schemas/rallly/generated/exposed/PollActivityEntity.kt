package rallly.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class PollActivityEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, PollActivityEntity>(PollActivityTable)

    var pollId by PollActivityTable.pollId

    /** Validated by a zod union in the feature, not a DB enum: the activity vocabulary grows without migrations */
    var type by PollActivityTable.type

    /** Subject ids are soft references (no FKs): activity must outlive its subjects — including account deletion — so refs persist as identity and the payload carries the display snapshot; whether a ref still resolves is the renderer's question */
    var userId by PollActivityTable.userId
    var participantId by PollActivityTable.participantId
    var inviteId by PollActivityTable.inviteId
    var optionId by PollActivityTable.optionId
    var payload by PollActivityTable.payload
    var createdAt by PollActivityTable.createdAt
    var poll by PollEntity referencedOn PollActivityTable.pollId
}
