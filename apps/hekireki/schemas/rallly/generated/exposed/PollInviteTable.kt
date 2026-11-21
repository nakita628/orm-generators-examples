package rallly.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object PollInviteTable : IdTable<String>("poll_invites") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val pollId =
        reference(
            "poll_id",
            PollTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "poll_invites_poll_id_fkey",
        )
    val email = pgCitext("email")

    /** Per invitee link. App minted, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet) — no schema default, so forgetting to mint fails loudly. Rotated on reactivation so a revoked (possibly leaked) token never regains access */
    val token = text("token").uniqueIndex("poll_invites_token_key")
    val openedAt = pgTimestamp("opened_at", 3).nullable()
    val remindedAt = pgTimestamp("reminded_at", 3).nullable()

    /** Revocation is always explicit: set when the host removes a pending invite, never as a side effect of close, schedule or reopen. Deleting a response instead hard deletes the participant, so the SetNull FK clears participantId and reverts the invite to pending with its token intact — the token the deleted response held is freed with it, so the next response through the same link can take it again. Cleared only by re-inviting (reactivation), which also rotates the token */
    val revokedAt = pgTimestamp("revoked_at", 3).nullable()

    /** Conversion join, merges invite history into the response timeline. Set only when the invitee responds via their token: a typed email is unverified, so a generic link response never claims an invite */
    val participantId =
        optReference(
            "participant_id",
            ParticipantTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "poll_invites_participant_id_fkey",
        ).uniqueIndex("poll_invites_participant_id_key")
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "poll_invites_pkey")

    init {
        uniqueIndex("poll_invites_poll_id_email_key", pollId, email)
    }
}
