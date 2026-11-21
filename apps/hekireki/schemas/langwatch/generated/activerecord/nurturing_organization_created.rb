# When an organization was created, from organization's created event: nurturing's own copy for
# the per-email usage's day offsets. Its own table, so an early creation never adds an unseeded
# NurturingOrganization row.
class NurturingOrganizationCreated < ApplicationRecord
  self.table_name = "NurturingOrganizationCreated"
  self.primary_key = "organizationId"

  alias_attribute :created_at, :createdAt
end
