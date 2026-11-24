defmodule Postiz.ItemUser do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          key: String.t(),
          user: Postiz.User.t() | nil
        }

  schema "ItemUser" do
    field(:key, :string)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, Postiz.User, foreign_key: :user_id, define_field: false)
  end
end