# Represents a team within an organization.
# Enables group-based access control and collaboration.
#
# @property id - Unique identifier for the team
# @property name - Display name of the team
# @property organization - The parent organization
# @property teamUsers - Users who are part of this team
# @property workspaceTeams - Workspaces this team has access to
class Team < ApplicationRecord
  self.table_name = "Team"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :name, presence: true, uniqueness: { scope: :organizationId }

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :teams
  has_many :team_users, foreign_key: "teamId", inverse_of: :team, dependent: :destroy
  has_many :workspace_teams, foreign_key: "teamId", inverse_of: :team, dependent: :destroy
end
