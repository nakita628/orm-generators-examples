defmodule LiteLLM.LiteLLMMCPUserEnvVars do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          user_id: String.t(),
          server_id: String.t(),
          values_b64: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_MCPUserEnvVars" do
    field(:user_id, :string)
    field(:server_id, :string)
    field(:values_b64, :string)
    timestamps(inserted_at_source: :created_at)
  end
end