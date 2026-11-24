class LiteLLMAppCustomerMemberTable < ApplicationRecord
  self.table_name = "LiteLLM_AppCustomerMemberTable"
  self.primary_key = %w[app_id customer_member_id]

  attribute :spend, default: 0

  belongs_to :budget, class_name: "LiteLLMBudgetTable", optional: true, inverse_of: :app_member_budgets
  belongs_to :app, class_name: "LiteLLMAppTable", inverse_of: :customer_members
  belongs_to :customer_member, class_name: "LiteLLMCustomerMemberTable", inverse_of: :apps
  belongs_to :budget_override, class_name: "LiteLLMBudgetTable", optional: true, inverse_of: :app_member_overrides
end
