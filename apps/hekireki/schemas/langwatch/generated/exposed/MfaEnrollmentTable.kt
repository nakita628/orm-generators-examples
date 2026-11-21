package langwatch.models

import org.jetbrains.exposed.v1.core.IntegerColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
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
object MfaEnrollmentTable : IdTable<String>("\"MfaEnrollment\"") {
    /**
     * The user. Also the aggregate id and the tenant - one identity history
     * per person (ADR-029 §4).
     */
    override val id = text("userId").entityId()

    /** The current enrollment, or NULL for somebody who never started one. */
    val enrollmentId = text("enrollmentId").nullable()

    /** totp. Never sms. */
    val method = text("method").nullable()

    /** NONE | PENDING | ENABLED | EXPIRED | DISABLED */
    val state = text("state")
    val enrolledAt = pgTimestamp("enrolledAt", 3).nullable()
    val confirmedAt = pgTimestamp("confirmedAt", 3).nullable()
    val expiredAt = pgTimestamp("expiredAt", 3).nullable()
    val disabledAt = pgTimestamp("disabledAt", 3).nullable()

    /** password+totp | org-admin - who ended it, and on whose authority. */
    val disabledVia = text("disabledVia").nullable()
    val backupCodeCount = integer("backupCodeCount").default(0)

    /** Positions already spent, ascending. Never codes. */
    val consumedBackupCodeIndexes =
        array("consumedBackupCodeIndexes", IntegerColumnType()).nullable().databaseDefault("'{}'") { emptyList() }
    val failedCount = integer("failedCount").default(0)
    val occurredAt = pgTimestamp("occurredAt", 3)
    val lastEventId = text("lastEventId")
    val acceptedAt = pgTimestamp("acceptedAt", 3)
    val projectionVersion = text("projectionVersion")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"MfaEnrollment_pkey\"")

    init {
        index("\"MfaEnrollment_state_idx\"", false, state)
    }
}
