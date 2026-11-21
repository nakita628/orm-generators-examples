package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * How a connected customer is invoiced (ADR-139, section 7): one payment
 * provider customer, one quarterly usage subscription in USD, and the
 * commercial terms the annual invoice was raised from.
 *
 * One row per organization. Every write through it is addressed by a stored
 * provider id, so running onboarding or a renewal twice changes nothing.
 */
class ConnectedBillingAccountEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ConnectedBillingAccountEntity>(ConnectedBillingAccountTable)

    var organizationId by ConnectedBillingAccountTable.organizationId

    /** The customer at the payment provider. Invoiced, never charged. */
    var stripeCustomerId by ConnectedBillingAccountTable.stripeCustomerId

    /**
     * The one metered subscription, and the item hosted usage is reported on.
     * Null between the customer being created and the subscription existing,
     * which is where a half-finished onboarding resumes from.
     */
    var usageSubscriptionId by ConnectedBillingAccountTable.usageSubscriptionId
    var usageSubscriptionItemId by ConnectedBillingAccountTable.usageSubscriptionItemId
    var termStartsAt by ConnectedBillingAccountTable.termStartsAt
    var termEndsAt by ConnectedBillingAccountTable.termEndsAt

    /** Prepaid usage for the term, in USD cents. The credit grants add up to it. */
    var commitUsdCents by ConnectedBillingAccountTable.commitUsdCents

    /** The seat contract, which may be in another currency than hosted usage. */
    var seatCurrency by ConnectedBillingAccountTable.seatCurrency
    var seatRateCents by ConnectedBillingAccountTable.seatRateCents
    var seats by ConnectedBillingAccountTable.seats

    /**
     * `us_bank_transfer` or `eu_bank_transfer` when the customer pays into a
     * virtual bank account at the provider. Null when it wires to LangWatch's
     * own account and finance marks the invoice paid.
     */
    var bankTransferType by ConnectedBillingAccountTable.bankTransferType
    var bankTransferCountry by ConnectedBillingAccountTable.bankTransferCountry
    var billingEmail by ConnectedBillingAccountTable.billingEmail

    /**
     * A renewal whose credit is waiting for the last usage invoice of the old
     * term to be finalized: the new term's end and commit, and who asked.
     */
    var pendingRenewal by ConnectedBillingAccountTable.pendingRenewal
    var createdById by ConnectedBillingAccountTable.createdById
    var createdAt by ConnectedBillingAccountTable.createdAt
    var updatedAt by ConnectedBillingAccountTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ConnectedBillingAccountTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
