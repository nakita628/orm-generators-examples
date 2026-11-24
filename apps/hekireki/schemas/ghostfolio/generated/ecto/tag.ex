defmodule Ghostfolio.Tag do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          user: Ghostfolio.User.t() | nil,
          accounts: [Ghostfolio.TagsOnAccounts.t()],
          activities: [Ghostfolio.Order.t()]
        }

  schema "Tag" do
    field(:name, :string)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, Ghostfolio.User, foreign_key: :user_id, define_field: false)
    has_many(:accounts, Ghostfolio.TagsOnAccounts, foreign_key: :tag_id)
    many_to_many(:activities, Ghostfolio.Order, join_through: "_OrderToTag", join_keys: [B: :id, A: :id])
  end
end