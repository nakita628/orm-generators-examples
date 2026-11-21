# *
# * Audit trail for all quantity changes on quantity-tracked assets.
# * Every event records who performed it, who received/returned items,
# * which booking it relates to, and the category of change.
# * Quantity is always positive — direction is determined by category.
class ConsumptionLog < ApplicationRecord
  self.table_name = "ConsumptionLog"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt

  enum :category, { checkout: "CHECKOUT", return: "RETURN", restock: "RESTOCK", adjustment: "ADJUSTMENT", loss: "LOSS", consume: "CONSUME", damage: "DAMAGE", move: "MOVE", initial: "INITIAL", opening_balance: "OPENING_BALANCE" },
       validate: true

  validates :quantity, presence: true

  belongs_to :asset, foreign_key: "assetId", inverse_of: :consumption_logs
  belongs_to :performed_by, class_name: "User", foreign_key: "userId", inverse_of: :consumption_logs
  belongs_to :booking, foreign_key: "bookingId", optional: true, inverse_of: :consumption_logs
  belongs_to :custodian, class_name: "TeamMember", foreign_key: "custodianId", optional: true,
                         inverse_of: :consumption_logs
  belongs_to :location, foreign_key: "locationId", optional: true, inverse_of: :consumption_logs
end
