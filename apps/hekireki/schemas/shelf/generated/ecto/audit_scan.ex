defmodule Shelf.AuditScan do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          code: String.t() | nil,
          asset_title: String.t() | nil,
          was_expected: boolean() | nil,
          metadata: map() | list() | String.t() | number() | boolean() | nil,
          scanned_at: DateTime.t(),
          inserted_at: DateTime.t(),
          audit_session: Shelf.AuditSession.t() | nil,
          audit_asset: Shelf.AuditAsset.t() | nil,
          asset: Shelf.Asset.t() | nil,
          scanned_by: Shelf.User.t() | nil
        }

  schema "AuditScan" do
    field(:id, :string, primary_key: true)
    field(:code, :string)
    # What was scanned, recorded BY VALUE at scan time.
    #
    # `asset` and `auditAsset` are both SetNull/Cascade, so deleting an asset
    # leaves this row pointing at nothing: no name, and no way to tell whether
    # it had been expected. An audit is a historical record, and a record whose
    # entries can silently lose their meaning is not one — so the two
    # descriptive facts are snapshotted here rather than only referenced.
    #
    # Readers prefer the LIVE asset when it still exists (a rename should show
    # the current name) and fall back to these.
    #
    # Null means the snapshot could not be taken: scans whose asset had already
    # been deleted before these columns were backfilled. Nothing can recover
    # those, and they degrade to the scanned `code`, which also survives. Every
    # other row carries a snapshot, so null is never merely "old".
    field(:asset_title, :string, source: :assetTitle)
    field(:was_expected, :boolean, source: :wasExpected)
    field(:metadata, Shelf.PrismaJson)
    field(:scanned_at, Shelf.PrismaDateTime, autogenerate: true, source: :scannedAt)
    field(:audit_session_id, :string, source: :auditSessionId)
    field(:audit_asset_id, :string, source: :auditAssetId)
    field(:asset_id, :string, source: :assetId)
    field(:scanned_by_id, :string, source: :scannedById)
    belongs_to(:audit_session, Shelf.AuditSession, foreign_key: :audit_session_id, define_field: false, type: :string)
    belongs_to(:audit_asset, Shelf.AuditAsset, foreign_key: :audit_asset_id, define_field: false, type: :string)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:scanned_by, Shelf.User, foreign_key: :scanned_by_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end