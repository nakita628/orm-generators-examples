defmodule LiteLLM.LiteLLMProjectTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :project_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          project_alias: String.t() | nil,
          description: String.t() | nil,
          metadata: map() | list() | String.t() | number() | boolean(),
          models: [String.t()],
          spend: float(),
          model_spend: map() | list() | String.t() | number() | boolean(),
          model_rpm_limit: map() | list() | String.t() | number() | boolean(),
          model_tpm_limit: map() | list() | String.t() | number() | boolean(),
          blocked: boolean(),
          created_by: String.t(),
          updated_by: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          litellm_team_table: LiteLLM.LiteLLMTeamTable.t() | nil,
          litellm_budget_table: LiteLLM.LiteLLMBudgetTable.t() | nil,
          object_permission: LiteLLM.LiteLLMObjectPermissionTable.t() | nil,
          keys: [LiteLLM.LiteLLMVerificationToken.t()]
        }

  schema "LiteLLM_ProjectTable" do
    field(:project_alias, :string)
    field(:description, :string)
    field(:metadata, LiteLLM.PrismaJson, default: %{})
    field(:models, {:array, :string})
    field(:spend, :float, default: 0.0)
    field(:model_spend, LiteLLM.PrismaJson, default: %{})
    field(:model_rpm_limit, LiteLLM.PrismaJson, default: %{})
    field(:model_tpm_limit, LiteLLM.PrismaJson, default: %{})
    field(:blocked, :boolean, default: false)
    field(:created_by, :string)
    field(:updated_by, :string)
    belongs_to(:litellm_team_table, LiteLLM.LiteLLMTeamTable, foreign_key: :team_id)
    belongs_to(:litellm_budget_table, LiteLLM.LiteLLMBudgetTable, foreign_key: :budget_id)
    belongs_to(:object_permission, LiteLLM.LiteLLMObjectPermissionTable, foreign_key: :object_permission_id)
    has_many(:keys, LiteLLM.LiteLLMVerificationToken, foreign_key: :project_id)
    timestamps(inserted_at_source: :created_at)
  end
end