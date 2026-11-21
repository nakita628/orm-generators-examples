# What a mid-term seat change of one connected license owes (ADR-141,
# section 7): the seats added on the reissued license, invoiced prorated to
# the end of the term.
#
# The row is written before the payment provider is called and completed
# after it, so a call that failed halfway is retried by the next seat
# invoicing pass without inventing a second invoice. One row per reissued license.
#
# licenseId:: The reissued `IssuedLicense` row that carries the new seat count. A
#   plain column ref, like `replacesId`: relationMode="prisma" carries no
#   foreign keys.
# accountId:: The `ConnectedBillingAccount` row the invoice is raised against. Null
#   when the customer had no billing account (`not_onboarded`).
# organizationId:: The customer organization the license belongs to.
# changedAt:: When the seats changed, which is the day the proration counts from.
# addedSeats:: Seats on the new license above the seats on the one it replaced.
# unitAmountCents:: What one added seat costs for the rest of the term, in cents of `currency`.
# amountCents:: `unitAmountCents` times `addedSeats`. Zero when nothing is owed.
# currency:: Null when no billing account named one (`not_onboarded`).
# stripeInvoiceId:: The one-off invoice the added seats were billed on.
class ConnectedSeatChange < ApplicationRecord
  self.table_name = "ConnectedSeatChange"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :currency, { usd: "USD", eur: "EUR" }, validate: { allow_nil: true }
  enum :state, { intent: "intent", invoiced: "invoiced", nothing_to_invoice: "nothing_to_invoice", not_onboarded: "not_onboarded" },
       validate: true

  validates :licenseId, presence: true, uniqueness: true
  validates :changedAt, presence: true
  validates :addedSeats, presence: true
  validates :unitAmountCents, presence: true
  validates :amountCents, presence: true
end
