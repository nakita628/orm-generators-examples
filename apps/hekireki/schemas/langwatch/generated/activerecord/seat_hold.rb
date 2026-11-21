# A full seat still counted after the person who took it stopped holding it (removed,
# disabled or demoted), until 30 days after they took it (specs/licensing/seat-rotation.feature).
# One per organization and person; a return within the hold takes the seat back.
#
# email:: Display copy for the Users list, kept because a removed person has no membership row.
# seatTakenAt:: When the person took the seat this hold keeps.
# heldUntil:: seatTakenAt plus 30 days; expired holds are ignored by every count and deleted hourly.
# reason:: REMOVED | DISABLED | DEMOTED
class SeatHold < ApplicationRecord
  self.table_name = "SeatHold"

  alias_attribute :created_at, :createdAt

  validates :organizationId, presence: true, uniqueness: { scope: :userId }
  validates :userId, presence: true
  validates :seatTakenAt, presence: true
  validates :heldUntil, presence: true
  validates :reason, presence: true
end
