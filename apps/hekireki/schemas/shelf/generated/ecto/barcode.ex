defmodule Shelf.Barcode do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          value: String.t(),
          type: atom(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          asset: Shelf.Asset.t() | nil,
          kit: Shelf.Kit.t() | nil,
          organization: Shelf.Organization.t() | nil,
          preferred_by_asset: Shelf.Asset.t() | nil,
          scans: [Shelf.Scan.t()]
        }

  schema "Barcode" do
    field(:id, :string, primary_key: true)
    field(:value, :string)
    field(:type, Ecto.Enum, values: [:Code128, :Code39, :DataMatrix, :ExternalQR, :EAN13], default: :Code128)
    field(:asset_id, :string, source: :assetId)
    field(:kit_id, :string, source: :kitId)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:kit, Shelf.Kit, foreign_key: :kit_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    has_one(:preferred_by_asset, Shelf.Asset, foreign_key: :preferred_barcode_id, references: :id)
    has_many(:scans, Shelf.Scan, foreign_key: :barcode_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end