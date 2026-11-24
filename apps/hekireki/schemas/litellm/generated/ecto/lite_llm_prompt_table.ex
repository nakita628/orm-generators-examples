defmodule LiteLLM.LiteLLMPromptTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          prompt_id: String.t(),
          version: integer(),
          environment: String.t(),
          created_by: String.t() | nil,
          litellm_params: map() | list() | String.t() | number() | boolean(),
          prompt_info: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_PromptTable" do
    field(:prompt_id, :string)
    field(:version, :integer, default: 1)
    field(:environment, :string, default: "development")
    field(:created_by, :string)
    field(:litellm_params, LiteLLM.PrismaJson)
    field(:prompt_info, LiteLLM.PrismaJson)
    timestamps(inserted_at_source: :created_at)
  end
end