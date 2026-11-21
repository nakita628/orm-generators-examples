package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class ConnectedSeatChangeEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ConnectedSeatChangeEntity>(ConnectedSeatChangeTable)

    /**
     * The reissued `IssuedLicense` row that carries the new seat count. A
     * plain column ref, like `replacesId`: relationMode="prisma" carries no
     * foreign keys.
     */
    var licenseId by ConnectedSeatChangeTable.licenseId

    /**
     * The `ConnectedBillingAccount` row the invoice is raised against. Null
     * when the customer had no billing account (`not_onboarded`).
     */
    var accountId by ConnectedSeatChangeTable.accountId

    /** The customer organization the license belongs to. */
    var organizationId by ConnectedSeatChangeTable.organizationId

    /** When the seats changed, which is the day the proration counts from. */
    var changedAt by ConnectedSeatChangeTable.changedAt

    /** Seats on the new license above the seats on the one it replaced. */
    var addedSeats by ConnectedSeatChangeTable.addedSeats

    /** What one added seat costs for the rest of the term, in cents of `currency`. */
    var unitAmountCents by ConnectedSeatChangeTable.unitAmountCents

    /** `unitAmountCents` times `addedSeats`. Zero when nothing is owed. */
    var amountCents by ConnectedSeatChangeTable.amountCents

    /** Null when no billing account named one (`not_onboarded`). */
    var currency by ConnectedSeatChangeTable.currency

    /** The one-off invoice the added seats were billed on. */
    var stripeInvoiceId by ConnectedSeatChangeTable.stripeInvoiceId
    var state by ConnectedSeatChangeTable.state
    var createdAt by ConnectedSeatChangeTable.createdAt
    var updatedAt by ConnectedSeatChangeTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ConnectedSeatChangeTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
