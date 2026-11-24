defmodule LiteLLM.LiteLLMAppGroupCustomerOrgTable do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          app_group_id: String.t(),
          customer_org_id: String.t(),
          spend: float(),
          spend_reset_window: String.t() | nil,
          budget: LiteLLM.LiteLLMBudgetTable.t() | nil,
          app_group: LiteLLM.LiteLLMAppGroupTable.t() | nil,
          customer_org: LiteLLM.LiteLLMCustomerOrganizationTable.t() | nil
        }

  schema "LiteLLM_AppGroupCustomerOrgTable" do
    field(:spend, :float, default: 0.0)
    field(:spend_reset_window, :string)
    field(:app_group_id, :string, primary_key: true)
    field(:customer_org_id, :string, primary_key: true)
    belongs_to(:budget, LiteLLM.LiteLLMBudgetTable, foreign_key: :budget_id, type: :string)
    belongs_to(:app_group, LiteLLM.LiteLLMAppGroupTable, foreign_key: :app_group_id, define_field: false, type: :string)
    belongs_to(:customer_org, LiteLLM.LiteLLMCustomerOrganizationTable, foreign_key: :customer_org_id, define_field: false, type: :string)
  end
end