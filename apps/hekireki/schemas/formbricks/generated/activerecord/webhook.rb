# Represents a webhook endpoint for receiving survey-related events.
# Webhooks can be configured to receive notifications about response creation, updates, and completion.
#
# @property id - Unique identifier for the webhook
# @property name - Optional display name for the webhook
# @property url - The endpoint URL where events will be sent
# @property source - Origin of the webhook (user, zapier, make, etc.)
# @property workspace - Associated workspace
# @property triggers - Types of events that trigger this webhook
# @property surveyIds - List of surveys this webhook is monitoring
class Webhook < ApplicationRecord
  self.table_name = "Webhook"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :triggers, default: -> { [] }
  attribute :surveyIds, default: -> { [] }

  enum :source, { user: "user", zapier: "zapier", make: "make", n8n: "n8n", activepieces: "activepieces" },
       default: :user, validate: true

  validates :url, presence: true

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :webhooks
end
