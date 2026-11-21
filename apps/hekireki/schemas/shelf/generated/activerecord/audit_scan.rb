# assetTitle:: What was scanned, recorded BY VALUE at scan time.
#
#   `asset` and `auditAsset` are both SetNull/Cascade, so deleting an asset
#   leaves this row pointing at nothing: no name, and no way to tell whether
#   it had been expected. An audit is a historical record, and a record whose
#   entries can silently lose their meaning is not one — so the two
#   descriptive facts are snapshotted here rather than only referenced.
#
#   Readers prefer the LIVE asset when it still exists (a rename should show
#   the current name) and fall back to these.
#
#   Null means the snapshot could not be taken: scans whose asset had already
#   been deleted before these columns were backfilled. Nothing can recover
#   those, and they degrade to the scanned `code`, which also survives. Every
#   other row carries a snapshot, so null is never merely "old".
class AuditScan < ApplicationRecord
  self.table_name = "AuditScan"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt

  belongs_to :audit_session, foreign_key: "auditSessionId", inverse_of: :scans
  belongs_to :audit_asset, foreign_key: "auditAssetId", optional: true, inverse_of: :scans
  belongs_to :asset, foreign_key: "assetId", optional: true, inverse_of: :audit_scans
  belongs_to :scanned_by, class_name: "User", foreign_key: "scannedById", optional: true, inverse_of: :audit_scans
end
