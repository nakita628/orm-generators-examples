class AuditAssignment < ApplicationRecord
  self.table_name = "AuditAssignment"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :role, { lead: "LEAD", participant: "PARTICIPANT" }, validate: { allow_nil: true }

  validates :auditSessionId, uniqueness: { scope: :userId }

  belongs_to :audit_session, foreign_key: "auditSessionId", inverse_of: :assignments
  belongs_to :user, foreign_key: "userId", inverse_of: :audit_assignments
end
