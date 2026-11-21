# Admin-created notification definitions
#
# title:: Admin-facing title for identification in admin tools
# payload:: Versioned JSON rendering content (see payload schema in spec)
# user:: Set when scope = USER
# organization:: Set when scope = ORGANIZATION
# project:: Set when scope = PROJECT
# startsAt:: When notification becomes active
# endsAt:: When notification expires
# cliMaxDaysAfterFirstSeen:: CLI: auto-expire N days after user first saw it
# cliMaxShowCount:: CLI: auto-expire after shown N times to user
# cliShowEvery:: CLI: only show every N-th request (e.g. cliShowEvery=3 means show on 3rd, 6th, 9th…)
# priority:: Ordering within same scope level (higher = more important)
# isDraft:: Staged-but-unpublished. While true the notification is hidden from all
#   user-facing reads regardless of startsAt/endsAt; publishing sets real
#   dates and flips this to false.
# archivedAt:: Soft delete
class PlatformNotification < ApplicationRecord
  self.table_name = "PlatformNotification"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :friendlyId, default: -> { Cuid.generate }
  attribute :priority, default: 0
  attribute :isDraft, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :surface, { webapp: "WEBAPP", cli: "CLI" }, validate: true
  enum :scope, { user: "USER", project: "PROJECT", organization: "ORGANIZATION", global: "GLOBAL" }, validate: true

  validates :friendlyId, uniqueness: true
  validates :title, presence: true
  validates :payload, exclusion: { in: [ nil ] }
  validates :startsAt, presence: true
  validates :endsAt, presence: true

  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :platform_notifications
  belongs_to :organization, foreign_key: "organizationId", optional: true, inverse_of: :platform_notifications
  belongs_to :project, foreign_key: "projectId", optional: true, inverse_of: :platform_notifications
  has_many :interactions, class_name: "PlatformNotificationInteraction", foreign_key: "notificationId",
                          inverse_of: :notification, dependent: :destroy
end
