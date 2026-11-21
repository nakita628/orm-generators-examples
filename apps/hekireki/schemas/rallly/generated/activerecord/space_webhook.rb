# secret:: AES-256-GCM under SECRET_PASSWORD (see @rallly/utils/encryption); the plaintext is shown once at creation
# events:: Subscribed event names (`poll.closed`, ...); validated by the feature schema, not a DB enum, so the vocabulary grows without migrations
# version:: The payload contract this endpoint was built against (WEBHOOK_VERSION at creation). No default: a writer that forgets it fails instead of recording a stale contract. A breaking payload change branches on this so existing endpoints keep their shape
# cursor:: Fan-out position in the poll activity log: activities created after this instant are still to be delivered. Starts at creation so a new endpoint never replays history
# consecutive_failures:: Exhausted deliveries in a row; reset by any success. The dispatcher disables the endpoint once it reaches its cap
class SpaceWebhook < ApplicationRecord
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :events, default: -> { [] }
  attribute :enabled, default: true
  attribute :consecutive_failures, default: 0

  validates :url, presence: true
  validates :secret, presence: true
  validates :version, presence: true

  belongs_to :space, inverse_of: :webhooks
  has_many :deliveries, class_name: "WebhookDelivery", foreign_key: "webhook_id", inverse_of: :webhook,
                        dependent: :destroy
end
