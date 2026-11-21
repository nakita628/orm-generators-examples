# name:: User-provided name for the token
# type:: Used to differentiate between user-generated and system-generated tokens
# hashedToken:: This is used to find the token in the database
# expiresAt:: Optional expiration date for the token
class OrganizationAccessToken < ApplicationRecord
  self.table_name = "OrganizationAccessToken"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { user: "USER", system: "SYSTEM" }, default: :user, prefix: true, validate: true

  validates :name, presence: true
  validates :hashedToken, presence: true, uniqueness: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :organization_access_tokens
end
