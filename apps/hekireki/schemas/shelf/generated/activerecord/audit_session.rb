class AuditSession < ApplicationRecord
  self.table_name = "AuditSession"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :expectedAssetCount, default: 0
  attribute :foundAssetCount, default: 0
  attribute :missingAssetCount, default: 0
  attribute :unexpectedAssetCount, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :status, { pending: "PENDING", active: "ACTIVE", completed: "COMPLETED", cancelled: "CANCELLED", archived: "ARCHIVED" },
       default: :pending, validate: true

  validates :name, presence: true

  belongs_to :created_by, class_name: "User", foreign_key: "createdById", inverse_of: :audits_created
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :audit_sessions
  has_many :assignments, class_name: "AuditAssignment", foreign_key: "auditSessionId", inverse_of: :audit_session,
                         dependent: :destroy
  has_many :assets, class_name: "AuditAsset", foreign_key: "auditSessionId", inverse_of: :audit_session,
                    dependent: :destroy
  has_many :scans, class_name: "AuditScan", foreign_key: "auditSessionId", inverse_of: :audit_session,
                   dependent: :destroy
  has_many :notes, class_name: "AuditNote", foreign_key: "auditSessionId", inverse_of: :audit_session,
                   dependent: :destroy
  has_many :images, class_name: "AuditImage", foreign_key: "auditSessionId", inverse_of: :audit_session,
                    dependent: :destroy
end
