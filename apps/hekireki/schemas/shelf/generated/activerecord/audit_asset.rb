# expectedQuantity:: Units the audit expects to find, snapshotted when the row is created: the
#   units in the audit's scope (placed at its locations, held in its kits, or
#   held by its person; at most `Asset.quantity`), or `Asset.quantity` when
#   none are in scope or the audit has no scope. Null for INDIVIDUAL assets,
#   which count as one.
#
#   A snapshot, not a live read: stock that moves during the audit must not
#   move the target the auditor is counting against.
# foundQuantity:: Units the auditor counted. Null until a quantity-tracked asset is
#   counted, and always null for INDIVIDUAL assets. A FOUND row with a count
#   below `expectedQuantity` is a shortfall: the status stays FOUND (the pool
#   was located), the gap is this number.
#
#   Both columns are non-negative, enforced by CHECK constraints in raw SQL
#   (migration `20261104120000_add_audit_asset_quantities`).
class AuditAsset < ApplicationRecord
  self.table_name = "AuditAsset"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :expected, default: true
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :status, { pending: "PENDING", found: "FOUND", missing: "MISSING", unexpected: "UNEXPECTED" }, default: :pending,
                                                                                                      validate: true

  validates :auditSessionId, uniqueness: { scope: :assetId }

  belongs_to :audit_session, foreign_key: "auditSessionId", inverse_of: :assets
  belongs_to :asset, foreign_key: "assetId", inverse_of: :audit_entries
  belongs_to :scanned_by, class_name: "User", foreign_key: "scannedById", optional: true,
                          inverse_of: :audit_assets_scanned
  has_many :scans, class_name: "AuditScan", foreign_key: "auditAssetId", inverse_of: :audit_asset, dependent: :nullify
  has_many :images, class_name: "AuditImage", foreign_key: "auditAssetId", inverse_of: :audit_asset, dependent: :nullify
  has_many :notes, class_name: "AuditNote", foreign_key: "auditAssetId", inverse_of: :audit_asset, dependent: :nullify
end
