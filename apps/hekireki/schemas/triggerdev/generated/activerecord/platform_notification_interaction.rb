# Per-user tracking of notification views and dismissals
#
# firstSeenAt:: Set by beacon/CLI GET on first impression
# showCount:: Times shown (incremented per beacon/CLI view)
# webappDismissedAt:: User dismissed in webapp
# webappClickedAt:: User clicked a link in webapp
# cliDismissedAt:: Auto-dismissed or expired in CLI
class PlatformNotificationInteraction < ApplicationRecord
  self.table_name = "PlatformNotificationInteraction"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :showCount, default: 1
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :notificationId, uniqueness: { scope: :userId }

  belongs_to :notification, class_name: "PlatformNotification", foreign_key: "notificationId", inverse_of: :interactions
  belongs_to :user, foreign_key: "userId", inverse_of: :platform_notification_interactions
end
