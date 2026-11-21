# One policy per organisation and kind of data: a notice `warnDays` before the action, and the action
# `periodDays` after the target's clock. The action follows from the kind of data: responses are
# deleted, surveys archived (the archive purge deletes them 30 days later, as it does every archived
# survey), members deactivated. The dates themselves are never stored; they are computed by the shared
# helper in `apps/web/modules/ee/data-retention/lib/schedule.ts`, so the UI, the emails and the sweep
# always agree.
#
# The allowed ranges are validated by the API. The migration repeats the ones that guard against
# acting early or re-notifying (a notice of at least 14 days and shorter than the period, an enabled
# policy always has `enabledAt`) and the structural ones (conditions on the surveys policy only) as
# CHECK constraints, which Prisma cannot express.
#
# enabledAt:: When the current configuration took effect: the policy was switched on or unpaused, or changed
#   in a way that could bring an action earlier (a shorter period or warning, different
#   conditions). Lengthening does not move it. The warning always runs in full from here, so no
#   change can delete data on the night it is made (ENG-3614). Also the clock for members with no
#   recorded sign-in.
# warnDays:: How long before the action the notice goes out. The action never runs before the notice has
#   been out for this long, so the warning always runs in full.
# periodDays:: The period, counted from the target's clock, after which the policy acts: responses are
#   deleted, surveys archived (the archive purge deletes them 30 days later), members deactivated.
# conditions:: Surveys policy only.
# deferredAt:: The last night the sweep's time ran out before it reached this policy's organisation, licensed.
#   The organisation was due and the sweep was running, so it counts as reached for the gap that
#   restarts the warning (`RETENTION_SWEEP_GAP_MS`): a backlog of organisations never restarts warnings.
class RetentionPolicy < ApplicationRecord
  self.table_name = "RetentionPolicy"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid2.call }
  attribute :enabled, default: false
  attribute :warnDays, default: 60

  enum :entity, { responses: "responses", surveys: "surveys", members: "members" }, validate: true

  validates :organizationId, uniqueness: { scope: :entity }
  validates :periodDays, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :retention_policies
  belongs_to :updated_by, class_name: "User", foreign_key: "updatedById", optional: true,
                          inverse_of: :retention_policies_updated
end
