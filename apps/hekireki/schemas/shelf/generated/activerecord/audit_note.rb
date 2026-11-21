class AuditNote < ApplicationRecord
  self.table_name = "AuditNote"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { comment: "COMMENT", update: "UPDATE" }, default: :comment, prefix: true, validate: true

  validates :content, presence: true

  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :audit_notes
  belongs_to :audit_session, foreign_key: "auditSessionId", inverse_of: :notes
  belongs_to :audit_asset, foreign_key: "auditAssetId", optional: true, inverse_of: :notes
end
