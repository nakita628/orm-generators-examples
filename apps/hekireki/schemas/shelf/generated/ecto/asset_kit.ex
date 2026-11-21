defmodule Shelf.AssetKit do
  @moduledoc """
  Pivot linking an `Asset` to the `Kit` it belongs to.

  `@@unique([assetId])` enforces "at most one kit per asset" today.
  The structure already allows multi-row-per-asset; if multi-kit
  allocation lands later, drop the `assetId` unique constraint and
  add a type-aware single-row trigger + sum-within-total CONSTRAINT
  TRIGGER (parallel to the Custody / BookingAsset pivots).
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          quantity: integer(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          asset: Shelf.Asset.t() | nil,
          kit: Shelf.Kit.t() | nil,
          organization: Shelf.Organization.t() | nil,
          asset_locations: [Shelf.AssetLocation.t()]
        }

  schema "AssetKit" do
    field(:id, :string, primary_key: true)
    field(:quantity, :integer, default: 1)
    field(:asset_id, :string, source: :assetId)
    field(:kit_id, :string, source: :kitId)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:kit, Shelf.Kit, foreign_key: :kit_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    has_many(:asset_locations, Shelf.AssetLocation, foreign_key: :asset_kit_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end