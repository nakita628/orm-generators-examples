defmodule LiteLLM.LiteLLMLensDataset do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          revision: integer(),
          created_at: DateTime.t(),
          data: map() | list() | String.t() | number() | boolean()
        }

  schema "LiteLLM_LensDataset" do
    field(:id, :string, primary_key: true)
    field(:revision, :integer, primary_key: true)
    field(:created_at, LiteLLM.PrismaDateTime)
    field(:data, LiteLLM.PrismaJson)
  end
end