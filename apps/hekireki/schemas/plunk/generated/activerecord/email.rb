# scheduledFor:: When a transactional email sent with `delivery.at` is due. Only meaningful while the
#   status is SCHEDULED; the delayed queue job is the timer and the scheduled-email sweep
#   re-queues any row still SCHEDULED past this time. The partial index on SCHEDULED rows
#   is built by the background index builder, not a migration (see background-index-builder.ts).
# unsubscribedAt:: When the recipient opted out from this email (unsubscribe or snooze). Only set for
#   opt-outs attributed to a campaign email and counted in `Campaign.unsubscribedCount`,
#   in the same transaction, so the stamped rows are the list behind that figure. Indexed
#   by a partial index built in the background (see background-index-builder.ts).
# simulated:: The address handed to SES belonged to the mailbox simulator
#   (`@simulator.amazonses.com`), stamped at send time from the address actually used --
#   which is the recipient override when one is set, not necessarily `contact.email`.
#
#   AWS itself excludes simulator mail from bounce rates, complaint rates, sending quotas
#   and reputation metrics, so Plunk's security rates exclude it from both the numerator
#   and the denominator. The events themselves are untouched: the webhook fires, the
#   activity feed shows the bounce, the contact is still suppressed.
class Email < ApplicationRecord
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { SecureRandom.uuid }
  attribute :opens, default: 0
  attribute :clicks, default: 0
  attribute :simulated, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :sourceType, { transactional: "TRANSACTIONAL", campaign: "CAMPAIGN", workflow: "WORKFLOW", inbound: "INBOUND" },
       validate: true
  enum :status, { scheduled: "SCHEDULED", cancelled: "CANCELLED", pending: "PENDING", sending: "SENDING", sent: "SENT", delivered: "DELIVERED", received: "RECEIVED", opened: "OPENED", clicked: "CLICKED", bounced: "BOUNCED", complained: "COMPLAINED", failed: "FAILED" },
       default: :pending, validate: true

  validates :subject, presence: true
  validates :body, presence: true
  validates :from, presence: true
  validates :messageId, uniqueness: { allow_nil: true }

  belongs_to :contact, foreign_key: "contactId", inverse_of: :emails
  belongs_to :template, foreign_key: "templateId", optional: true, inverse_of: :emails
  belongs_to :campaign, foreign_key: "campaignId", optional: true, inverse_of: :emails
  belongs_to :workflow_execution, foreign_key: "workflowExecutionId", optional: true, inverse_of: :emails
  belongs_to :workflow_step_execution, foreign_key: "workflowStepExecutionId", optional: true, inverse_of: :emails
  belongs_to :project, foreign_key: "projectId", inverse_of: :emails
  has_many :events, foreign_key: "emailId", inverse_of: :email, dependent: :destroy
end
