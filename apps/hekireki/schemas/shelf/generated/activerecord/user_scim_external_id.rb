# *
# * Org-scoped SCIM external ID linking a Shelf user to an identity-provider
# * object within a specific organization.  A single Shelf user can appear in
# * multiple organizations, each managed by a different identity provider, so the
# * external ID must be scoped to the (user, organization) pair.
class UserScimExternalId < ApplicationRecord
  self.table_name = "UserScimExternalId"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :userId, uniqueness: { scope: :organizationId }
  validates :scimExternalId, presence: true, uniqueness: { scope: :organizationId }

  belongs_to :user, foreign_key: "userId", inverse_of: :scim_external_ids
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :scim_external_ids
end
