defmodule Shelf.Asset do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          title: String.t(),
          description: String.t() | nil,
          main_image: String.t() | nil,
          thumbnail_image: String.t() | nil,
          main_image_expiration: DateTime.t() | nil,
          status: atom(),
          valuation: float() | nil,
          available_to_book: boolean(),
          sequential_id: String.t() | nil,
          archived_at: DateTime.t() | nil,
          type: atom(),
          quantity: integer() | nil,
          min_quantity: integer() | nil,
          consumption_type: atom() | nil,
          unit_of_measure: String.t() | nil,
          low_stock_notified_at: DateTime.t() | nil,
          stock_ledger_started_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          asset_model: Shelf.AssetModel.t() | nil,
          user: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil,
          category: Shelf.Category.t() | nil,
          preferred_barcode: Shelf.Barcode.t() | nil,
          custody: [Shelf.Custody.t()],
          notes: [Shelf.Note.t()],
          qr_codes: [Shelf.Qr.t()],
          barcodes: [Shelf.Barcode.t()],
          reports: [Shelf.ReportFound.t()],
          custom_fields: [Shelf.AssetCustomFieldValue.t()],
          booking_assets: [Shelf.BookingAsset.t()],
          consumption_logs: [Shelf.ConsumptionLog.t()],
          reminders: [Shelf.AssetReminder.t()],
          audit_entries: [Shelf.AuditAsset.t()],
          audit_scans: [Shelf.AuditScan.t()],
          scans: [Shelf.Scan.t()],
          asset_kits: [Shelf.AssetKit.t()],
          asset_locations: [Shelf.AssetLocation.t()],
          tags: [Shelf.Tag.t()]
        }

  schema "Asset" do
    field(:id, :string, primary_key: true)
    field(:title, :string)
    field(:description, :string)
    field(:main_image, :string, source: :mainImage)
    field(:thumbnail_image, :string, source: :thumbnailImage)
    field(:main_image_expiration, Shelf.PrismaDateTime, source: :mainImageExpiration)
    field(:status, Ecto.Enum, values: [:AVAILABLE, :IN_CUSTODY, :CHECKED_OUT], default: :AVAILABLE)
    field(:valuation, :float, source: :value)
    field(:available_to_book, :boolean, default: true, source: :availableToBook)
    field(:sequential_id, :string, source: :sequentialId)
    field(:archived_at, Shelf.PrismaDateTime, source: :archivedAt)
    field(:type, Ecto.Enum, values: [:INDIVIDUAL, :QUANTITY_TRACKED], default: :INDIVIDUAL)
    field(:quantity, :integer)
    field(:min_quantity, :integer, source: :minQuantity)
    field(:consumption_type, Ecto.Enum, values: [:ONE_WAY, :TWO_WAY], source: :consumptionType)
    field(:unit_of_measure, :string, source: :unitOfMeasure)
    field(:low_stock_notified_at, Shelf.PrismaDateTime, source: :lowStockNotifiedAt)
    field(:stock_ledger_started_at, Shelf.PrismaDateTime, source: :stockLedgerStartedAt)
    field(:asset_model_id, :string, source: :assetModelId)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    field(:category_id, :string, source: :categoryId)
    field(:preferred_barcode_id, :string, source: :preferredBarcodeId)
    belongs_to(:asset_model, Shelf.AssetModel, foreign_key: :asset_model_id, define_field: false, type: :string)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:category, Shelf.Category, foreign_key: :category_id, define_field: false, type: :string)
    belongs_to(:preferred_barcode, Shelf.Barcode, foreign_key: :preferred_barcode_id, define_field: false, type: :string)
    has_many(:custody, Shelf.Custody, foreign_key: :asset_id, references: :id)
    has_many(:notes, Shelf.Note, foreign_key: :asset_id, references: :id)
    has_many(:qr_codes, Shelf.Qr, foreign_key: :asset_id, references: :id)
    has_many(:barcodes, Shelf.Barcode, foreign_key: :asset_id, references: :id)
    has_many(:reports, Shelf.ReportFound, foreign_key: :asset_id, references: :id)
    has_many(:custom_fields, Shelf.AssetCustomFieldValue, foreign_key: :asset_id, references: :id)
    has_many(:booking_assets, Shelf.BookingAsset, foreign_key: :asset_id, references: :id)
    has_many(:consumption_logs, Shelf.ConsumptionLog, foreign_key: :asset_id, references: :id)
    has_many(:reminders, Shelf.AssetReminder, foreign_key: :asset_id, references: :id)
    has_many(:audit_entries, Shelf.AuditAsset, foreign_key: :asset_id, references: :id)
    has_many(:audit_scans, Shelf.AuditScan, foreign_key: :asset_id, references: :id)
    has_many(:scans, Shelf.Scan, foreign_key: :asset_id, references: :id)
    has_many(:asset_kits, Shelf.AssetKit, foreign_key: :asset_id, references: :id)
    has_many(:asset_locations, Shelf.AssetLocation, foreign_key: :asset_id, references: :id)
    many_to_many(:tags, Shelf.Tag, join_through: "_AssetToTag", join_keys: [A: :id, B: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end