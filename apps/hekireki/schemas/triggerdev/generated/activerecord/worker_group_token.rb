class WorkerGroupToken < ApplicationRecord
  self.table_name = "WorkerGroupToken"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :tokenHash, presence: true, uniqueness: true
  validates :previousTokenHash, uniqueness: { allow_nil: true }

  belongs_to :child_of, class_name: "WorkerInstanceGroup", foreign_key: "workerGroupId", optional: true,
                        inverse_of: :child_tokens
  has_one :worker_group, class_name: "WorkerInstanceGroup", foreign_key: "tokenId", inverse_of: :token,
                         dependent: :destroy
end
