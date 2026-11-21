class Qr < ApplicationRecord
  self.table_name = "Qr"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :version, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :errorCorrection, { l: "L", m: "M", q: "Q", h: "H" }, default: :l, validate: true

  belongs_to :asset, foreign_key: "assetId", optional: true, inverse_of: :qr_codes
  belongs_to :kit, foreign_key: "kitId", optional: true, inverse_of: :qr_codes
  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :qr_codes
  belongs_to :organization, foreign_key: "organizationId", optional: true, inverse_of: :qr_codes
  belongs_to :batch, class_name: "PrintBatch", foreign_key: "batchId", optional: true, inverse_of: :qr_codes
  has_many :scans, foreign_key: "qrId", inverse_of: :qr, dependent: :nullify
end
