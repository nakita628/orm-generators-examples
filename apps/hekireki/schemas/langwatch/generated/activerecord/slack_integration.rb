# One named Slack connection (ADR-093 §5a); an organization holds as many as
# it likes, each scoped to the organization or to one project. The secret is
# AES-256-GCM ciphertext and never returned to a client. `secretFingerprint`
# is an HMAC of the plaintext, so one secret is one connection per scope
# without decrypting anything. No foreign key: `scopeId` is
# polymorphic (ADR-021), as on RetentionPolicy and CustomLLMModelCost.
class SlackIntegration < ApplicationRecord
  self.table_name = "SlackIntegration"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :kind, { bot: "BOT", incoming_webhook: "INCOMING_WEBHOOK" }, validate: true
  enum :scopeType, { organization: "ORGANIZATION", project: "PROJECT" }, validate: true

  validates :name, presence: true
  validates :scopeId, presence: true
  validates :organizationId, presence: true, uniqueness: { scope: [ :scopeType, :scopeId, :secretFingerprint ] }
  validates :secretFingerprint, presence: true
  validates :secretHint, presence: true
  validates :createdById, presence: true
  validates :updatedById, presence: true
end
