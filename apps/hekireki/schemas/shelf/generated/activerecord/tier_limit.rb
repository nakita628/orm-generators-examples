class TierLimit < ApplicationRecord
  self.table_name = "TierLimit"

  attribute :canImportAssets, default: false
  attribute :canExportAssets, default: false
  attribute :canImportNRM, default: false
  attribute :canHideShelfBranding, default: false
  attribute :canUseReports, default: false
  attribute :canUseRecurringReminders, default: false
  attribute :canUsePushNotifications, default: false
  attribute :maxCustomFields, default: 0
  attribute :maxOrganizations, default: 1
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :id, inclusion: { in: %w[free tier_1 tier_2 custom] }

  has_one :tier, foreign_key: "tierLimitId", inverse_of: :tier_limit, dependent: :nullify
end
