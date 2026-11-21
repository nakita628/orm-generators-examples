# rbacRoleId:: Optional RBAC role to assign on invite acceptance. When set, the
#   accept-invite flow calls the loaded RBAC plugin's setUserRole with
#   this id after creating the OrgMember. Null = legacy behaviour, the
#   runtime fallback derives the role from `role` above.
#
#   Plain text (not an FK) — the RBAC plugin's RbacRole table lives on
#   a separate schema (Drizzle, not Prisma) so we can't model the FK
#   here. Validation happens at write time (action) and read time
#   (acceptInvite).
class OrgMemberInvite < ApplicationRecord
  self.table_name = "OrgMemberInvite"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :token, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :role, { admin: "ADMIN", member: "MEMBER" }, default: :member, validate: true

  validates :token, uniqueness: true
  validates :email, presence: true, uniqueness: { scope: :organizationId }

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :invites
  belongs_to :inviter, class_name: "User", foreign_key: "inviterId", inverse_of: :sent_invites
end
