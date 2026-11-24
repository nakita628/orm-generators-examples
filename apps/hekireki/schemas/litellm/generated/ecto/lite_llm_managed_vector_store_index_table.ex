defmodule LiteLLM.LiteLLMManagedVectorStoreIndexTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          index_name: String.t(),
          litellm_params: map() | list() | String.t() | number() | boolean(),
          index_info: map() | list() | String.t() | number() | boolean() | nil,
          created_by: String.t() | nil,
          updated_by: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_ManagedVectorStoreIndexTable" do
    field(:index_name, :string)
    field(:litellm_params, LiteLLM.PrismaJson)
    field(:index_info, LiteLLM.PrismaJson)
    field(:created_by, :string)
    field(:updated_by, :string)
    timestamps(inserted_at_source: :created_at)
  end
end