package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class UserEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, UserEntity>(UserTable)

    var name by UserTable.name
    var email by UserTable.email
    var emailVerified by UserTable.emailVerified
    var signupConfirmationPending by UserTable.signupConfirmationPending
    var passkeySignupClaimHash by UserTable.passkeySignupClaimHash
    var image by UserTable.image
    var pendingSsoSetup by UserTable.pendingSsoSetup
    var userHashKey by UserTable.userHashKey

    /**
     * Owned by the two-factor plugin (D06), which is why it is a plain
     * boolean rather than a state: the plugin flips it when a setup is
     * confirmed and when it is turned off, and challenges every sign-in while
     * it is true. The lifecycle - and the history - lives on MfaEnrollment.
     */
    var twoFactorEnabled by UserTable.twoFactorEnabled
    var createdAt by UserTable.createdAt
    var updatedAt by UserTable.updatedAt
    var lastLoginAt by UserTable.lastLoginAt
    var deactivatedAt by UserTable.deactivatedAt
    var lastHomePath by UserTable.lastHomePath
    var tracesExplorerTourDismissedAt by UserTable.tracesExplorerTourDismissedAt

    /**
     * How Langy should reach this person's code when a task needs a change
     * (ADR-129). "github" means they chose the pull request path and asked to
     * be remembered; NULL means ask. A shared local folder is never stored
     * here, because a folder has to be shared again each time.
     * Spec: specs/langy/langy-code-access.feature
     */
    var langyCodeAccessPreference by UserTable.langyCodeAccessPreference

    /**
     * When this person last said "not now" to being offered a passkey
     * (ADR-120). On the account rather than in browser storage for two
     * reasons: a per-browser record forgets on a new device, which is exactly
     * where the offer should be MORE eager, and an interval nobody can enforce
     * is not an interval. NULL means never asked.
     */
    var passkeyNudgeDismissedAt by UserTable.passkeyNudgeDismissedAt

    /**
     * The domains this person has said "no thanks" to being offered their
     * colleagues on (D12). Per domain rather than one flag: somebody who
     * changes jobs and verifies a second address should be offered their new
     * employer's organization, not silenced forever by a decision they made
     * about the old one. Empty means never dismissed anything.
     */
    var joinOfferDismissedDomains by UserTable.joinOfferDismissedDomains

    /**
     * This person's answer about browser notifications, one entry per topic:
     * { "langy": "enabled" | "declined" }. A topic that is absent was never
     * answered. The browser's own permission is not stored here.
     * Spec: specs/langy/langy-notifications.feature
     */
    var notificationPreferences by UserTable.notificationPreferences

    /**
     * When an operator marked this person's password as no longer trusted. A
     * password sign-in opens no session while it is set; any write of a new
     * password hash clears it. Spec: modules/auth/specs/forced-password-reset.feature
     */
    var passwordResetRequiredAt by UserTable.passwordResetRequiredAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(UserTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
