defmodule Formbricks.RetentionNotice do
  @moduledoc """
  The notice for one target under one policy, so each notice goes out once: a survey or a member, or
  the one-time responses reminder for a survey. It counts only once delivered, and is void if it was
  claimed before the policy's `enabledAt`, before the target's latest exemption ended or (surveys and
  members) before the target's clock: then a later cycle claims it again and sends a new one.

  A CHECK in the migration ties each entity to its target: `surveys` and `responses` → `surveyId`,
  `members` → `userId`; another makes a sent email imply a delivery time. The survey notices that still
  hold a clock are the partial index `RetentionNotice_surveys_clock_pending_idx` (`WHERE "entity" =
  'surveys' AND "clockAt" IS NOT NULL`) in
  `migration/20261012120000_retention_notice_void_on_response_delete/migration.sql`, with the triggers
  that clear `clockAt`; Prisma's `@@index` has no `where`, so it is deliberately not declared here (see
  "Indexes Prisma cannot express" in the README).
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          entity: atom(),
          sent_at: DateTime.t(),
          delivered_at: DateTime.t() | nil,
          email_sent: boolean(),
          claim_token: String.t() | nil,
          clock_at: DateTime.t() | nil,
          organization: Formbricks.Organization.t() | nil,
          survey: Formbricks.Survey.t() | nil,
          user: Formbricks.User.t() | nil
        }

  schema "RetentionNotice" do
    field(:id, :string, primary_key: true)
    field(:entity, Ecto.Enum, values: [:responses, :surveys, :members])
    # When the notice was claimed for sending. A claim that never got delivered can be taken over.
    field(:sent_at, Formbricks.PrismaDateTime, autogenerate: true, source: :sentAt)
    # When the notice was handed to the mail transport, or recorded without one (no SMTP). The
    # warning only counts from a delivered notice, so a crash between claim and send can't let the
    # action run without it.
    field(:delivered_at, Formbricks.PrismaDateTime, source: :deliveredAt)
    # Whether an email actually went out. False when SMTP isn't configured.
    field(:email_sent, :boolean, default: false, source: :emailSent)
    # Set by each claim; delivery is recorded only for the claim that sent the email, so a slow sender
    # can't mark a newer claim delivered.
    field(:claim_token, :string, source: :claimToken)
    # The target's clock the notice was computed for (surveys and members). A notice counts only while
    # the target's clock is still this one, so activity between the sweep's read and its claim voids
    # it. Null for the responses reminder, which is sent once per survey and follows no one clock, and
    # for a survey notice voided for good: a person deleted or moved one of its responses, or the
    # responses policy's deletion left it computed for another clock (see `NOTICE_FOLLOWS_CLOCK`).
    field(:clock_at, Formbricks.PrismaDateTime, source: :clockAt)
    field(:organization_id, :string, source: :organizationId)
    field(:survey_id, :string, source: :surveyId)
    field(:user_id, :string, source: :userId)
    belongs_to(:organization, Formbricks.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:survey, Formbricks.Survey, foreign_key: :survey_id, define_field: false, type: :string)
    belongs_to(:user, Formbricks.User, foreign_key: :user_id, define_field: false, type: :string)
  end
end