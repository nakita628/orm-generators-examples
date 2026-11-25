# One person's digest for one ISO week, accepted by our own mailer. The id is `userId:isoWeek`.
class DigestSend < ApplicationRecord
  self.table_name = "DigestSend"

  alias_attribute :created_at, :createdAt

  validates :userId, presence: true
  validates :organizationId, presence: true
  validates :isoWeek, presence: true
  validates :templateId, presence: true
  validates :sentAt, presence: true
end
