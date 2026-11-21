# One experiment setting per (flag, subject). A USER row is that person's
# own opt-in; an ORGANIZATION or PROJECT row is an owner's explicit policy
# for that exact scope. No row means inherit.
class FeatureFlagExperimentSetting < ApplicationRecord
  self.table_name = "FeatureFlagExperimentSetting"
  self.primary_key = %w[flagKey subjectType subjectId]

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :subjectType, { user: "USER", organization: "ORGANIZATION", project: "PROJECT" }, validate: true

  validates :enabled, inclusion: { in: [ true, false ] }
end
