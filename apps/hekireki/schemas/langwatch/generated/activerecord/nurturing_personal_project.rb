# Nurturing's copy of each personal project's owner, from organization's and project's facts.
# The first owner recorded stays: a personal project's owner never changes.
class NurturingPersonalProject < ApplicationRecord
  self.table_name = "NurturingPersonalProject"
  self.primary_key = "projectId"

  validates :organizationId, presence: true
  validates :ownerUserId, presence: true
end
