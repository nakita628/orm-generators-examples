# ADR-201: an organisation's braid setting. Owned by authz; read into the organisation's snapshot.
class AuthzOrganizationPolicy < ApplicationRecord
  self.table_name = "AuthzOrganizationPolicy"
  self.primary_key = "organizationId"

  attribute :braidsOrganizationAdminsOnly, default: false
  alias_attribute :updated_at, :updatedAt

  validates :updatedByUserId, presence: true
end
