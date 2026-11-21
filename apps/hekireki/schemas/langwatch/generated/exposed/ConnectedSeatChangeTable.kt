package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * What a mid-term seat change of one connected license owes (ADR-141,
 * section 7): the seats added on the reissued license, invoiced prorated to
 * the end of the term.
 *
 * The row is written before the payment provider is called and completed
 * after it, so a call that failed halfway is retried by the next seat
 * invoicing pass without inventing a second invoice. One row per reissued license.
 */
object ConnectedSeatChangeTable : IdTable<String>("\"ConnectedSeatChange\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()

    /**
     * The reissued `IssuedLicense` row that carries the new seat count. A
     * plain column ref, like `replacesId`: relationMode="prisma" carries no
     * foreign keys.
     */
    val licenseId = text("licenseId").uniqueIndex("\"ConnectedSeatChange_licenseId_key\"")

    /**
     * The `ConnectedBillingAccount` row the invoice is raised against. Null
     * when the customer had no billing account (`not_onboarded`).
     */
    val accountId = text("accountId").nullable()

    /** The customer organization the license belongs to. */
    val organizationId = text("organizationId").nullable()

    /** When the seats changed, which is the day the proration counts from. */
    val changedAt = pgTimestamp("changedAt", 3)

    /** Seats on the new license above the seats on the one it replaced. */
    val addedSeats = integer("addedSeats")

    /** What one added seat costs for the rest of the term, in cents of `currency`. */
    val unitAmountCents = integer("unitAmountCents")

    /** `unitAmountCents` times `addedSeats`. Zero when nothing is owed. */
    val amountCents = integer("amountCents")

    /** Null when no billing account named one (`not_onboarded`). */
    val currency = pgEnum("currency", "\"Currency\"", Currency.entries, Currency::dbName).nullable()

    /** The one-off invoice the added seats were billed on. */
    val stripeInvoiceId = text("stripeInvoiceId").nullable()
    val state = pgEnum("state", "\"ConnectedSeatChangeState\"", ConnectedSeatChangeState.entries, ConnectedSeatChangeState::dbName)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"ConnectedSeatChange_pkey\"")

    init {
        index("\"ConnectedSeatChange_accountId_idx\"", false, accountId)
        index("\"ConnectedSeatChange_organizationId_idx\"", false, organizationId)
        index("\"ConnectedSeatChange_state_idx\"", false, state)
    }
}
