class ScimToken < ApplicationRecord
  self.table_name = "ScimToken"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt

  validates :tokenHash, presence: true, uniqueness: true
  validates :label, presence: true
  validates :createdById, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :scim_tokens
end
