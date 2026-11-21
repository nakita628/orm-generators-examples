# Nurturing's copy of each project that received a real trace, from trace's and project's facts.
class NurturingTracedProject < ApplicationRecord
  self.table_name = "NurturingTracedProject"
  self.primary_key = "projectId"

  validates :organizationId, presence: true
end
