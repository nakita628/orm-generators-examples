# A successful callback through one SSO connection. Identifier.lastUsedAt
# also advances during routing and resolution, so it cannot prove that a
# member actually authenticated through the replacement before cutover.
#
# providerAccountId:: The provider-side account the assertion named; null before #7633's arrival door recorded it.
class SsoAuthenticationActivity < ApplicationRecord
  self.table_name = "SsoAuthenticationActivity"

  validates :organizationId, presence: true
  validates :connectionId, presence: true
  validates :userId, presence: true
  validates :authenticatedAt, presence: true
end
