package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/**
 * One renewed licence per paid Stripe invoice of a self-serve licence
 * subscription (specs/licensing/self-serving-license-purchase.feature). A
 * redelivered invoice event finds the row and mints nothing. A claim never
 * delivered is taken over once `claimedAt` is older than the claim lease.
 */
object LicenseRenewalClaimTable : IdTable<String>("\"LicenseRenewalClaim\"") {
    override val id = text("stripeInvoiceId").entityId()
    val stripeSubscriptionId = text("stripeSubscriptionId")

    /** Fences completion and release to the delivery that holds the claim. */
    val claimToken = text("claimToken")
    val claimedAt = pgTimestamp("claimedAt", 3)

    /** The signed licence between minting and mailing, so a retry mails it again instead of minting another; cleared once delivered. */
    val issuedLicense = jsonb<String>("issuedLicense", { it }, { it }).nullable()
    val deliveredAt = pgTimestamp("deliveredAt", 3).nullable()

    /** The `licenseId` inside the delivered licence. */
    val licenseId = text("licenseId").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"LicenseRenewalClaim_pkey\"")
}
