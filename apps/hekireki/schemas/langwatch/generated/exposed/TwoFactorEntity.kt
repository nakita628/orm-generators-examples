package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The two-factor plugin's own table (D06). Protocol state, and row-truth for
 * good (ADR-101 R12): the shared secret and the backup codes live here and
 * in no event, ever. Nothing in this repo reads or writes these columns -
 * the plugin owns them end to end, including the lockout counters, which is
 * why `failedVerificationCount` and `lockedUntil` are here rather than
 * rebuilt on the projection.
 *
 * `backupCodes` holds an ENCRYPTED blob, not hashes. A backup code has to be
 * compared against what somebody types, and the plugin decrypts and compares;
 * hashing them would leave it unable to verify its own codes.
 */
class TwoFactorEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TwoFactorEntity>(TwoFactorTable)

    var userId by TwoFactorTable.userId
    var secret by TwoFactorTable.secret
    var backupCodes by TwoFactorTable.backupCodes
    var verified by TwoFactorTable.verified
    var failedVerificationCount by TwoFactorTable.failedVerificationCount
    var lockedUntil by TwoFactorTable.lockedUntil
    var createdAt by TwoFactorTable.createdAt
    var updatedAt by TwoFactorTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(TwoFactorTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
