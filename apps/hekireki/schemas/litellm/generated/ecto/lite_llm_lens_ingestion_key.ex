defmodule LiteLLM.LiteLLMLensIngestionKey do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          data: map() | list() | String.t() | number() | boolean()
        }

  schema "LiteLLM_LensIngestionKey" do
    field(:id, :string, primary_key: true)
    field(:data, LiteLLM.PrismaJson)
  end
end