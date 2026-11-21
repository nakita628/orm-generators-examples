# Held by the one seat-taking write running on an organization (seat-requests.feature). A
# row past expiresAt is free to take; a holder deletes its own row when its write is done.
#
# holder:: Random per hold, so a release deletes only its own lease.
class OrganizationSeatLease < ApplicationRecord
  self.table_name = "OrganizationSeatLease"
  self.primary_key = "organizationId"

  alias_attribute :created_at, :createdAt

  validates :holder, presence: true
  validates :expiresAt, presence: true
end
