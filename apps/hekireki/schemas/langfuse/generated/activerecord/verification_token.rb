class VerificationToken < ApplicationRecord
  self.primary_key = "token"

  validates :identifier, presence: true, uniqueness: { scope: :token }
  validates :token, presence: true, uniqueness: true
  validates :expires, presence: true
end
