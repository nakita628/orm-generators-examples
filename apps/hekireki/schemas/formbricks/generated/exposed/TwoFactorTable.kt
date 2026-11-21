package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Better Auth (ENG-1054) two-factor secrets. Supersedes the legacy User.twoFactorSecret /
 * User.backupCodes columns (kept until cutover). `secret` and `backupCodes` are encrypted
 * at rest by Better Auth using BETTER_AUTH_SECRET.
 */
object TwoFactorTable : IdTable<String>("\"TwoFactor\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val secret = text("secret")
    val backupCodes = text("backupCodes")
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TwoFactor_userId_fkey\"",
        ).uniqueIndex("\"TwoFactor_userId_key\"")
    val verified = bool("verified").default(true)

    /** Better Auth 1.6.23+ two-factor brute-force lockout: consecutive failed verification attempts. */
    val failedVerificationCount = integer("failedVerificationCount").default(0)

    /** Better Auth 1.6.23+ two-factor brute-force lockout: set once failedVerificationCount trips the limit. */
    val lockedUntil = pgTimestamp("lockedUntil", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"TwoFactor_pkey\"")
}
