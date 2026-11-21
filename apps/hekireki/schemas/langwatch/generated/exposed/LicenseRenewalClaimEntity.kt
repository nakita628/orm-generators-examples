package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One renewed licence per paid Stripe invoice of a self-serve licence
 * subscription (specs/licensing/self-serving-license-purchase.feature). A
 * redelivered invoice event finds the row and mints nothing. A claim never
 * delivered is taken over once `claimedAt` is older than the claim lease.
 */
class LicenseRenewalClaimEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, LicenseRenewalClaimEntity>(LicenseRenewalClaimTable)

    var stripeSubscriptionId by LicenseRenewalClaimTable.stripeSubscriptionId

    /** Fences completion and release to the delivery that holds the claim. */
    var claimToken by LicenseRenewalClaimTable.claimToken
    var claimedAt by LicenseRenewalClaimTable.claimedAt

    /** The signed licence between minting and mailing, so a retry mails it again instead of minting another; cleared once delivered. */
    var issuedLicense by LicenseRenewalClaimTable.issuedLicense
    var deliveredAt by LicenseRenewalClaimTable.deliveredAt

    /** The `licenseId` inside the delivered licence. */
    var licenseId by LicenseRenewalClaimTable.licenseId
}
