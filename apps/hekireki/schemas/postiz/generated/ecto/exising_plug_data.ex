defmodule Postiz.ExisingPlugData do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          method_name: String.t(),
          value: String.t(),
          integration: Postiz.Integration.t() | nil
        }

  schema "ExisingPlugData" do
    field(:method_name, :string, source: :methodName)
    field(:value, :string)
    field(:integration_id, :string, source: :integrationId)
    belongs_to(:integration, Postiz.Integration, foreign_key: :integration_id, define_field: false)
  end
end