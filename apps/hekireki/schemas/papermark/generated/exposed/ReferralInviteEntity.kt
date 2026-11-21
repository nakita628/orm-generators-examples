package papermark.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class ReferralInviteEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ReferralInviteEntity>(ReferralInviteTable)

    var userId by ReferralInviteTable.userId
    var user by UserEntity referencedOn ReferralInviteTable.userId
    var teamId by ReferralInviteTable.teamId
    var team by TeamEntity optionalReferencedOn ReferralInviteTable.teamId

    /**
     * *
     * * Whose cap and discount this invite counts against — the fund's team, or the
     * * sender alone. One non-null key rather than a pair of nullable ones: on a
     * * nullable teamId every personal invite is distinct to Postgres, so the two
     * * keys guarded different scopes and a member could not invite a company both
     * * for their fund and for themselves.
     */
    var scopeId by ReferralInviteTable.scopeId
    var email by ReferralInviteTable.email
    var referralLink by ReferralInviteTable.referralLink
    var sentAt by ReferralInviteTable.sentAt
}
