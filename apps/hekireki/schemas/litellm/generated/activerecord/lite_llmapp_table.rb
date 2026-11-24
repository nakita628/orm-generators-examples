class LiteLLMAppTable < ApplicationRecord
  self.table_name = "LiteLLM_AppTable"
  self.primary_key = "app_id"
  self.implicit_order_column = "created_at"

  attribute :app_id, default: -> { SecureRandom.uuid }
  attribute :spend, default: 0
  attribute :auto_create_customers, default: true

  validates :resource_type,
            uniqueness: { scope: :resource_id, conditions: -> { where.not(resource_id: nil) }, allow_nil: true }
  validates :alias, presence: true, uniqueness: true
  validates :created_by, presence: true
  validates :updated_by, presence: true

  belongs_to :app_group, class_name: "LiteLLMAppGroupTable", optional: true, inverse_of: :apps
  belongs_to :budget, class_name: "LiteLLMBudgetTable", optional: true, inverse_of: :apps
  belongs_to :default_tier_budget, class_name: "LiteLLMBudgetTable", optional: true, inverse_of: :app_default_tiers
  belongs_to :default_customer_member_budget, class_name: "LiteLLMBudgetTable", optional: true,
                                              inverse_of: :app_default_members
  has_many :customer_organizations, class_name: "LiteLLMAppCustomerOrgTable", foreign_key: "app_id", inverse_of: :app,
                                    dependent: :destroy
  has_many :customer_members, class_name: "LiteLLMAppCustomerMemberTable", foreign_key: "app_id", inverse_of: :app,
                              dependent: :destroy
end
