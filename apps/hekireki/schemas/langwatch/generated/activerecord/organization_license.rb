# The licence an organization activated and its dates, owned by licensing
# (round 37 D6). A row whose key is null is a cleared licence. An organization
# with no row is read from its Organization columns until the copy step runs.
class OrganizationLicense < ApplicationRecord
  self.table_name = "OrganizationLicense"
  self.primary_key = "organizationId"

  alias_attribute :updated_at, :updatedAt
end
