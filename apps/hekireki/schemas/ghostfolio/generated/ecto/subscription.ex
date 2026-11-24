defmodule Ghostfolio.Subscription do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Ghostfolio.PrismaDateTime, autogenerate: {Ghostfolio.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          expires_at: DateTime.t(),
          price: float() | nil,
          stripe_checkout_session_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Ghostfolio.User.t() | nil
        }

  schema "Subscription" do
    field(:expires_at, Ghostfolio.PrismaDateTime, source: :expiresAt)
    field(:price, :float)
    field(:stripe_checkout_session_id, :string, source: :stripeCheckoutSessionId)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, Ghostfolio.User, foreign_key: :user_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end