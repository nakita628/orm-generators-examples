class Scan < ApplicationRecord
  self.table_name = "Scan"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :manuallyGenerated, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :codeType, { qr: "QR", barcode: "BARCODE", sam_id: "SAM_ID" }, default: :qr, validate: true
  enum :source, { qr_link: "QR_LINK", web_scanner: "WEB_SCANNER", web_drawer: "WEB_DRAWER", audit: "AUDIT", companion: "COMPANION", gps_update: "GPS_UPDATE" },
       validate: { allow_nil: true }

  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :scans
  belongs_to :qr, foreign_key: "qrId", optional: true, inverse_of: :scans
  belongs_to :barcode, foreign_key: "barcodeId", optional: true, inverse_of: :scans
  belongs_to :asset, foreign_key: "assetId", optional: true, inverse_of: :scans
  belongs_to :kit, foreign_key: "kitId", optional: true, inverse_of: :scans
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :scans
end
