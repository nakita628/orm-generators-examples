# maxMembersDeveloper:: Developer seats the organization pays for; null falls to the plan's own cap.
# seatsKept:: Seats paid for this period that stay usable after a seat decrease, until seatsKeptUntil.
# seatLineQuantity:: The seat quantity the last finished seat change set the seat line to; provider events never write it.
# seatLineWriteQuantities:: A Stripe seat line write recorded before it is sent and cleared once it is known to have
#   finished under the lease; one left behind is settled by the seat line repair.
#   Every seat quantity written while this record was pending.
# seatLineWriteSince:: The first write still pending; a later holder's write keeps it.
# seatLineWriteAt:: The latest write sent; the repair waits a grace window after it.
# seatRepairQuantity:: The seat quantity Stripe should bill if a seat change stops after lifting the seat line; null once settled.
# seatRepairLiftedQuantity:: The quantity the change lifted the line to; the repair moves the line back only while Stripe still bills it.
# seatRepairFlaggedAt:: Set once the seat line repair has failed often enough that Ops should look at it.
# seatLineLeaseToken:: The lease a seat change or seat line repair holds while it reads and moves the seat line.
class Subscription < ApplicationRecord
  self.table_name = "Subscription"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :seatLineWriteQuantities, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :plan, { free: "FREE", pro: "PRO", growth: "GROWTH", growth_seat_eur_monthly: "GROWTH_SEAT_EUR_MONTHLY", growth_seat_eur_annual: "GROWTH_SEAT_EUR_ANNUAL", growth_seat_usd_monthly: "GROWTH_SEAT_USD_MONTHLY", growth_seat_usd_annual: "GROWTH_SEAT_USD_ANNUAL", enterprise: "ENTERPRISE", launch: "LAUNCH", accelerate: "ACCELERATE", launch_annual: "LAUNCH_ANNUAL", accelerate_annual: "ACCELERATE_ANNUAL" },
       validate: true
  enum :status, { pending: "PENDING", failed: "FAILED", active: "ACTIVE", cancelled: "CANCELLED" }, validate: true

  validates :stripeSubscriptionId, uniqueness: { allow_nil: true }

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :subscriptions
  has_many :invoices, foreign_key: "subscriptionId", inverse_of: :subscription, dependent: :restrict_with_error
end
