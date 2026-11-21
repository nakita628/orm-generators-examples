defmodule Shelf.AuditAsset do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          expected: boolean(),
          status: atom(),
          scanned_at: DateTime.t() | nil,
          metadata: map() | nil,
          expected_quantity: integer() | nil,
          found_quantity: integer() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          audit_session: Shelf.AuditSession.t() | nil,
          asset: Shelf.Asset.t() | nil,
          scanned_by: Shelf.User.t() | nil,
          scans: [Shelf.AuditScan.t()],
          images: [Shelf.AuditImage.t()],
          notes: [Shelf.AuditNote.t()]
        }

  schema "AuditAsset" do
    field(:id, :string, primary_key: true)
    field(:expected, :boolean, default: true)
    field(:status, Ecto.Enum, values: [:PENDING, :FOUND, :MISSING, :UNEXPECTED], default: :PENDING)
    field(:scanned_at, Shelf.PrismaDateTime, source: :scannedAt)
    field(:metadata, :map)
    # Units the audit expects to find, snapshotted when the row is created: the
    # units in the audit's scope (placed at its locations, held in its kits, or
    # held by its person; at most `Asset.quantity`), or `Asset.quantity` when
    # none are in scope or the audit has no scope. Null for INDIVIDUAL assets,
    # which count as one.
    #
    # A snapshot, not a live read: stock that moves during the audit must not
    # move the target the auditor is counting against.
    field(:expected_quantity, :integer, source: :expectedQuantity)
    # Units the auditor counted. Null until a quantity-tracked asset is
    # counted, and always null for INDIVIDUAL assets. A FOUND row with a count
    # below `expectedQuantity` is a shortfall: the status stays FOUND (the pool
    # was located), the gap is this number.
    #
    # Both columns are non-negative, enforced by CHECK constraints in raw SQL
    # (migration `20261104120000_add_audit_asset_quantities`).
    field(:found_quantity, :integer, source: :foundQuantity)
    field(:audit_session_id, :string, source: :auditSessionId)
    field(:asset_id, :string, source: :assetId)
    field(:scanned_by_id, :string, source: :scannedById)
    belongs_to(:audit_session, Shelf.AuditSession, foreign_key: :audit_session_id, define_field: false, type: :string)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:scanned_by, Shelf.User, foreign_key: :scanned_by_id, define_field: false, type: :string)
    has_many(:scans, Shelf.AuditScan, foreign_key: :audit_asset_id, references: :id)
    has_many(:images, Shelf.AuditImage, foreign_key: :audit_asset_id, references: :id)
    has_many(:notes, Shelf.AuditNote, foreign_key: :audit_asset_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end