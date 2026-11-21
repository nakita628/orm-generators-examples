# Holds one survey back from one policy until `until`. Keyed by organisation + entity rather than a
# policy id, so removing and re-adding a policy doesn't orphan its exemptions.
#
# Active means `revokedAt IS NULL AND until > now()`, so expiry needs no job. "One active exemption
# per survey and policy" is the partial unique index `RetentionExemption_surveyId_entity_active_key`
# (`WHERE "revokedAt" IS NULL`) in
# `migration/20261007120000_add_data_retention_tables/migration.sql`; Prisma's `@@unique` has no
# `where`, so it is deliberately not declared here (see "Indexes Prisma cannot express" in the
# README). An expired row still holds that slot until it is revoked, so creating a new exemption
# must close an expired one in the same transaction.
#
# `surveyId` is nullable and a CHECK requires exactly one target, so a later `workflowId` is one more
# column in that CHECK. The survey's organisation can't be tied to `organizationId` in the database
# (a survey carries its workspace, not its organisation): writers derive it from the survey.
#
# reason:: Required for `manual`, absent for `restore` (`RetentionExemption_reason_check`).
class RetentionExemption < ApplicationRecord
  self.table_name = "RetentionExemption"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid2.call }

  enum :entity, { responses: "responses", surveys: "surveys", members: "members" }, validate: true
  enum :source, { manual: "manual", restore: "restore" }, default: :manual, validate: true

  validates :until, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :retention_exemptions
  belongs_to :survey, foreign_key: "surveyId", optional: true, inverse_of: :retention_exemptions
  belongs_to :created_by, class_name: "User", foreign_key: "createdById", optional: true,
                          inverse_of: :exemptions_created
  belongs_to :revoked_by, class_name: "User", foreign_key: "revokedById", optional: true,
                          inverse_of: :exemptions_revoked
end
