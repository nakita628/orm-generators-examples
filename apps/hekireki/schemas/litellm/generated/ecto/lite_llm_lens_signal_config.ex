defmodule LiteLLM.LiteLLMLensSignalConfig do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          data: map() | list() | String.t() | number() | boolean()
        }

  schema "LiteLLM_LensSignalConfig" do
    field(:id, :string, primary_key: true)
    field(:data, LiteLLM.PrismaJson)
  end
end