defmodule Shelf.Kit do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          status: atom(),
          image: String.t() | nil,
          image_expiration: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Shelf.Organization.t() | nil,
          created_by: Shelf.User.t() | nil,
          category: Shelf.Category.t() | nil,
          location: Shelf.Location.t() | nil,
          custody: Shelf.KitCustody.t() | nil,
          asset_kits: [Shelf.AssetKit.t()],
          qr_codes: [Shelf.Qr.t()],
          barcodes: [Shelf.Barcode.t()],
          reports: [Shelf.ReportFound.t()],
          scans: [Shelf.Scan.t()]
        }

  schema "Kit" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    field(:status, Ecto.Enum, values: [:AVAILABLE, :IN_CUSTODY, :CHECKED_OUT], default: :AVAILABLE)
    field(:image, :string)
    field(:image_expiration, Shelf.PrismaDateTime, source: :imageExpiration)
    field(:organization_id, :string, source: :organizationId)
    field(:created_by_id, :string, source: :createdById)
    field(:category_id, :string, source: :categoryId)
    field(:location_id, :string, source: :locationId)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:created_by, Shelf.User, foreign_key: :created_by_id, define_field: false, type: :string)
    belongs_to(:category, Shelf.Category, foreign_key: :category_id, define_field: false, type: :string)
    belongs_to(:location, Shelf.Location, foreign_key: :location_id, define_field: false, type: :string)
    has_one(:custody, Shelf.KitCustody, foreign_key: :kit_id, references: :id)
    has_many(:asset_kits, Shelf.AssetKit, foreign_key: :kit_id, references: :id)
    has_many(:qr_codes, Shelf.Qr, foreign_key: :kit_id, references: :id)
    has_many(:barcodes, Shelf.Barcode, foreign_key: :kit_id, references: :id)
    has_many(:reports, Shelf.ReportFound, foreign_key: :kit_id, references: :id)
    has_many(:scans, Shelf.Scan, foreign_key: :kit_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end