# One calendar month's (UTC) seat peaks for an organization (ADR-002).
#
# month:: `YYYY-MM`.
class OrganizationSeatPeakMonth < ApplicationRecord
  self.table_name = "OrganizationSeatPeakMonth"
  self.primary_key = %w[organizationId month]

  validates :fullPeak, presence: true
  validates :fullPeakAt, presence: true
  validates :litePeak, presence: true
  validates :litePeakAt, presence: true
  validates :fullAtEnd, presence: true
  validates :liteAtEnd, presence: true
end
