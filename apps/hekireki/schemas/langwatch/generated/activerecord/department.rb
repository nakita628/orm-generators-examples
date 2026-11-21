# A department is a single-valued accounting dimension, separate from
# RBAC, for rolling up AI spend by org unit (Engineering, Marketing,
# ...). Assigned to OrganizationUser, Team, and Project via their
# `departmentId` column. Pure accounting: never grants or restricts
# access. See specs/ai-gateway/governance/departments.feature.
class Department < ApplicationRecord
  self.table_name = "Department"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true
  validates :name, presence: true
end
