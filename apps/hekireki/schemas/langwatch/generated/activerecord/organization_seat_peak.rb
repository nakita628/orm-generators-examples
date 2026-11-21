# Organization's monthly seat peak, folded from its seat changes (modules/organization, ADR-002).
# A read model, one row per organization: rebuildable by replaying the event log, so no relation.
# Epochs in milliseconds.
# Organization's seat-peak counts (ADR-002): one row per organization once counting opened.
#
# organizationId:: The organization, which is also the aggregate id.
# openedAt:: When the opening snapshot was taken, epoch milliseconds.
# latestChangeAt:: The latest change time applied to any holder; invitations expire up to it.
# mismatches:: Changes whose `from` disagreed with the seat the fold held.
class OrganizationSeatPeak < ApplicationRecord
  self.table_name = "OrganizationSeatPeak"
  self.primary_key = "organizationId"

  validates :openedAt, presence: true
  validates :fullSeats, presence: true
  validates :liteSeats, presence: true
  validates :developerSeats, presence: true
  validates :latestChangeAt, presence: true
  validates :mismatches, presence: true
end
