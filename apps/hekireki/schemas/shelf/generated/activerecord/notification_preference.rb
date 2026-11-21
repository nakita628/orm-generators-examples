# A person's push notification switches. A missing row means every switch
# is on, so no backfill is needed. The three category columns are named
# after `PushCategory` in apps/webapp/app/modules/push/types.ts.
class NotificationPreference < ApplicationRecord
  self.table_name = "NotificationPreference"
  self.primary_key = "userId"

  attribute :pushEnabled, default: true
  attribute :bookings, default: true
  attribute :audits, default: true
  attribute :stockAndReminders, default: true
  alias_attribute :updated_at, :updatedAt

  belongs_to :user, foreign_key: "userId", inverse_of: :notification_preference
end
