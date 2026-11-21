defmodule Shelf.Role do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: atom(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          users: [Shelf.User.t()]
        }

  schema "Role" do
    field(:id, :string, primary_key: true)
    field(:name, Ecto.Enum, values: [:USER, :ADMIN], default: :USER)
    many_to_many(:users, Shelf.User, join_through: "_RoleToUser", join_keys: [A: :id, B: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end