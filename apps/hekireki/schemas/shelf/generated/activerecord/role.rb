class Role < ApplicationRecord
  self.table_name = "Role"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :name, { user: "USER", admin: "ADMIN" }, default: :user, validate: true

  has_and_belongs_to_many :users, join_table: "_RoleToUser", foreign_key: "A", association_foreign_key: "B"
end
