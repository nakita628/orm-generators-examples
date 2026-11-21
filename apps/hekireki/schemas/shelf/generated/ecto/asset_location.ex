defmodule Shelf.AssetLocation do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          quantity: integer(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          asset: Shelf.Asset.t() | nil,
          location: Shelf.Location.t() | nil,
          organization: Shelf.Organization.t() | nil,
          asset_kit: Shelf.AssetKit.t() | nil
        }

  schema "AssetLocation" do
    field(:id, :string, primary_key: true)
    field(:quantity, :integer, default: 1)
    field(:asset_id, :string, source: :assetId)
    field(:location_id, :string, source: :locationId)
    field(:organization_id, :string, source: :organizationId)
    field(:asset_kit_id, :string, source: :assetKitId)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:location, Shelf.Location, foreign_key: :location_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:asset_kit, Shelf.AssetKit, foreign_key: :asset_kit_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end