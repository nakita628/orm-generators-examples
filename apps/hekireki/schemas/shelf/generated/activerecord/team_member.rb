class TeamMember < ApplicationRecord
  self.table_name = "TeamMember"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :name, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :members
  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :team_members
  has_many :custodies, foreign_key: "teamMemberId", inverse_of: :custodian, dependent: :restrict_with_error
  has_many :received_invites, class_name: "Invite", foreign_key: "teamMemberId", inverse_of: :invitee_team_member,
                              dependent: :restrict_with_error
  has_many :kit_custodies, foreign_key: "custodianId", inverse_of: :custodian, dependent: :restrict_with_error
  has_many :bookings, foreign_key: "custodianTeamMemberId", inverse_of: :custodian_team_member, dependent: :destroy
  has_many :consumption_logs, foreign_key: "custodianId", inverse_of: :custodian, dependent: :nullify
  has_many :team_member_notes, foreign_key: "teamMemberId", inverse_of: :team_member, dependent: :destroy
  has_and_belongs_to_many :asset_reminders, join_table: "_AssetReminderToTeamMember", foreign_key: "B",
                                            association_foreign_key: "A"
  has_and_belongs_to_many :always_notify_booking_settings, class_name: "BookingSettings",
                                                           join_table: "_BookingSettingsAlwaysNotify", foreign_key: "B",
                                                           association_foreign_key: "A"
  has_and_belongs_to_many :booking_notifications, class_name: "Booking", join_table: "_BookingNotificationRecipients",
                                                  foreign_key: "B", association_foreign_key: "A"
end
