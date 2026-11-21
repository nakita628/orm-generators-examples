defmodule Shelf.UserUpdateRead do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          read_at: DateTime.t(),
          user: Shelf.User.t() | nil,
          update: Shelf.Update.t() | nil
        }

  schema "UserUpdateRead" do
    field(:id, :string, primary_key: true)
    field(:read_at, Shelf.PrismaDateTime, autogenerate: true, source: :readAt)
    field(:user_id, :string, source: :userId)
    field(:update_id, :string, source: :updateId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:update, Shelf.Update, foreign_key: :update_id, define_field: false, type: :string)
  end
end