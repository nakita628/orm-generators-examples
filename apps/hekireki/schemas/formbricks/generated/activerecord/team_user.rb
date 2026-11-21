# Links users to teams with specific roles.
# Manages team membership and permissions.
#
# @property team - The associated team
# @property user - The team member
# @property role - User's role within the team
class TeamUser < ApplicationRecord
  self.table_name = "TeamUser"
  self.primary_key = %w[teamId userId]

  enum :role, { admin: "admin", contributor: "contributor" }, validate: true

  belongs_to :team, foreign_key: "teamId", inverse_of: :team_users
  belongs_to :user, foreign_key: "userId", inverse_of: :team_users
end
