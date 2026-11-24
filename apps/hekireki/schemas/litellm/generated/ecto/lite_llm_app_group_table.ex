defmodule LiteLLM.LiteLLMAppGroupTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :app_group_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          metadata: map() | list() | String.t() | number() | boolean() | nil,
          spend: float(),
          spend_reset_window: String.t() | nil,
          created_by: String.t(),
          updated_by: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          budget: LiteLLM.LiteLLMBudgetTable.t() | nil,
          apps: [LiteLLM.LiteLLMAppTable.t()],
          customer_organizations: [LiteLLM.LiteLLMAppGroupCustomerOrgTable.t()]
        }

  schema "LiteLLM_AppGroupTable" do
    field(:name, :string)
    field(:description, :string)
    field(:metadata, LiteLLM.PrismaJson)
    field(:spend, :float, default: 0.0)
    field(:spend_reset_window, :string)
    field(:created_by, :string)
    field(:updated_by, :string)
    belongs_to(:budget, LiteLLM.LiteLLMBudgetTable, foreign_key: :budget_id)
    has_many(:apps, LiteLLM.LiteLLMAppTable, foreign_key: :app_group_id)
    has_many(:customer_organizations, LiteLLM.LiteLLMAppGroupCustomerOrgTable, foreign_key: :app_group_id)
    timestamps(inserted_at_source: :created_at)
  end
end