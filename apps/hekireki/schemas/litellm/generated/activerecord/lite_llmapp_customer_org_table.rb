class LiteLLMAppCustomerOrgTable < ApplicationRecord
  self.table_name = "LiteLLM_AppCustomerOrgTable"
  self.primary_key = %w[app_id customer_org_id]

  attribute :spend, default: 0

  belongs_to :budget, class_name: "LiteLLMBudgetTable", optional: true, inverse_of: :app_org_budgets
  belongs_to :app, class_name: "LiteLLMAppTable", inverse_of: :customer_organizations
  belongs_to :customer_org, class_name: "LiteLLMCustomerOrganizationTable", inverse_of: :apps
  belongs_to :tier_budget, class_name: "LiteLLMBudgetTable", optional: true, inverse_of: :app_org_tiers
  belongs_to :budget_override, class_name: "LiteLLMBudgetTable", optional: true, inverse_of: :app_org_overrides
  belongs_to :customer_member_budget, class_name: "LiteLLMBudgetTable", optional: true, inverse_of: :app_org_members
end
