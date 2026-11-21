# One member's or pending invitation's seat as the seat peak last applied it (ADR-002).
#
# holderKind:: `member`, `invitation` or `hold` (a seat held after somebody left).
# holderId:: The member's or the hold's userId, or the invitation's id.
# seat:: `full`, `lite`, `developer` or `none`.
# lastChangeId:: Absent for a seat the opening snapshot recorded.
# expiresAt:: An invitation's expiry; null for a member or an invitation that never expires.
# expiredByFold:: The fold gave the invitation's seat back at `expiresAt`.
# suppressedBy:: A hold whose seat a live full invitation to its address took: that invitation's id.
# restoredByFold:: The fold gave a suppressed hold its seat back when that invitation expired.
class OrganizationSeatPeakHolder < ApplicationRecord
  self.table_name = "OrganizationSeatPeakHolder"
  self.primary_key = %w[organizationId holderKind holderId]

  attribute :expiredByFold, default: false
  attribute :restoredByFold, default: false

  validates :seat, presence: true
  validates :lastChangedAt, presence: true
end
