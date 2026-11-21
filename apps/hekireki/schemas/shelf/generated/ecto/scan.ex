defmodule Shelf.Scan do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          latitude: String.t() | nil,
          longitude: String.t() | nil,
          user_agent: String.t() | nil,
          raw_qr_id: String.t() | nil,
          code_type: atom(),
          code: String.t() | nil,
          source: atom() | nil,
          manually_generated: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil,
          qr: Shelf.Qr.t() | nil,
          barcode: Shelf.Barcode.t() | nil,
          asset: Shelf.Asset.t() | nil,
          kit: Shelf.Kit.t() | nil,
          organization: Shelf.Organization.t() | nil
        }

  schema "Scan" do
    field(:id, :string, primary_key: true)
    field(:latitude, :string)
    field(:longitude, :string)
    field(:user_agent, :string, source: :userAgent)
    field(:raw_qr_id, :string, source: :rawQrId)
    field(:code_type, Ecto.Enum, values: [:QR, :BARCODE, :SAM_ID], default: :QR, source: :codeType)
    field(:code, :string)
    field(:source, Ecto.Enum, values: [:QR_LINK, :WEB_SCANNER, :WEB_DRAWER, :AUDIT, :COMPANION, :GPS_UPDATE])
    field(:manually_generated, :boolean, default: false, source: :manuallyGenerated)
    field(:user_id, :string, source: :userId)
    field(:qr_id, :string, source: :qrId)
    field(:barcode_id, :string, source: :barcodeId)
    field(:asset_id, :string, source: :assetId)
    field(:kit_id, :string, source: :kitId)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:qr, Shelf.Qr, foreign_key: :qr_id, define_field: false, type: :string)
    belongs_to(:barcode, Shelf.Barcode, foreign_key: :barcode_id, define_field: false, type: :string)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:kit, Shelf.Kit, foreign_key: :kit_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end