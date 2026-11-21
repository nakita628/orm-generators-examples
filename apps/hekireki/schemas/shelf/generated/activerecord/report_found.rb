class ReportFound < ApplicationRecord
  self.table_name = "ReportFound"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :email, presence: true
  validates :content, presence: true

  belongs_to :asset, foreign_key: "assetId", optional: true, inverse_of: :reports
  belongs_to :kit, foreign_key: "kitId", optional: true, inverse_of: :reports
end
