# The notice for one target under one policy, so each notice goes out once: a survey or a member, or
# the one-time responses reminder for a survey. It counts only once delivered, and is void if it was
# claimed before the policy's `enabledAt`, before the target's latest exemption ended or (surveys and
# members) before the target's clock: then a later cycle claims it again and sends a new one.
#
# A CHECK in the migration ties each entity to its target: `surveys` and `responses` → `surveyId`,
# `members` → `userId`; another makes a sent email imply a delivery time. The survey notices that still
# hold a clock are the partial index `RetentionNotice_surveys_clock_pending_idx` (`WHERE "entity" =
# 'surveys' AND "clockAt" IS NOT NULL`) in
# `migration/20261012120000_retention_notice_void_on_response_delete/migration.sql`, with the triggers
# that clear `clockAt`; Prisma's `@@index` has no `where`, so it is deliberately not declared here (see
# "Indexes Prisma cannot express" in the README).
#
# sentAt:: When the notice was claimed for sending. A claim that never got delivered can be taken over.
# deliveredAt:: When the notice was handed to the mail transport, or recorded without one (no SMTP). The
#   warning only counts from a delivered notice, so a crash between claim and send can't let the
#   action run without it.
# emailSent:: Whether an email actually went out. False when SMTP isn't configured.
# claimToken:: Set by each claim; delivery is recorded only for the claim that sent the email, so a slow sender
#   can't mark a newer claim delivered.
# clockAt:: The target's clock the notice was computed for (surveys and members). A notice counts only while
#   the target's clock is still this one, so activity between the sweep's read and its claim voids
#   it. Null for the responses reminder, which is sent once per survey and follows no one clock, and
#   for a survey notice voided for good: a person deleted or moved one of its responses, or the
#   responses policy's deletion left it computed for another clock (see `NOTICE_FOLLOWS_CLOCK`).
class RetentionNotice < ApplicationRecord
  self.table_name = "RetentionNotice"

  attribute :id, default: -> { Cuid2.call }
  attribute :emailSent, default: false

  enum :entity, { responses: "responses", surveys: "surveys", members: "members" }, validate: true

  validates :surveyId, uniqueness: { scope: :entity, allow_nil: true }
  validates :userId, uniqueness: { scope: [ :organizationId, :entity ], allow_nil: true }

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :retention_notices
  belongs_to :survey, foreign_key: "surveyId", optional: true, inverse_of: :retention_notices
  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :retention_notices
end
