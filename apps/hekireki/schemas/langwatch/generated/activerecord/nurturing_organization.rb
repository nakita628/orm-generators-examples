# Nurturing's read model per organization: its admin and its milestone counts.
# `seeded` marks an organization learned from the project backfill: its first_* already happened.
class NurturingOrganization < ApplicationRecord
  self.table_name = "NurturingOrganization"
  self.primary_key = "organizationId"

  attribute :seeded, default: false
  attribute :evaluationCount, default: 0
  attribute :simulationRunCount, default: 0
  attribute :codingAgentTypes, default: -> { [] }
  attribute :codingActiveDays, default: 0
  attribute :codingTerminalTypes, default: -> { [] }
  attribute :codingRepositoryKeys, default: -> { [] }
  attribute :codingFirstDays, default: -> { [] }
  alias_attribute :updated_at, :updatedAt
end
