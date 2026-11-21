class Note < ApplicationRecord
  self.table_name = "Note"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { comment: "COMMENT", update: "UPDATE" }, default: :comment, prefix: true, validate: true

  validates :content, presence: true

  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :notes
  belongs_to :asset, foreign_key: "assetId", inverse_of: :notes
end
