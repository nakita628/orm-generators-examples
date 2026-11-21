package langwatch.models

import org.jetbrains.exposed.v1.core.IntegerColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A way back in (D05 - see specs/identity/sso-onboarding-tiers.feature).
 *
 * One named person who may still sign in without the identity provider
 * while their organization's sign-in belongs to one, and the date on which
 * that stops. Activation asks whether such a person exists, because the
 * failure it prevents is the only one that cannot be recovered from inside
 * the product: a misconfigured connection goes live, every member is handed
 * to an identity provider that will not authenticate them, and nobody is
 * left who can turn it off.
 *
 * Rows are immutable except for `supersededAt` and `warnedDays`. A renewal
 * writes a NEW row naming the one it replaced, so the date a way in
 * previously ended is still readable in the history - which is what makes a
 * renewal auditable rather than a column somebody moved.
 *
 * No relation to Organization or User, matching the identity heads above:
 * the binding is the record of who could get in, and a cascade would erase
 * exactly the history a lockout post-mortem reads.
 */
object SsoBreakGlassBindingTable : IdTable<String>("\"SsoBreakGlassBinding\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")

    /** Who holds the way in. */
    val userId = text("userId")

    /** Who granted it. A way back in is never self-served. */
    val grantedByUserId = text("grantedByUserId")
    val grantedAt = pgTimestamp("grantedAt", 3)
    val expiresAt = pgTimestamp("expiresAt", 3)

    /** Set when a renewal replaced this row; the row itself is never edited. */
    val supersededAt = pgTimestamp("supersededAt", 3).nullable()

    /** The binding this one renewed, so the previous end date stays readable. */
    val renewedFromId = text("renewedFromId").nullable()

    /**
     * Which of the fourteen, seven and one day warnings have been sent, so a
     * sweep that runs twice in a day never says the same thing twice.
     */
    val warnedDays = array("warnedDays", IntegerColumnType()).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"SsoBreakGlassBinding_pkey\"")

    init {
        index("\"SsoBreakGlassBinding_organizationId_expiresAt_idx\"", false, organizationId, expiresAt)
        index("\"SsoBreakGlassBinding_userId_idx\"", false, userId)
    }
}
