package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
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
object ConnectedBillingAccountTable : IdTable<String>("\"ConnectedBillingAccount\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId").uniqueIndex("\"ConnectedBillingAccount_organizationId_key\"")

    /** The customer at the payment provider. Invoiced, never charged. */
    val stripeCustomerId = text("stripeCustomerId").uniqueIndex("\"ConnectedBillingAccount_stripeCustomerId_key\"")

    /**
     * The one metered subscription, and the item hosted usage is reported on.
     * Null between the customer being created and the subscription existing,
     * which is where a half-finished onboarding resumes from.
     */
    val usageSubscriptionId = text("usageSubscriptionId").nullable()
    val usageSubscriptionItemId = text("usageSubscriptionItemId").nullable()
    val termStartsAt = pgTimestamp("termStartsAt", 3)
    val termEndsAt = pgTimestamp("termEndsAt", 3)

    /** Prepaid usage for the term, in USD cents. The credit grants add up to it. */
    val commitUsdCents = integer("commitUsdCents").default(0)

    /** The seat contract, which may be in another currency than hosted usage. */
    val seatCurrency = pgEnum("seatCurrency", "\"Currency\"", Currency.entries, Currency::dbName)
    val seatRateCents = integer("seatRateCents")
    val seats = integer("seats")

    /**
     * `us_bank_transfer` or `eu_bank_transfer` when the customer pays into a
     * virtual bank account at the provider. Null when it wires to LangWatch's
     * own account and finance marks the invoice paid.
     */
    val bankTransferType = text("bankTransferType").nullable()
    val bankTransferCountry = text("bankTransferCountry").nullable()
    val billingEmail = text("billingEmail")

    /**
     * A renewal whose credit is waiting for the last usage invoice of the old
     * term to be finalized: the new term's end and commit, and who asked.
     */
    val pendingRenewal = jsonb<String>("pendingRenewal", { it }, { it }).nullable()
    val createdById = text("createdById").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"ConnectedBillingAccount_pkey\"")
}
