class Booking < ApplicationRecord
  self.table_name = "Booking"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :description, default: ""
  attribute :archivedWithoutCheckin, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :status, { draft: "DRAFT", reserved: "RESERVED", ongoing: "ONGOING", overdue: "OVERDUE", complete: "COMPLETE", archived: "ARCHIVED", cancelled: "CANCELLED" },
       default: :draft, validate: true

  validates :name, presence: true
  validates :from, presence: true
  validates :to, presence: true

  belongs_to :creator, class_name: "User", foreign_key: "creatorId", inverse_of: :bookings
  belongs_to :custodian_user, class_name: "User", foreign_key: "custodianUserId", optional: true, inverse_of: :custodies
  belongs_to :custodian_team_member, class_name: "TeamMember", foreign_key: "custodianTeamMemberId", optional: true,
                                     inverse_of: :bookings
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :bookings
  has_many :booking_assets, foreign_key: "bookingId", inverse_of: :booking, dependent: :destroy
  has_many :consumption_logs, foreign_key: "bookingId", inverse_of: :booking, dependent: :nullify
  has_many :notes, class_name: "BookingNote", foreign_key: "bookingId", inverse_of: :booking, dependent: :destroy
  has_many :model_requests, class_name: "BookingModelRequest", foreign_key: "bookingId", inverse_of: :booking,
                            dependent: :destroy
  has_many :partial_checkins, class_name: "PartialBookingCheckin", foreign_key: "bookingId", inverse_of: :booking,
                              dependent: :destroy
  has_many :partial_checkouts, class_name: "PartialBookingCheckout", foreign_key: "bookingId", inverse_of: :booking,
                               dependent: :destroy
  has_and_belongs_to_many :tags, join_table: "_BookingToTag", foreign_key: "A", association_foreign_key: "B"
  has_and_belongs_to_many :notification_recipients, class_name: "TeamMember",
                                                    join_table: "_BookingNotificationRecipients", foreign_key: "A",
                                                    association_foreign_key: "B"
end
