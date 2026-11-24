defmodule LiteLLM.LiteLLMGuardrailsTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :guardrail_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          guardrail_name: String.t(),
          litellm_params: map() | list() | String.t() | number() | boolean(),
          guardrail_info: map() | list() | String.t() | number() | boolean() | nil,
          team_id: String.t() | nil,
          status: String.t(),
          submitted_at: DateTime.t() | nil,
          reviewed_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_GuardrailsTable" do
    field(:guardrail_name, :string)
    field(:litellm_params, LiteLLM.PrismaJson)
    field(:guardrail_info, LiteLLM.PrismaJson)
    field(:team_id, :string)
    field(:status, :string, default: "active")
    field(:submitted_at, LiteLLM.PrismaDateTime)
    field(:reviewed_at, LiteLLM.PrismaDateTime)
    timestamps(inserted_at_source: :created_at)
  end
end