defmodule LiteLLM.LiteLLMAppTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :app_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          resource_type: String.t() | nil,
          resource_id: String.t() | nil,
          alias: String.t(),
          spend: float(),
          spend_reset_window: String.t() | nil,
          auto_create_customers: boolean(),
          customer_margin: float() | nil,
          created_by: String.t(),
          updated_by: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          app_group: LiteLLM.LiteLLMAppGroupTable.t() | nil,
          budget: LiteLLM.LiteLLMBudgetTable.t() | nil,
          default_tier_budget: LiteLLM.LiteLLMBudgetTable.t() | nil,
          default_customer_member_budget: LiteLLM.LiteLLMBudgetTable.t() | nil,
          customer_organizations: [LiteLLM.LiteLLMAppCustomerOrgTable.t()],
          customer_members: [LiteLLM.LiteLLMAppCustomerMemberTable.t()]
        }

  schema "LiteLLM_AppTable" do
    field(:resource_type, :string)
    field(:resource_id, :string)
    field(:alias, :string)
    field(:spend, :float, default: 0.0)
    field(:spend_reset_window, :string)
    field(:auto_create_customers, :boolean, default: true)
    field(:customer_margin, :float)
    field(:created_by, :string)
    field(:updated_by, :string)
    belongs_to(:app_group, LiteLLM.LiteLLMAppGroupTable, foreign_key: :app_group_id)
    belongs_to(:budget, LiteLLM.LiteLLMBudgetTable, foreign_key: :budget_id)
    belongs_to(:default_tier_budget, LiteLLM.LiteLLMBudgetTable, foreign_key: :default_tier_budget_id)
    belongs_to(:default_customer_member_budget, LiteLLM.LiteLLMBudgetTable, foreign_key: :default_customer_member_budget_id)
    has_many(:customer_organizations, LiteLLM.LiteLLMAppCustomerOrgTable, foreign_key: :app_id)
    has_many(:customer_members, LiteLLM.LiteLLMAppCustomerMemberTable, foreign_key: :app_id)
    timestamps(inserted_at_source: :created_at)
  end
end