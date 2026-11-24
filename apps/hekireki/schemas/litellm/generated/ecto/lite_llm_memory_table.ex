defmodule LiteLLM.LiteLLMMemoryTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :memory_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          key: String.t(),
          value: String.t(),
          metadata: map() | list() | String.t() | number() | boolean() | nil,
          user_id: String.t() | nil,
          team_id: String.t() | nil,
          created_by: String.t() | nil,
          updated_by: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_MemoryTable" do
    field(:key, :string)
    field(:value, :string)
    field(:metadata, LiteLLM.PrismaJson)
    field(:user_id, :string)
    field(:team_id, :string)
    field(:created_by, :string)
    field(:updated_by, :string)
    timestamps(inserted_at_source: :created_at)
  end
end