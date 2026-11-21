package rallly.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class PollInviteEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, PollInviteEntity>(PollInviteTable)

    var pollId by PollInviteTable.pollId
    var email by PollInviteTable.email

    /** Per invitee link. App minted, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet) — no schema default, so forgetting to mint fails loudly. Rotated on reactivation so a revoked (possibly leaked) token never regains access */
    var token by PollInviteTable.token
    var openedAt by PollInviteTable.openedAt
    var remindedAt by PollInviteTable.remindedAt

    /** Revocation is always explicit: set when the host removes a pending invite, never as a side effect of close, schedule or reopen. Deleting a response instead hard deletes the participant, so the SetNull FK clears participantId and reverts the invite to pending with its token intact — the token the deleted response held is freed with it, so the next response through the same link can take it again. Cleared only by re-inviting (reactivation), which also rotates the token */
    var revokedAt by PollInviteTable.revokedAt

    /** Conversion join, merges invite history into the response timeline. Set only when the invitee responds via their token: a typed email is unverified, so a generic link response never claims an invite */
    var participantId by PollInviteTable.participantId
    var createdAt by PollInviteTable.createdAt
    var updatedAt by PollInviteTable.updatedAt
    var poll by PollEntity referencedOn PollInviteTable.pollId
    var participant by ParticipantEntity optionalReferencedOn PollInviteTable.participantId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(PollInviteTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
