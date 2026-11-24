class LiteLLMAppGroupTable < ApplicationRecord
  self.table_name = "LiteLLM_AppGroupTable"
  self.primary_key = "app_group_id"
  self.implicit_order_column = "created_at"

  attribute :app_group_id, default: -> { SecureRandom.uuid }
  attribute :spend, default: 0

  validates :name, presence: true, uniqueness: true
  validates :created_by, presence: true
  validates :updated_by, presence: true

  belongs_to :budget, class_name: "LiteLLMBudgetTable", optional: true, inverse_of: :app_groups
  has_many :apps, class_name: "LiteLLMAppTable", foreign_key: "app_group_id", inverse_of: :app_group,
                  dependent: :restrict_with_error
  has_many :customer_organizations, class_name: "LiteLLMAppGroupCustomerOrgTable", foreign_key: "app_group_id",
                                    inverse_of: :app_group, dependent: :destroy
end
