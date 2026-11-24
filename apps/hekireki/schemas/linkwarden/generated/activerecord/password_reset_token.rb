class PasswordResetToken < ApplicationRecord
  self.table_name = "PasswordResetToken"
  self.primary_key = "token"

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :identifier, presence: true
  validates :token, presence: true, uniqueness: true
  validates :expires, presence: true
end
