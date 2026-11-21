class PrintBatch < ApplicationRecord
  self.table_name = "PrintBatch"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :printed, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :name, presence: true, uniqueness: true

  has_many :qr_codes, class_name: "Qr", foreign_key: "batchId", inverse_of: :batch, dependent: :nullify
end
