class BookingSettings < ApplicationRecord
  self.table_name = "BookingSettings"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :bufferStartTime, default: 0
  attribute :tagsRequired, default: false
  attribute :maxBookingLengthSkipClosedDays, default: false
  attribute :autoArchiveBookings, default: false
  attribute :autoArchiveDays, default: 2
  attribute :autoArchiveExpiredReservations, default: false
  attribute :requireExplicitCheckinForAdmin, default: false
  attribute :requireExplicitCheckinForSelfService, default: false
  attribute :requireExplicitCheckoutForAdmin, default: false
  attribute :requireExplicitCheckoutForSelfService, default: false
  attribute :countKitsAsSingleUnit, default: false
  attribute :notifyBookingCreator, default: true
  attribute :notifyAdminsOnNewBooking, default: true
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, uniqueness: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :booking_settings
  has_and_belongs_to_many :always_notify_team_members, class_name: "TeamMember",
                                                       join_table: "_BookingSettingsAlwaysNotify", foreign_key: "A",
                                                       association_foreign_key: "B"
end
