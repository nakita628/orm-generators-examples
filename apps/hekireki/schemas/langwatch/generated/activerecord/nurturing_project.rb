# Nurturing's read model: which organization a project belongs to, from project's created event.
# The per-email usage places a project and lists an organization's projects from these rows only.
class NurturingProject < ApplicationRecord
  self.table_name = "NurturingProject"
  self.primary_key = "projectId"

  validates :organizationId, presence: true
end
