defmodule Shelf.Tag do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          color: String.t() | nil,
          use_for: atom(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil,
          assets: [Shelf.Asset.t()],
          bookings: [Shelf.Booking.t()]
        }

  schema "Tag" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    field(:color, :string)
    field(:use_for, {:array, Ecto.Enum}, values: [:ASSET, :BOOKING], source: :useFor)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    many_to_many(:assets, Shelf.Asset, join_through: "_AssetToTag", join_keys: [B: :id, A: :id])
    many_to_many(:bookings, Shelf.Booking, join_through: "_BookingToTag", join_keys: [B: :id, A: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end