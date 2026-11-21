# Who authz treats as gone, from user's and identity's facts; authz never reads User (ARCHITECTURE.md, "Platform operators are a grant").
class AuthzUserStanding < ApplicationRecord
  self.table_name = "AuthzUserStanding"
  self.primary_key = "userId"

  validates :standingChangedAt, presence: true
end
