defmodule Shelf.LocationNote do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          content: String.t(),
          type: atom(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil,
          location: Shelf.Location.t() | nil
        }

  schema "LocationNote" do
    field(:id, :string, primary_key: true)
    field(:content, :string)
    field(:type, Ecto.Enum, values: [:COMMENT, :UPDATE], default: :COMMENT)
    field(:user_id, :string, source: :userId)
    field(:location_id, :string, source: :locationId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:location, Shelf.Location, foreign_key: :location_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end