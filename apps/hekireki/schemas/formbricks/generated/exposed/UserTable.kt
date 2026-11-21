package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents a user in the Formbricks system.
 * Central model for user authentication and profile management.
 *
 * @property id - Unique identifier for the user
 * @property name - Display name of the user
 * @property email - User's email address
 * @property twoFactorEnabled - Whether 2FA is active
 * @property memberships - Organizations the user belongs to
 * @property notificationSettings - User's notification preferences
 */
object UserTable : IdTable<String>("\"User\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val name = text("name")
    val email = text("email").uniqueIndex("\"User_email_key\"")
    val emailVerified = bool("email_verified").default(false)
    val twoFactorSecret = text("twoFactorSecret").nullable()
    val twoFactorEnabled = bool("twoFactorEnabled").default(false)
    val backupCodes = text("backupCodes").nullable()
    val password = text("password").nullable()
    val identityProvider =
        pgEnum("identityProvider", "\"IdentityProvider\"", IdentityProvider.entries, IdentityProvider::dbName).default(
            IdentityProvider.EMAIL,
        )
    val identityProviderAccountId = text("identityProviderAccountId").nullable()
    val groupId = text("groupId").nullable()

    /** [UserNotificationSettings] */
    val notificationSettings = jsonb<String>("notificationSettings", { it }, { it }).default("{}")

    /** [Locale] */
    val locale = text("locale").default("en-US")
    val lastLoginAt = pgTimestamp("lastLoginAt", 3).nullable()
    val isActive = bool("isActive").default(true)

    /**
     * When an admin last reactivated this user. Restarts the members retention clock (ENG-3713)
     * without touching `lastLoginAt`, which the member list, usage counts and APIs report as a real
     * sign-in.
     */
    val reactivatedAt = pgTimestamp("reactivatedAt", 3).nullable()

    /**
     * When the user last completed a sign-in (`lastActivePlugin`, after any second factor) or renewed a
     * session (Better Auth `databaseHooks.session.update`), on the database clock and never moved
     * backwards. The members retention clock reads it: sessions are
     * deleted on sign-out and expiry, so their own timestamps can't serve as a record of activity.
     */
    val lastActiveAt = pgTimestamp("lastActiveAt", 3).nullable()
    val isBootstrapAdmin =
        pgEnum("isBootstrapAdmin", "\"BootstrapAdminMarker\"", BootstrapAdminMarker.entries, BootstrapAdminMarker::dbName)
            .nullable()
            .uniqueIndex("\"User_isBootstrapAdmin_key\"")
    override val primaryKey = PrimaryKey(id, name = "\"User_pkey\"")
}
