# Defines team access to specific workspaces.
# Manages workspace-level permissions for teams.
#
# @property workspace - The accessed workspace
# @property team - The team receiving access
# @property permission - Level of access granted
class WorkspaceTeam < ApplicationRecord
  self.table_name = "WorkspaceTeam"
  self.primary_key = %w[workspaceId teamId]

  enum :permission, { read: "read", read_write: "readWrite", manage: "manage" }, default: :read, validate: true

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :workspace_teams
  belongs_to :team, foreign_key: "teamId", inverse_of: :workspace_teams
end
