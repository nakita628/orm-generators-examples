# Postgres-authoritative one-turn admission for a conversation. PREPARING
# claims are leased so a crash before the durable start event can recover;
# COMMITTED claims are released only by the matching terminal event.
class LangyActiveTurn < ApplicationRecord
  self.table_name = "LangyActiveTurn"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :projectId, presence: true, uniqueness: { scope: :conversationId }
  validates :conversationId, presence: true
  validates :turnId, presence: true
  validates :requestId, presence: true
  validates :userId, presence: true
  validates :status, presence: true
  validates :leaseOwner, presence: true
  validates :leaseExpiresAt, presence: true
end
