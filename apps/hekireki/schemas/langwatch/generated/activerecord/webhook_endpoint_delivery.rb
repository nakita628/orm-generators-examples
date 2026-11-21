# The per-ATTEMPT delivery log for BOTH webhook channels: the endpoints
# platform's batches and the automations channel's trigger fires. One row per
# HTTP attempt; `dispatchId` groups every attempt of one logical send (the
# value of the X-LangWatch-Delivery-Id or X-LangWatch-Event-Id header). A slim
# facts table: outcome, status, latency, a capped error message, and a
# truncated copy of a failure response. Our request content (URL, headers,
# body) is never stored. Pruned after 30 days.
#
# The two channels used to keep separate tables with the same columns and two
# prune implementations. They are one table with a discriminator now, so the
# tenancy columns of the channel that did not write a row are null.
#
# organizationId:: Platform tenancy. Null on automations rows.
# projectId:: Automations tenancy. Null on platform rows.
# attempt:: 1-based attempt and the batch's envelope count. Both null on automations
#   rows, which record neither: one trigger fire is one event, and the
#   channel's sender does not thread its attempt number into the log.
class WebhookEndpointDelivery < ApplicationRecord
  self.table_name = "WebhookEndpointDelivery"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt

  enum :channel, { platform: "platform", automations: "automations" }, default: :platform, validate: true
  enum :outcome, { success: "success", retryable: "retryable", terminal: "terminal", pending: "pending" },
       validate: true

  validates :dispatchId, presence: true

  belongs_to :endpoint, class_name: "WebhookEndpoint", foreign_key: "endpointId", optional: true,
                        inverse_of: :deliveries
  belongs_to :project, foreign_key: "projectId", optional: true, inverse_of: :webhook_deliveries
  belongs_to :trigger, foreign_key: "triggerId", optional: true, inverse_of: :webhook_deliveries
end
