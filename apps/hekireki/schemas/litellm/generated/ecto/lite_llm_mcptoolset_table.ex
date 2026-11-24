defmodule LiteLLM.LiteLLMMCPToolsetTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :toolset_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          toolset_name: String.t(),
          description: String.t() | nil,
          tools: map() | list() | String.t() | number() | boolean(),
          created_by: String.t() | nil,
          updated_by: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_MCPToolsetTable" do
    field(:toolset_name, :string)
    field(:description, :string)
    field(:tools, LiteLLM.PrismaJson, default: [])
    field(:created_by, :string)
    field(:updated_by, :string)
    timestamps(inserted_at_source: :created_at)
  end
end