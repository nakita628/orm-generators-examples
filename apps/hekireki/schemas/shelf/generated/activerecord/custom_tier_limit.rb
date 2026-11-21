class CustomTierLimit < ApplicationRecord
  self.table_name = "CustomTierLimit"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :canImportAssets, default: true
  attribute :canExportAssets, default: true
  attribute :canImportNRM, default: true
  attribute :canHideShelfBranding, default: true
  attribute :canUseReports, default: true
  attribute :canUseRecurringReminders, default: true
  attribute :canUsePushNotifications, default: true
  attribute :maxCustomFields, default: 1000
  attribute :maxOrganizations, default: 1
  attribute :isEnterprise, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :userId, uniqueness: { allow_nil: true }

  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :custom_tier_limit
end
