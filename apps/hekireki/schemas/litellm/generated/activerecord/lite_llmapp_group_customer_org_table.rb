class LiteLLMAppGroupCustomerOrgTable < ApplicationRecord
  self.table_name = "LiteLLM_AppGroupCustomerOrgTable"
  self.primary_key = %w[app_group_id customer_org_id]

  attribute :spend, default: 0

  belongs_to :budget, class_name: "LiteLLMBudgetTable", optional: true, inverse_of: :app_group_org_budgets
  belongs_to :app_group, class_name: "LiteLLMAppGroupTable", inverse_of: :customer_organizations
  belongs_to :customer_org, class_name: "LiteLLMCustomerOrganizationTable", inverse_of: :app_groups
end
