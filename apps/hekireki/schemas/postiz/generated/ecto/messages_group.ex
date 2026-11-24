defmodule Postiz.MessagesGroup do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Postiz.PrismaDateTime, autogenerate: {Postiz.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          buyer: Postiz.User.t() | nil,
          buyer_organization: Postiz.Organization.t() | nil,
          seller: Postiz.User.t() | nil,
          messages: [Postiz.Messages.t()],
          orders: [Postiz.Orders.t()]
        }

  schema "MessagesGroup" do
    field(:buyer_id, :string, source: :buyerId)
    field(:buyer_organization_id, :string, source: :buyerOrganizationId)
    field(:seller_id, :string, source: :sellerId)
    belongs_to(:buyer, Postiz.User, foreign_key: :buyer_id, define_field: false)
    belongs_to(:buyer_organization, Postiz.Organization, foreign_key: :buyer_organization_id, define_field: false)
    belongs_to(:seller, Postiz.User, foreign_key: :seller_id, define_field: false)
    has_many(:messages, Postiz.Messages, foreign_key: :group_id)
    has_many(:orders, Postiz.Orders, foreign_key: :message_group_id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end