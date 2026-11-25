# An organization enrolled in the weekly digest. No row: not enrolled, so nothing sends.
class DigestCohortOrganization < ApplicationRecord
  self.table_name = "DigestCohortOrganization"
  self.primary_key = "organizationId"

  validates :enrolledBy, presence: true
end
