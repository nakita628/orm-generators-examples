defmodule Langfuse.GatewayApiKeyAssociation do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          api_key_id: String.t(),
          metadata: map() | list() | String.t() | number() | boolean(),
          allowed_models: map() | list() | String.t() | number() | boolean() | nil,
          api_key: Langfuse.ApiKey.t() | nil
        }

  schema "gateway_api_key_associations" do
    field(:metadata, Langfuse.PrismaJson, default: %{})
    field(:allowed_models, Langfuse.PrismaJson)
    field(:api_key_id, :string, primary_key: true)
    belongs_to(:api_key, Langfuse.ApiKey, foreign_key: :api_key_id, define_field: false, type: :string)
  end
end