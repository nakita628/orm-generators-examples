defmodule Shelf.Location do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          address: String.t() | nil,
          latitude: float() | nil,
          longitude: float() | nil,
          image_url: String.t() | nil,
          thumbnail_url: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          image: Shelf.Image.t() | nil,
          user: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil,
          parent: Shelf.Location.t() | nil,
          asset_locations: [Shelf.AssetLocation.t()],
          kits: [Shelf.Kit.t()],
          notes: [Shelf.LocationNote.t()],
          custodies: [Shelf.Custody.t()],
          consumption_logs: [Shelf.ConsumptionLog.t()],
          children: [Shelf.Location.t()]
        }

  schema "Location" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    field(:address, :string)
    field(:latitude, :float)
    field(:longitude, :float)
    field(:image_url, :string, source: :imageUrl)
    field(:thumbnail_url, :string, source: :thumbnailUrl)
    field(:image_id, :string, source: :imageId)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    field(:parent_id, :string, source: :parentId)
    belongs_to(:image, Shelf.Image, foreign_key: :image_id, define_field: false, type: :string)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:parent, Shelf.Location, foreign_key: :parent_id, define_field: false, type: :string)
    has_many(:asset_locations, Shelf.AssetLocation, foreign_key: :location_id, references: :id)
    has_many(:kits, Shelf.Kit, foreign_key: :location_id, references: :id)
    has_many(:notes, Shelf.LocationNote, foreign_key: :location_id, references: :id)
    has_many(:custodies, Shelf.Custody, foreign_key: :location_id, references: :id)
    has_many(:consumption_logs, Shelf.ConsumptionLog, foreign_key: :location_id, references: :id)
    has_many(:children, Shelf.Location, foreign_key: :parent_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end