defmodule Plunk.Email do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Plunk.PrismaDateTime, autogenerate: {Plunk.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          to_name: String.t() | nil,
          subject: String.t(),
          body: String.t(),
          from: String.t(),
          from_name: String.t() | nil,
          reply_to: String.t() | nil,
          headers: map() | list() | String.t() | number() | boolean() | nil,
          attachments: map() | list() | String.t() | number() | boolean() | nil,
          message_id: String.t() | nil,
          source_type: atom(),
          status: atom(),
          scheduled_for: DateTime.t() | nil,
          sent_at: DateTime.t() | nil,
          delivered_at: DateTime.t() | nil,
          opened_at: DateTime.t() | nil,
          clicked_at: DateTime.t() | nil,
          bounced_at: DateTime.t() | nil,
          complained_at: DateTime.t() | nil,
          unsubscribed_at: DateTime.t() | nil,
          opens: integer(),
          clicks: integer(),
          simulated: boolean(),
          error: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          contact: Plunk.Contact.t() | nil,
          template: Plunk.Template.t() | nil,
          campaign: Plunk.Campaign.t() | nil,
          workflow_execution: Plunk.WorkflowExecution.t() | nil,
          workflow_step_execution: Plunk.WorkflowStepExecution.t() | nil,
          project: Plunk.Project.t() | nil,
          events: [Plunk.Event.t()]
        }

  schema "emails" do
    field(:to_name, :string, source: :toName)
    field(:subject, :string)
    field(:body, :string)
    field(:from, :string)
    field(:from_name, :string, source: :fromName)
    field(:reply_to, :string, source: :replyTo)
    field(:headers, Plunk.PrismaJson)
    field(:attachments, Plunk.PrismaJson)
    field(:message_id, :string, source: :messageId)
    field(:source_type, Ecto.Enum, values: [:TRANSACTIONAL, :CAMPAIGN, :WORKFLOW, :INBOUND], source: :sourceType)
    field(:status, Ecto.Enum, values: [:SCHEDULED, :CANCELLED, :PENDING, :SENDING, :SENT, :DELIVERED, :RECEIVED, :OPENED, :CLICKED, :BOUNCED, :COMPLAINED, :FAILED], default: :PENDING)
    # When a transactional email sent with `delivery.at` is due. Only meaningful while the
    # status is SCHEDULED; the delayed queue job is the timer and the scheduled-email sweep
    # re-queues any row still SCHEDULED past this time. The partial index on SCHEDULED rows
    # is built by the background index builder, not a migration (see background-index-builder.ts).
    field(:scheduled_for, Plunk.PrismaDateTime, source: :scheduledFor)
    field(:sent_at, Plunk.PrismaDateTime, source: :sentAt)
    field(:delivered_at, Plunk.PrismaDateTime, source: :deliveredAt)
    field(:opened_at, Plunk.PrismaDateTime, source: :openedAt)
    field(:clicked_at, Plunk.PrismaDateTime, source: :clickedAt)
    field(:bounced_at, Plunk.PrismaDateTime, source: :bouncedAt)
    field(:complained_at, Plunk.PrismaDateTime, source: :complainedAt)
    # When the recipient opted out from this email (unsubscribe or snooze). Only set for
    # opt-outs attributed to a campaign email and counted in `Campaign.unsubscribedCount`,
    # in the same transaction, so the stamped rows are the list behind that figure. Indexed
    # by a partial index built in the background (see background-index-builder.ts).
    field(:unsubscribed_at, Plunk.PrismaDateTime, source: :unsubscribedAt)
    field(:opens, :integer, default: 0)
    field(:clicks, :integer, default: 0)
    # The address handed to SES belonged to the mailbox simulator
    # (`@simulator.amazonses.com`), stamped at send time from the address actually used --
    # which is the recipient override when one is set, not necessarily `contact.email`.
    #
    # AWS itself excludes simulator mail from bounce rates, complaint rates, sending quotas
    # and reputation metrics, so Plunk's security rates exclude it from both the numerator
    # and the denominator. The events themselves are untouched: the webhook fires, the
    # activity feed shows the bounce, the contact is still suppressed.
    field(:simulated, :boolean, default: false)
    field(:error, :string)
    field(:contact_id, :string, source: :contactId)
    field(:template_id, :string, source: :templateId)
    field(:campaign_id, :string, source: :campaignId)
    field(:workflow_execution_id, :string, source: :workflowExecutionId)
    field(:workflow_step_execution_id, :string, source: :workflowStepExecutionId)
    field(:project_id, :string, source: :projectId)
    belongs_to(:contact, Plunk.Contact, foreign_key: :contact_id, define_field: false)
    belongs_to(:template, Plunk.Template, foreign_key: :template_id, define_field: false)
    belongs_to(:campaign, Plunk.Campaign, foreign_key: :campaign_id, define_field: false)
    belongs_to(:workflow_execution, Plunk.WorkflowExecution, foreign_key: :workflow_execution_id, define_field: false)
    belongs_to(:workflow_step_execution, Plunk.WorkflowStepExecution, foreign_key: :workflow_step_execution_id, define_field: false)
    belongs_to(:project, Plunk.Project, foreign_key: :project_id, define_field: false)
    has_many(:events, Plunk.Event, foreign_key: :email_id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end