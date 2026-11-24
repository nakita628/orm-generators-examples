defmodule Postiz.PayoutProblems do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Postiz.PrismaDateTime, autogenerate: {Postiz.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          status: String.t(),
          amount: integer(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          order: Postiz.Orders.t() | nil,
          post: Postiz.Post.t() | nil,
          user: Postiz.User.t() | nil
        }

  schema "PayoutProblems" do
    field(:status, :string)
    field(:amount, :integer)
    field(:order_id, :string, source: :orderId)
    field(:post_id, :string, source: :postId)
    field(:user_id, :string, source: :userId)
    belongs_to(:order, Postiz.Orders, foreign_key: :order_id, define_field: false)
    belongs_to(:post, Postiz.Post, foreign_key: :post_id, define_field: false)
    belongs_to(:user, Postiz.User, foreign_key: :user_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end