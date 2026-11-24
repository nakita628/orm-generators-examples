defmodule LiteLLM.LiteLLMProxyModelTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :model_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          model_name: String.t(),
          litellm_params: map() | list() | String.t() | number() | boolean(),
          model_info: map() | list() | String.t() | number() | boolean() | nil,
          blocked: boolean(),
          created_by: String.t(),
          updated_by: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_ProxyModelTable" do
    field(:model_name, :string)
    field(:litellm_params, LiteLLM.PrismaJson)
    field(:model_info, LiteLLM.PrismaJson)
    field(:blocked, :boolean, default: false)
    field(:created_by, :string)
    field(:updated_by, :string)
    timestamps(inserted_at_source: :created_at)
  end
end