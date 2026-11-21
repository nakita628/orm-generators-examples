class Tier < ApplicationRecord
  self.table_name = "Tier"

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :id, inclusion: { in: %w[free tier_1 tier_2 custom] }
  enum :tierLimitId, { free: "free", tier_1: "tier_1", tier_2: "tier_2", custom: "custom" },
       validate: { allow_nil: true }

  validates :name, presence: true

  belongs_to :tier_limit, foreign_key: "tierLimitId", optional: true, inverse_of: :tier
  has_many :subscribers, class_name: "User", foreign_key: "tierId", inverse_of: :tier, dependent: :restrict_with_error
end
