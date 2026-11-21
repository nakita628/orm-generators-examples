# calendarTokenId:: Secret, unguessable token for this member's iCal calendar-feed subscription
#   in this workspace. Null until the user generates a feed; regenerating it
#   rotates the value (invalidating any previously-shared subscription URL).
#   `@unique` doubles as the lookup index for the public feed route.
class UserOrganization < ApplicationRecord
  self.table_name = "UserOrganization"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :roles, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :userId, uniqueness: { scope: :organizationId }
  validates :calendarTokenId, uniqueness: { allow_nil: true }

  belongs_to :user, foreign_key: "userId", inverse_of: :user_organizations
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :user_organizations
end
