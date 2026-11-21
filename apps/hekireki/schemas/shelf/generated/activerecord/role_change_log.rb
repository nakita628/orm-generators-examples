# source:: How the change was made. SSO rows name the member themself as `changedBy`.
class RoleChangeLog < ApplicationRecord
  self.table_name = "RoleChangeLog"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt

  enum :previousRole, { admin: "ADMIN", base: "BASE", owner: "OWNER", self_service: "SELF_SERVICE", manager: "MANAGER" },
       validate: true
  enum :newRole, { admin: "ADMIN", base: "BASE", owner: "OWNER", self_service: "SELF_SERVICE", manager: "MANAGER" },
       prefix: true, validate: true
  enum :source, { manual: "MANUAL", sso: "SSO" }, default: :manual, validate: true

  belongs_to :user, foreign_key: "userId", inverse_of: :role_changes_received
  belongs_to :changed_by, class_name: "User", foreign_key: "changedById", inverse_of: :role_changes_performed
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :role_change_logs
end
