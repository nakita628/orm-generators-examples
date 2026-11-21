package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Better Auth (ENG-1054) two-factor secrets. Supersedes the legacy User.twoFactorSecret /
 * User.backupCodes columns (kept until cutover). `secret` and `backupCodes` are encrypted
 * at rest by Better Auth using BETTER_AUTH_SECRET.
 */
class TwoFactorEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TwoFactorEntity>(TwoFactorTable)

    var secret by TwoFactorTable.secret
    var backupCodes by TwoFactorTable.backupCodes
    var userId by TwoFactorTable.userId
    var user by UserEntity referencedOn TwoFactorTable.userId
    var verified by TwoFactorTable.verified

    /** Better Auth 1.6.23+ two-factor brute-force lockout: consecutive failed verification attempts. */
    var failedVerificationCount by TwoFactorTable.failedVerificationCount

    /** Better Auth 1.6.23+ two-factor brute-force lockout: set once failedVerificationCount trips the limit. */
    var lockedUntil by TwoFactorTable.lockedUntil
}
