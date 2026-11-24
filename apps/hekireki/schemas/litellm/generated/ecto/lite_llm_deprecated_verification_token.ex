defmodule LiteLLM.LiteLLMDeprecatedVerificationToken do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          token: String.t(),
          active_token_id: String.t(),
          revoke_at: DateTime.t(),
          inserted_at: DateTime.t()
        }

  schema "LiteLLM_DeprecatedVerificationToken" do
    field(:token, :string)
    field(:active_token_id, :string)
    field(:revoke_at, LiteLLM.PrismaDateTime)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end