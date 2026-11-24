defmodule LiteLLM.LiteLLMBackgroundResponseCreator do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          response_id: String.t(),
          upstream_url: String.t(),
          create_context: map() | list() | String.t() | number() | boolean(),
          poll_lease_expires_at: DateTime.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "LiteLLM_BackgroundResponseCreator" do
    field(:response_id, :string, primary_key: true)
    field(:upstream_url, :string)
    field(:create_context, LiteLLM.PrismaJson)
    field(:poll_lease_expires_at, LiteLLM.PrismaDateTime)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end