# One notification per limit, subject, alert type, threshold, counting start and window, so the gateway may report a limit any number of times.
class GatewaySpendLimitAlert < ApplicationRecord
  self.inheritance_column = nil
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :subject_id, default: ""
  attribute :threshold_percent, default: 100

  enum :type, { warning: "WARNING", exceeded: "EXCEEDED" }, validate: true

  validates :subject_id,
            uniqueness: { scope: %i[spend_limit_id type threshold_percent counting_started_at window_start] }
  validates :counting_started_at, presence: true
  validates :window_start, presence: true

  belongs_to :spend_limit, class_name: "GatewaySpendLimit", inverse_of: :alerts
end
