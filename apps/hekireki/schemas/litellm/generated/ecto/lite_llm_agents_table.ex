defmodule LiteLLM.LiteLLMAgentsTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :agent_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          agent_name: String.t(),
          litellm_params: map() | list() | String.t() | number() | boolean() | nil,
          agent_card_params: map() | list() | String.t() | number() | boolean(),
          static_headers: map() | list() | String.t() | number() | boolean() | nil,
          extra_headers: [String.t()],
          kill_switch: map() | list() | String.t() | number() | boolean() | nil,
          agent_access_groups: [String.t()],
          access_group_ids: [String.t()],
          spend: float(),
          identity_managed: boolean(),
          enabled: boolean(),
          execution_mode: String.t(),
          tpm_limit: integer() | nil,
          rpm_limit: integer() | nil,
          session_tpm_limit: integer() | nil,
          session_rpm_limit: integer() | nil,
          created_by: String.t(),
          updated_by: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          object_permission: LiteLLM.LiteLLMObjectPermissionTable.t() | nil,
          identity: LiteLLM.LiteLLMAgentIdentity.t() | nil,
          retired_identities: [LiteLLM.LiteLLMRetiredAgentIdentity.t()]
        }

  schema "LiteLLM_AgentsTable" do
    field(:agent_name, :string)
    field(:litellm_params, LiteLLM.PrismaJson)
    field(:agent_card_params, LiteLLM.PrismaJson)
    field(:static_headers, LiteLLM.PrismaJson, default: %{})
    field(:extra_headers, {:array, :string})
    field(:kill_switch, LiteLLM.PrismaJson)
    field(:agent_access_groups, {:array, :string})
    field(:access_group_ids, {:array, :string})
    field(:spend, :float, default: 0.0)
    field(:identity_managed, :boolean, default: false)
    field(:enabled, :boolean, default: true)
    field(:execution_mode, :string, default: "autonomous")
    field(:tpm_limit, :integer)
    field(:rpm_limit, :integer)
    field(:session_tpm_limit, :integer)
    field(:session_rpm_limit, :integer)
    field(:created_by, :string)
    field(:updated_by, :string)
    belongs_to(:object_permission, LiteLLM.LiteLLMObjectPermissionTable, foreign_key: :object_permission_id)
    has_one(:identity, LiteLLM.LiteLLMAgentIdentity, foreign_key: :agent_id)
    has_many(:retired_identities, LiteLLM.LiteLLMRetiredAgentIdentity, foreign_key: :agent_id)
    timestamps(inserted_at_source: :created_at)
  end
end