class UserContact < ApplicationRecord
  self.table_name = "UserContact"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :userId, uniqueness: true

  belongs_to :user, foreign_key: "userId", inverse_of: :contact
end
