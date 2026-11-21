# A progressive check-in SESSION, mirroring PartialBookingCheckout.
#
# `assetIds` lists only FULLY reconciled assets: INDIVIDUAL always, and
# QUANTITY_TRACKED only once its remaining units hit zero (partial returns
# live in `ConsumptionLog`). Completing a booking outright writes no row here,
# so the same warning applies — absence is not evidence of anything.
# `BookingAsset.checkedInAt` is the per-slice fact.
class PartialBookingCheckin < ApplicationRecord
  self.table_name = "PartialBookingCheckin"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :assetIds, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :checkinCount, presence: true

  belongs_to :booking, foreign_key: "bookingId", inverse_of: :partial_checkins
  belongs_to :checked_in_by, class_name: "User", foreign_key: "checkedInById", inverse_of: :partial_checkins
end
