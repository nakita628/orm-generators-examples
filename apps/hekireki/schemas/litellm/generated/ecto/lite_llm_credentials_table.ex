defmodule LiteLLM.LiteLLMCredentialsTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :credential_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          credential_name: String.t(),
          display_name: String.t() | nil,
          credential_values: map() | list() | String.t() | number() | boolean(),
          credential_info: map() | list() | String.t() | number() | boolean() | nil,
          created_by: String.t(),
          updated_by: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_CredentialsTable" do
    field(:credential_name, :string)
    field(:display_name, :string)
    field(:credential_values, LiteLLM.PrismaJson)
    field(:credential_info, LiteLLM.PrismaJson)
    field(:created_by, :string)
    field(:updated_by, :string)
    timestamps(inserted_at_source: :created_at)
  end
end