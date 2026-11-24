class Waitlist < ApplicationRecord
  self.table_name = "waitlist"
  self.primary_key = "email"

  validates :email, presence: true, uniqueness: true
end
