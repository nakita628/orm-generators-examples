defmodule TypeHero.Role do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          role: atom(),
          user: [TypeHero.User.t()]
        }

  schema "Role" do
    field(:role, Ecto.Enum, values: [:USER, :ADMIN, :MODERATOR, :CREATOR, :SUPPORTER, :CONTRIBUTOR])
    many_to_many(:user, TypeHero.User, join_through: "_RoleToUser", join_keys: [A: :id, B: :id])
  end
end