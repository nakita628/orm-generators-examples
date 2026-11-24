# An organization a grant with `organizations` access reaches, workspaces created there later included.
class McpGrantOrganization < ApplicationRecord
  self.table_name = "McpGrantOrganization"
  self.primary_key = %w[grantId organizationId]

  belongs_to :grant, class_name: "McpGrant", foreign_key: "grantId", inverse_of: :organizations
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :mcp_grants
end
