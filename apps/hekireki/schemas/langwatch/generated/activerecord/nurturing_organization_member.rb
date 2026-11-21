# Nurturing's own copy of each organization's members, from the owners' member facts (§9):
# organization's signed_up and invite_accepted and auth's sso_auto_added join, member_removed
# and member_disabled leave, member_enabled returns. Each column keeps its latest time.
class NurturingOrganizationMember < ApplicationRecord
  self.table_name = "NurturingOrganizationMember"
  self.primary_key = %w[organizationId userId]
end
