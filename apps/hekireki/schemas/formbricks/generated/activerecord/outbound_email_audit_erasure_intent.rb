# Holds an account erasure request across an in-flight email send. The address is never stored.
# The user id is removed once deletion is confirmed; one-way recovery keys expire after 90 days.
class OutboundEmailAuditErasureIntent < ApplicationRecord
  self.table_name = "OutboundEmailAuditErasureIntent"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { SecureRandom.uuid }
  alias_attribute :created_at, :createdAt

  validates :userHash, presence: true
  validates :recipientMac, presence: true
  validates :expiresAt, presence: true
end
