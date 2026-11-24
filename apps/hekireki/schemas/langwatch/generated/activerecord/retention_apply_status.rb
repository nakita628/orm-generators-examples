class RetentionApplyStatus < ApplicationRecord
  self.table_name = "RetentionApplyStatus"

  enum :scopeType, { organization: "ORGANIZATION", team: "TEAM", project: "PROJECT" }, validate: true

  validates :organizationId, presence: true, uniqueness: { scope: [ :scopeType, :scopeId ] }
  validates :scopeId, presence: true
  validates :createdAt, presence: true
  validates :updatedAt, presence: true
  validates :occurredAt, presence: true
  validates :acceptedAt, presence: true
  validates :lastEventId, presence: true
  validates :projectionVersion, presence: true
end
