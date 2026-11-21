# activity_id:: Soft reference to the PollActivity row this delivery was fanned out from; with webhookId it makes fan-out idempotent
# payload:: Frozen at fan-out: retries resend exactly what the first attempt sent
class WebhookDelivery < ApplicationRecord
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :attempts, default: 0

  enum :status, { pending: "pending", in_flight: "in_flight", succeeded: "succeeded", failed: "failed", exhausted: "exhausted" },
       default: :pending, validate: true

  validates :activity_id, presence: true, uniqueness: { scope: :webhook_id }
  validates :event_type, presence: true
  validates :payload, exclusion: { in: [ nil ] }

  belongs_to :webhook, class_name: "SpaceWebhook", inverse_of: :deliveries
end
