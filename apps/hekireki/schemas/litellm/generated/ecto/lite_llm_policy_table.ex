defmodule LiteLLM.LiteLLMPolicyTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :policy_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          policy_name: String.t(),
          version_number: integer(),
          version_status: String.t(),
          parent_version_id: String.t() | nil,
          is_latest: boolean(),
          published_at: DateTime.t() | nil,
          production_at: DateTime.t() | nil,
          inherit: String.t() | nil,
          description: String.t() | nil,
          guardrails_add: [String.t()],
          guardrails_remove: [String.t()],
          condition: map() | list() | String.t() | number() | boolean() | nil,
          pipeline: map() | list() | String.t() | number() | boolean() | nil,
          created_by: String.t() | nil,
          updated_by: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_PolicyTable" do
    field(:policy_name, :string)
    field(:version_number, :integer, default: 1)
    field(:version_status, :string, default: "production")
    field(:parent_version_id, :string)
    field(:is_latest, :boolean, default: true)
    field(:published_at, LiteLLM.PrismaDateTime)
    field(:production_at, LiteLLM.PrismaDateTime)
    field(:inherit, :string)
    field(:description, :string)
    field(:guardrails_add, {:array, :string})
    field(:guardrails_remove, {:array, :string})
    field(:condition, LiteLLM.PrismaJson, default: %{})
    field(:pipeline, LiteLLM.PrismaJson)
    field(:created_by, :string)
    field(:updated_by, :string)
    timestamps(inserted_at_source: :created_at)
  end
end