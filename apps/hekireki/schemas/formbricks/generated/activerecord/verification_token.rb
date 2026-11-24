# Stores one-time verification tokens used by Auth.js adapter flows.
class VerificationToken < ApplicationRecord
  self.table_name = "VerificationToken"
  self.primary_key = %w[identifier token]

  validates :identifier, presence: true, uniqueness: { scope: :token }
  validates :token, presence: true
  validates :expires, presence: true
end
