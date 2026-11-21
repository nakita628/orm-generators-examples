# One CSV content import. Holds the parsed rows from preview until the import
# completes, so the job imports exactly the file the user previewed.
# A partial unique index on (organizationId) WHERE status IN (QUEUED, RUNNING) allows one
# active import per workspace. Prisma cannot express it; it lives in migration
# 20261015120100_asset_import_one_active_per_workspace.
#
# processedRows:: Rows committed so far. Advanced in the same transaction as each chunk.
# rows:: Parsed rows (`StoredImportRow[]`). Cleared when the import completes.
# resolved:: Name to id maps written by the PREPARE step (`ResolvedImportEntities`).
# errorMessage:: User-facing reason the import failed. Null unless status is FAILED.
# warnings:: Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500.
# warningCount:: Exact number of warnings, including any beyond the stored cap.
# lockedUntil:: A step holds the import until this time. Null when no step is running.
class AssetImport < ApplicationRecord
  self.table_name = "AssetImport"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :processedRows, default: 0
  attribute :canUseBarcodes, default: false
  attribute :warningCount, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :status, { draft: "DRAFT", queued: "QUEUED", running: "RUNNING", completed: "COMPLETED", failed: "FAILED" },
       default: :draft, validate: true
  enum :phase, { prepare: "PREPARE", rows: "ROWS", finalize: "FINALIZE" }, default: :prepare, validate: true

  validates :totalRows, presence: true

  belongs_to :created_by, class_name: "User", foreign_key: "createdById", inverse_of: :asset_imports
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :asset_imports
end
