# One renewed licence per paid Stripe invoice of a self-serve licence
# subscription (specs/licensing/self-serving-license-purchase.feature). A
# redelivered invoice event finds the row and mints nothing. A claim never
# delivered is taken over once `claimedAt` is older than the claim lease.
#
# claimToken:: Fences completion and release to the delivery that holds the claim.
# issuedLicense:: The signed licence between minting and mailing, so a retry mails it again instead of minting another; cleared once delivered.
# licenseId:: The `licenseId` inside the delivered licence.
class LicenseRenewalClaim < ApplicationRecord
  self.table_name = "LicenseRenewalClaim"
  self.primary_key = "stripeInvoiceId"

  validates :stripeSubscriptionId, presence: true
  validates :claimToken, presence: true
  validates :claimedAt, presence: true
end
