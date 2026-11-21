# How a connected customer is invoiced (ADR-139, section 7): one payment
# provider customer, one quarterly usage subscription in USD, and the
# commercial terms the annual invoice was raised from.
#
# One row per organization. Every write through it is addressed by a stored
# provider id, so running onboarding or a renewal twice changes nothing.
#
# stripeCustomerId:: The customer at the payment provider. Invoiced, never charged.
# usageSubscriptionId:: The one metered subscription, and the item hosted usage is reported on.
#   Null between the customer being created and the subscription existing,
#   which is where a half-finished onboarding resumes from.
# commitUsdCents:: Prepaid usage for the term, in USD cents. The credit grants add up to it.
# seatCurrency:: The seat contract, which may be in another currency than hosted usage.
# bankTransferType:: `us_bank_transfer` or `eu_bank_transfer` when the customer pays into a
#   virtual bank account at the provider. Null when it wires to LangWatch's
#   own account and finance marks the invoice paid.
# pendingRenewal:: A renewal whose credit is waiting for the last usage invoice of the old
#   term to be finalized: the new term's end and commit, and who asked.
class ConnectedBillingAccount < ApplicationRecord
  self.table_name = "ConnectedBillingAccount"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :commitUsdCents, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :seatCurrency, { usd: "USD", eur: "EUR" }, validate: true

  validates :organizationId, uniqueness: true
  validates :stripeCustomerId, presence: true, uniqueness: true
  validates :termStartsAt, presence: true
  validates :termEndsAt, presence: true
  validates :seatRateCents, presence: true
  validates :seats, presence: true
  validates :billingEmail, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :connected_billing
end
