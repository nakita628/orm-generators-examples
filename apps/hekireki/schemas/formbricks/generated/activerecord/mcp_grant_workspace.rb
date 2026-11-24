# A workspace a grant with `workspaces` access reaches.
class McpGrantWorkspace < ApplicationRecord
  self.table_name = "McpGrantWorkspace"
  self.primary_key = %w[grantId workspaceId]

  belongs_to :grant, class_name: "McpGrant", foreign_key: "grantId", inverse_of: :workspaces
  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :mcp_grants
end
