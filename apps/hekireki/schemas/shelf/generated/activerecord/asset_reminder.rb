class AssetReminder < ApplicationRecord
  self.table_name = "AssetReminder"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :recurrenceUnit, { day: "DAY", week: "WEEK", month: "MONTH", year: "YEAR" }, validate: { allow_nil: true }

  validates :name, presence: true
  validates :message, presence: true
  validates :alertDateTime, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :asset_reminders
  belongs_to :asset, foreign_key: "assetId", inverse_of: :reminders
  belongs_to :created_by, class_name: "User", foreign_key: "createdById", inverse_of: :asset_reminders
  has_and_belongs_to_many :team_members, join_table: "_AssetReminderToTeamMember", foreign_key: "A",
                                         association_foreign_key: "B"
end
