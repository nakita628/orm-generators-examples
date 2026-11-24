defmodule LiteLLM.LiteLLMVerificationToken do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          key_id: String.t() | nil,
          token: String.t(),
          key_name: String.t() | nil,
          key_alias: String.t() | nil,
          soft_budget_cooldown: boolean(),
          spend: float(),
          total_spend: float(),
          expires: DateTime.t() | nil,
          models: [String.t()],
          aliases: map() | list() | String.t() | number() | boolean(),
          config: map() | list() | String.t() | number() | boolean(),
          router_settings: map() | list() | String.t() | number() | boolean() | nil,
          user_id: String.t() | nil,
          team_id: String.t() | nil,
          agent_id: String.t() | nil,
          permissions: map() | list() | String.t() | number() | boolean(),
          max_parallel_requests: integer() | nil,
          metadata: map() | list() | String.t() | number() | boolean(),
          blocked: boolean() | nil,
          tpm_limit: integer() | nil,
          rpm_limit: integer() | nil,
          tpd_limit: integer() | nil,
          max_budget: float() | nil,
          budget_duration: String.t() | nil,
          budget_reset_at: DateTime.t() | nil,
          allowed_cache_controls: [String.t()],
          allowed_routes: [String.t()],
          key_type: String.t() | nil,
          policies: [String.t()],
          access_group_ids: [String.t()],
          model_spend: map() | list() | String.t() | number() | boolean(),
          model_max_budget: map() | list() | String.t() | number() | boolean(),
          budget_fallbacks: map() | list() | String.t() | number() | boolean(),
          created_by: String.t() | nil,
          updated_by: String.t() | nil,
          settings_updated_at: DateTime.t() | nil,
          last_active: DateTime.t() | nil,
          rotation_count: integer() | nil,
          auto_rotate: boolean() | nil,
          rotation_interval: String.t() | nil,
          last_rotation_at: DateTime.t() | nil,
          key_rotation_at: DateTime.t() | nil,
          budget_limits: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t() | nil,
          updated_at: DateTime.t() | nil,
          litellm_budget_table: LiteLLM.LiteLLMBudgetTable.t() | nil,
          litellm_organization_table: LiteLLM.LiteLLMOrganizationTable.t() | nil,
          litellm_project_table: LiteLLM.LiteLLMProjectTable.t() | nil,
          object_permission: LiteLLM.LiteLLMObjectPermissionTable.t() | nil,
          jwt_key_mappings: [LiteLLM.LiteLLMJWTKeyMapping.t()]
        }

  schema "LiteLLM_VerificationToken" do
    field(:key_id, :string, autogenerate: {Ecto.UUID, :generate, []})
    field(:token, :string, primary_key: true)
    field(:key_name, :string)
    field(:key_alias, :string)
    field(:soft_budget_cooldown, :boolean, default: false)
    field(:spend, :float, default: 0.0)
    field(:total_spend, :float, default: 0.0)
    field(:expires, LiteLLM.PrismaDateTime)
    field(:models, {:array, :string})
    field(:aliases, LiteLLM.PrismaJson, default: %{})
    field(:config, LiteLLM.PrismaJson, default: %{})
    field(:router_settings, LiteLLM.PrismaJson, default: %{})
    field(:user_id, :string)
    field(:team_id, :string)
    field(:agent_id, :string)
    field(:permissions, LiteLLM.PrismaJson, default: %{})
    field(:max_parallel_requests, :integer)
    field(:metadata, LiteLLM.PrismaJson, default: %{})
    field(:blocked, :boolean)
    field(:tpm_limit, :integer)
    field(:rpm_limit, :integer)
    field(:tpd_limit, :integer)
    field(:max_budget, :float)
    field(:budget_duration, :string)
    field(:budget_reset_at, LiteLLM.PrismaDateTime)
    field(:allowed_cache_controls, {:array, :string})
    field(:allowed_routes, {:array, :string})
    field(:key_type, :string)
    field(:policies, {:array, :string})
    field(:access_group_ids, {:array, :string})
    field(:model_spend, LiteLLM.PrismaJson, default: %{})
    field(:model_max_budget, LiteLLM.PrismaJson, default: %{})
    field(:budget_fallbacks, LiteLLM.PrismaJson, default: %{})
    field(:created_by, :string)
    field(:updated_by, :string)
    field(:settings_updated_at, LiteLLM.PrismaDateTime)
    field(:last_active, LiteLLM.PrismaDateTime)
    field(:rotation_count, :integer, default: 0)
    field(:auto_rotate, :boolean, default: false)
    field(:rotation_interval, :string)
    field(:last_rotation_at, LiteLLM.PrismaDateTime)
    field(:key_rotation_at, LiteLLM.PrismaDateTime)
    field(:budget_limits, LiteLLM.PrismaJson)
    belongs_to(:litellm_budget_table, LiteLLM.LiteLLMBudgetTable, foreign_key: :budget_id, type: :string)
    belongs_to(:litellm_organization_table, LiteLLM.LiteLLMOrganizationTable, foreign_key: :organization_id, type: :string)
    belongs_to(:litellm_project_table, LiteLLM.LiteLLMProjectTable, foreign_key: :project_id, type: :string)
    belongs_to(:object_permission, LiteLLM.LiteLLMObjectPermissionTable, foreign_key: :object_permission_id, type: :string)
    has_many(:jwt_key_mappings, LiteLLM.LiteLLMJWTKeyMapping, foreign_key: :token, references: :token)
    timestamps(inserted_at_source: :created_at)
  end
end