class LiteLLMBudgetTable < ApplicationRecord
  self.table_name = "LiteLLM_BudgetTable"
  self.primary_key = "budget_id"
  self.implicit_order_column = "created_at"

  attribute :budget_id, default: -> { SecureRandom.uuid }
  attribute :budget_reset_count, default: 0
  attribute :allowed_models, default: -> { [] }

  validates :created_by, presence: true
  validates :updated_by, presence: true

  has_many :apps, class_name: "LiteLLMAppTable", foreign_key: "budget_id", inverse_of: :budget, dependent: :nullify
  has_many :app_groups, class_name: "LiteLLMAppGroupTable", foreign_key: "budget_id", inverse_of: :budget,
                        dependent: :restrict_with_error
  has_many :app_group_org_budgets, class_name: "LiteLLMAppGroupCustomerOrgTable", foreign_key: "budget_id",
                                   inverse_of: :budget, dependent: :restrict_with_error
  has_many :app_default_tiers, class_name: "LiteLLMAppTable", foreign_key: "default_tier_budget_id",
                               inverse_of: :default_tier_budget, dependent: :nullify
  has_many :app_default_members, class_name: "LiteLLMAppTable", foreign_key: "default_customer_member_budget_id",
                                 inverse_of: :default_customer_member_budget, dependent: :nullify
  has_many :app_org_tiers, class_name: "LiteLLMAppCustomerOrgTable", foreign_key: "tier_budget_id",
                           inverse_of: :tier_budget, dependent: :nullify
  has_many :app_org_budgets, class_name: "LiteLLMAppCustomerOrgTable", foreign_key: "budget_id", inverse_of: :budget,
                             dependent: :restrict_with_error
  has_many :app_member_budgets, class_name: "LiteLLMAppCustomerMemberTable", foreign_key: "budget_id",
                                inverse_of: :budget, dependent: :restrict_with_error
  has_many :app_org_overrides, class_name: "LiteLLMAppCustomerOrgTable", foreign_key: "budget_override_id",
                               inverse_of: :budget_override, dependent: :nullify
  has_many :app_org_members, class_name: "LiteLLMAppCustomerOrgTable", foreign_key: "customer_member_budget_id",
                             inverse_of: :customer_member_budget, dependent: :nullify
  has_many :app_member_overrides, class_name: "LiteLLMAppCustomerMemberTable", foreign_key: "budget_override_id",
                                  inverse_of: :budget_override, dependent: :nullify
  has_many :customer_organization_defaults, class_name: "LiteLLMCustomerOrganizationTable",
                                            foreign_key: "default_budget_template_id",
                                            inverse_of: :default_budget_template, dependent: :nullify
  has_many :customer_organizations, class_name: "LiteLLMCustomerOrganizationTable", foreign_key: "budget_id",
                                    inverse_of: :budget, dependent: :nullify
  has_many :customer_teams, class_name: "LiteLLMCustomerTeamTable", foreign_key: "budget_id", inverse_of: :budget,
                            dependent: :nullify
  has_many :customer_members, class_name: "LiteLLMCustomerMemberTable", foreign_key: "budget_id", inverse_of: :budget,
                              dependent: :restrict_with_error
  has_many :customer_member_overrides, class_name: "LiteLLMCustomerMemberTable", foreign_key: "budget_override_id",
                                       inverse_of: :budget_override, dependent: :restrict_with_error
  has_many :customer_team_overrides, class_name: "LiteLLMCustomerTeamTable", foreign_key: "budget_override_id",
                                     inverse_of: :budget_override, dependent: :restrict_with_error
  has_many :organization, class_name: "LiteLLMOrganizationTable", foreign_key: "budget_id",
                          inverse_of: :litellm_budget_table, dependent: :nullify
  has_many :projects, class_name: "LiteLLMProjectTable", foreign_key: "budget_id", inverse_of: :litellm_budget_table,
                      dependent: :nullify
  has_many :keys, class_name: "LiteLLMVerificationToken", foreign_key: "budget_id", inverse_of: :litellm_budget_table,
                  dependent: :nullify
  has_many :end_users, class_name: "LiteLLMEndUserTable", foreign_key: "budget_id", inverse_of: :litellm_budget_table,
                       dependent: :nullify
  has_many :tags, class_name: "LiteLLMTagTable", foreign_key: "budget_id", inverse_of: :litellm_budget_table,
                  dependent: :nullify
  has_many :model_access_groups, class_name: "LiteLLMModelAccessGroupBudgetTable", foreign_key: "budget_id",
                                 inverse_of: :litellm_budget_table, dependent: :nullify
  has_many :team_membership, class_name: "LiteLLMTeamMembership", foreign_key: "budget_id",
                             inverse_of: :litellm_budget_table, dependent: :nullify
  has_many :organization_membership, class_name: "LiteLLMOrganizationMembership", foreign_key: "budget_id",
                                     inverse_of: :litellm_budget_table, dependent: :nullify
end
