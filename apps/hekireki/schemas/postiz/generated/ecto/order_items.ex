defmodule Postiz.OrderItems do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          quantity: integer(),
          price: integer(),
          integration: Postiz.Integration.t() | nil,
          order: Postiz.Orders.t() | nil
        }

  schema "OrderItems" do
    field(:quantity, :integer)
    field(:price, :integer)
    field(:integration_id, :string, source: :integrationId)
    field(:order_id, :string, source: :orderId)
    belongs_to(:integration, Postiz.Integration, foreign_key: :integration_id, define_field: false)
    belongs_to(:order, Postiz.Orders, foreign_key: :order_id, define_field: false)
  end
end