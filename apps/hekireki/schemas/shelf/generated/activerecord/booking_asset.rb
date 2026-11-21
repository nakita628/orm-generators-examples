# *
# * Explicit pivot table between Booking and Asset.
# * Replaces the implicit M2M (_AssetToBooking) via a rename migration in Phase 3.
class BookingAsset < ApplicationRecord
  self.table_name = "BookingAsset"

  attribute :id, default: -> { Cuid.generate }
  attribute :quantity, default: 1
  attribute :checkedOutQuantity, default: 0

  belongs_to :booking, foreign_key: "bookingId", inverse_of: :booking_assets
  belongs_to :asset, foreign_key: "assetId", inverse_of: :booking_assets
end
