defmodule LiteLLM.LiteLLMManagedFileContentTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          content: binary(),
          inserted_at: DateTime.t()
        }

  schema "LiteLLM_ManagedFileContentTable" do
    field(:content, :binary)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end