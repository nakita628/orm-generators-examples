class Barcode < ApplicationRecord
  self.table_name = "Barcode"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { code128: "Code128", code39: "Code39", data_matrix: "DataMatrix", external_qr: "ExternalQR", ean13: "EAN13" },
       default: :code128, validate: true

  validates :value, presence: true, uniqueness: { scope: :organizationId }

  belongs_to :asset, foreign_key: "assetId", optional: true, inverse_of: :barcodes
  belongs_to :kit, foreign_key: "kitId", optional: true, inverse_of: :barcodes
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :barcodes
  has_one :preferred_by_asset, class_name: "Asset", foreign_key: "preferredBarcodeId", inverse_of: :preferred_barcode,
                               dependent: :nullify
  has_many :scans, foreign_key: "barcodeId", inverse_of: :barcode, dependent: :nullify
end
