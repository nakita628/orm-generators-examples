package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Consecutive failed sign-ins, and the lock they earn (GAC-09).
 * Spec: specs/identity/org-account-lockout.feature.
 *
 * KEYED ON THE ADDRESS THAT WAS TYPED, NOT ON A USER, and that is the whole
 * design. We already answer "no such account" and "wrong password" with one
 * refusal so that sign-in cannot be used to discover who has an account. A
 * lock that existed only for real accounts would put that back: five wrong
 * guesses saying "locked" for one address and "incorrect" for another tells
 * an attacker which addresses are worth attacking. So an address with
 * nothing behind it gets a row, a count and a lock exactly like one that
 * resolves, and both answer identically.
 *
 * `identifierHash` is a KEYED hash of the normalised address, never the
 * address. The key is the deployment's own secret, so this table cannot be
 * lifted and reversed into the list of every address anybody has ever tried
 * to sign in as - which is what it would otherwise become, including every
 * typo and every address that has no account here at all.
 *
 * Row-truth, deliberately (ADR-101 R12): this is protocol state read and
 * written on the sign-in path, like TwoFactor's own counters next door. The
 * EVIDENCE an auditor reads is the identity event emitted when a lock is
 * taken, which is durable; these rows are reaped once they stop meaning
 * anything.
 */
class SignInAttemptLockEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SignInAttemptLockEntity>(SignInAttemptLockTable)

    /** Keyed hash of the normalised address. Never the address itself. */
    var identifierHash by SignInAttemptLockTable.identifierHash

    /**
     * Consecutive failures since the last success. Reset by a success, not by
     * time - "five consecutive", not five ever.
     */
    var failedCount by SignInAttemptLockTable.failedCount

    /** When the current lock lifts. Null when not locked. */
    var lockedUntil by SignInAttemptLockTable.lockedUntil

    /**
     * Locks taken since the last success. At five this stops lifting by
     * itself and becomes an incident - see `heldForReview`.
     */
    var consecutiveLockouts by SignInAttemptLockTable.consecutiveLockouts

    /**
     * Held past its lock because the count of locks reached the ceiling the
     * control names. Cleared by an administrator, or by the owner proving
     * mailbox control through a password reset - the second is what keeps a
     * held account from being a denial of service anyone can inflict by
     * typing a colleague's address wrong often enough.
     */
    var heldForReview by SignInAttemptLockTable.heldForReview

    /**
     * Whose account this resolved to at the time, when it resolved to one.
     * Null for an address with no account - those rows are the reason this is
     * nullable, and the reason this table is not hung off User.
     */
    var userId by SignInAttemptLockTable.userId
    var createdAt by SignInAttemptLockTable.createdAt
    var updatedAt by SignInAttemptLockTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SignInAttemptLockTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
