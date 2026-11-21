defmodule Shelf.Qr do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          version: integer(),
          error_correction: atom(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          asset: Shelf.Asset.t() | nil,
          kit: Shelf.Kit.t() | nil,
          user: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil,
          batch: Shelf.PrintBatch.t() | nil,
          scans: [Shelf.Scan.t()]
        }

  schema "Qr" do
    field(:id, :string, primary_key: true)
    field(:version, :integer, default: 0)
    field(:error_correction, Ecto.Enum, values: [:L, :M, :Q, :H], default: :L, source: :errorCorrection)
    field(:asset_id, :string, source: :assetId)
    field(:kit_id, :string, source: :kitId)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    field(:batch_id, :string, source: :batchId)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:kit, Shelf.Kit, foreign_key: :kit_id, define_field: false, type: :string)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:batch, Shelf.PrintBatch, foreign_key: :batch_id, define_field: false, type: :string)
    has_many(:scans, Shelf.Scan, foreign_key: :qr_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end