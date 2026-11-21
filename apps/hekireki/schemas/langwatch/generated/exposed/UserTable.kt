package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object UserTable : IdTable<String>("\"User\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val name = text("name").nullable()
    val email = text("email").nullable().uniqueIndex("\"User_email_key\"")
    val emailVerified = bool("emailVerified").default(false)
    val signupConfirmationPending = bool("signupConfirmationPending").default(false)
    val passkeySignupClaimHash = text("passkeySignupClaimHash").nullable()
    val image = text("image").nullable()
    val pendingSsoSetup = bool("pendingSsoSetup").default(false)
    val userHashKey = text("userHashKey").nullable()

    /**
     * Owned by the two-factor plugin (D06), which is why it is a plain
     * boolean rather than a state: the plugin flips it when a setup is
     * confirmed and when it is turned off, and challenges every sign-in while
     * it is true. The lifecycle - and the history - lives on MfaEnrollment.
     */
    val twoFactorEnabled = bool("twoFactorEnabled").default(false)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val lastLoginAt = pgTimestamp("lastLoginAt", 3).nullable()
    val deactivatedAt = pgTimestamp("deactivatedAt", 3).nullable()
    val lastHomePath = text("lastHomePath").nullable()
    val tracesExplorerTourDismissedAt = pgTimestamp("tracesExplorerTourDismissedAt", 3).nullable()

    /**
     * How Langy should reach this person's code when a task needs a change
     * (ADR-129). "github" means they chose the pull request path and asked to
     * be remembered; NULL means ask. A shared local folder is never stored
     * here, because a folder has to be shared again each time.
     * Spec: specs/langy/langy-code-access.feature
     */
    val langyCodeAccessPreference = text("langyCodeAccessPreference").nullable()

    /**
     * When this person last said "not now" to being offered a passkey
     * (ADR-120). On the account rather than in browser storage for two
     * reasons: a per-browser record forgets on a new device, which is exactly
     * where the offer should be MORE eager, and an interval nobody can enforce
     * is not an interval. NULL means never asked.
     */
    val passkeyNudgeDismissedAt = pgTimestamp("passkeyNudgeDismissedAt", 3).nullable()

    /**
     * The domains this person has said "no thanks" to being offered their
     * colleagues on (D12). Per domain rather than one flag: somebody who
     * changes jobs and verifies a second address should be offered their new
     * employer's organization, not silenced forever by a decision they made
     * about the old one. Empty means never dismissed anything.
     */
    val joinOfferDismissedDomains = array("joinOfferDismissedDomains", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /**
     * This person's answer about browser notifications, one entry per topic:
     * { "langy": "enabled" | "declined" }. A topic that is absent was never
     * answered. The browser's own permission is not stored here.
     * Spec: specs/langy/langy-notifications.feature
     */
    val notificationPreferences = jsonb<String>("notificationPreferences", { it }, { it }).default("{}")

    /**
     * When an operator marked this person's password as no longer trusted. A
     * password sign-in opens no session while it is set; any write of a new
     * password hash clears it. Spec: modules/auth/specs/forced-password-reset.feature
     */
    val passwordResetRequiredAt = pgTimestamp("passwordResetRequiredAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"User_pkey\"")
}
