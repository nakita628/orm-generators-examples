package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The two-step verification pipeline's projection (D06 - see
 * specs/identity/mfa-and-session-shape.feature). One row per PERSON, because
 * one person has one setup however many organizations they belong to.
 *
 * A pure event-truth head like `SsoConnection`: every column is fold-written,
 * rows are never deleted (DISABLED is a tombstone, so history survives being
 * turned off), and a replay rebuilds it whole-row from the event log. The
 * cursor columns are inline because the fold produces exactly one row per
 * aggregate - the `Identifier` fan-out needed its own cursor table, this
 * does not.
 *
 * Every column here is a lifecycle fact. There is deliberately nowhere in
 * this model a secret or a backup code could be put: `consumedBackupCodeIndexes`
 * holds POSITIONS, which is what lets "how many are left" be answered
 * without the projection ever knowing a code.
 */
class MfaEnrollmentEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, MfaEnrollmentEntity>(MfaEnrollmentTable)

    /** The current enrollment, or NULL for somebody who never started one. */
    var enrollmentId by MfaEnrollmentTable.enrollmentId

    /** totp. Never sms. */
    var method by MfaEnrollmentTable.method

    /** NONE | PENDING | ENABLED | EXPIRED | DISABLED */
    var state by MfaEnrollmentTable.state
    var enrolledAt by MfaEnrollmentTable.enrolledAt
    var confirmedAt by MfaEnrollmentTable.confirmedAt
    var expiredAt by MfaEnrollmentTable.expiredAt
    var disabledAt by MfaEnrollmentTable.disabledAt

    /** password+totp | org-admin - who ended it, and on whose authority. */
    var disabledVia by MfaEnrollmentTable.disabledVia
    var backupCodeCount by MfaEnrollmentTable.backupCodeCount

    /** Positions already spent, ascending. Never codes. */
    var consumedBackupCodeIndexes by MfaEnrollmentTable.consumedBackupCodeIndexes
    var failedCount by MfaEnrollmentTable.failedCount
    var occurredAt by MfaEnrollmentTable.occurredAt
    var lastEventId by MfaEnrollmentTable.lastEventId
    var acceptedAt by MfaEnrollmentTable.acceptedAt
    var projectionVersion by MfaEnrollmentTable.projectionVersion
    var createdAt by MfaEnrollmentTable.createdAt
    var updatedAt by MfaEnrollmentTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(MfaEnrollmentTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
