# A progressive checkout SESSION: one scan batch, covering N assets via the
# positional `assetIds` / `quantities` / `bookingAssetIds` arrays.
#
# Records sessions and the units they claimed — NOT whether an asset is out.
# The all-at-once checkout (`checkoutBooking`, the Check out button) flips
# statuses and writes NO row here, so **absence proves nothing** and must never
# be read as "was not checked out". Reading it that way is what let a single
# later scan mark every button-checked-out asset on a booking as
# "never checked out". `BookingAsset.checkedOutAt` answers that question.
class PartialBookingCheckout < ApplicationRecord
  self.table_name = "PartialBookingCheckout"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :assetIds, default: -> { [] }
  attribute :quantities, default: -> { [] }
  attribute :bookingAssetIds, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :checkoutCount, presence: true

  belongs_to :booking, foreign_key: "bookingId", inverse_of: :partial_checkouts
  belongs_to :checked_out_by, class_name: "User", foreign_key: "checkedOutById", inverse_of: :partial_checkouts
end
