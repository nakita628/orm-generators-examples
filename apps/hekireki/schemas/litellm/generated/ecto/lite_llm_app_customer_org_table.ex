defmodule LiteLLM.LiteLLMAppCustomerOrgTable do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          app_id: String.t(),
          customer_org_id: String.t(),
          spend: float(),
          spend_reset_window: String.t() | nil,
          budget_override_expires_at: DateTime.t() | nil,
          budget: LiteLLM.LiteLLMBudgetTable.t() | nil,
          app: LiteLLM.LiteLLMAppTable.t() | nil,
          customer_org: LiteLLM.LiteLLMCustomerOrganizationTable.t() | nil,
          tier_budget: LiteLLM.LiteLLMBudgetTable.t() | nil,
          budget_override: LiteLLM.LiteLLMBudgetTable.t() | nil,
          customer_member_budget: LiteLLM.LiteLLMBudgetTable.t() | nil
        }

  schema "LiteLLM_AppCustomerOrgTable" do
    field(:spend, :float, default: 0.0)
    field(:spend_reset_window, :string)
    field(:budget_override_expires_at, LiteLLM.PrismaDateTime)
    field(:app_id, :string, primary_key: true)
    field(:customer_org_id, :string, primary_key: true)
    belongs_to(:budget, LiteLLM.LiteLLMBudgetTable, foreign_key: :budget_id, type: :string)
    belongs_to(:app, LiteLLM.LiteLLMAppTable, foreign_key: :app_id, define_field: false, type: :string)
    belongs_to(:customer_org, LiteLLM.LiteLLMCustomerOrganizationTable, foreign_key: :customer_org_id, define_field: false, type: :string)
    belongs_to(:tier_budget, LiteLLM.LiteLLMBudgetTable, foreign_key: :tier_budget_id, type: :string)
    belongs_to(:budget_override, LiteLLM.LiteLLMBudgetTable, foreign_key: :budget_override_id, type: :string)
    belongs_to(:customer_member_budget, LiteLLM.LiteLLMBudgetTable, foreign_key: :customer_member_budget_id, type: :string)
  end
end