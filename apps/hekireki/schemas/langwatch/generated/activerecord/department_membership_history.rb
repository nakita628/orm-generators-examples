# One dated person → department link (ADR-128 §13, #7882).
#
# `OrganizationUser.departmentId` stays the CURRENT assignment every screen
# reads; these rows are its history, appended by the same `assignUser` write
# so the two cannot drift. They exist for the cost reads: January's spend
# stays with January's department, so a read resolves against the link that
# was open on the cost's day, not against today's pointer. Members with no
# link on a day resolve to "unassigned", never to an error.
#
# Same shape as IdentityMatch's dated links, same one-open-link guard: the
# partial unique index in the migration rejects a second open row per
# (organizationId, userId) with SQLSTATE 23505.
#
# validTo:: Null means the link is open. Reassignment and clearing close it.
class DepartmentMembershipHistory < ApplicationRecord
  self.table_name = "DepartmentMembershipHistory"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt

  validates :organizationId, presence: true
  validates :userId, presence: true
  validates :departmentId, presence: true
  validates :validFrom, presence: true
end
