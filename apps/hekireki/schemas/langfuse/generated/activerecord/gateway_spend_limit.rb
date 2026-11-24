class GatewaySpendLimit < ApplicationRecord
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :alert_thresholds, default: -> { [ 100 ] }

  enum :scope, { api_key: "API_KEY", user: "USER", api_key_default: "API_KEY_DEFAULT", user_default: "USER_DEFAULT", organization: "ORGANIZATION" },
       validate: true
  enum :interval, { daily: "DAILY", weekly: "WEEKLY", monthly: "MONTHLY", total: "TOTAL" }, validate: true

  validates :organization_id, uniqueness: { scope: %i[user_id interval], conditions: -> { where.not(user_id: nil) } }
  validates :api_key_id, uniqueness: { scope: :interval, allow_nil: true }
  validates :limit_usd, presence: true
  validates :counting_started_at, presence: true

  belongs_to :organization, inverse_of: :gateway_spend_limits
  belongs_to :api_key, optional: true, inverse_of: :gateway_spend_limits
  belongs_to :user, optional: true, inverse_of: :gateway_spend_limits
  belongs_to :owner_user, class_name: "User", optional: true, inverse_of: :gateway_spend_limits_owned
  has_many :alerts, class_name: "GatewaySpendLimitAlert", foreign_key: "spend_limit_id", inverse_of: :spend_limit,
                    dependent: :destroy
end
