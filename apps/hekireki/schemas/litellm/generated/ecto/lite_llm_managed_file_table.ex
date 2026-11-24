defmodule LiteLLM.LiteLLMManagedFileTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          unified_file_id: String.t(),
          file_object: map() | list() | String.t() | number() | boolean() | nil,
          model_mappings: map() | list() | String.t() | number() | boolean(),
          flat_model_file_ids: [String.t()],
          storage_backend: String.t() | nil,
          storage_url: String.t() | nil,
          created_by: String.t() | nil,
          team_id: String.t() | nil,
          updated_by: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_ManagedFileTable" do
    field(:unified_file_id, :string)
    field(:file_object, LiteLLM.PrismaJson)
    field(:model_mappings, LiteLLM.PrismaJson)
    field(:flat_model_file_ids, {:array, :string})
    field(:storage_backend, :string)
    field(:storage_url, :string)
    field(:created_by, :string)
    field(:team_id, :string)
    field(:updated_by, :string)
    timestamps(inserted_at_source: :created_at)
  end
end