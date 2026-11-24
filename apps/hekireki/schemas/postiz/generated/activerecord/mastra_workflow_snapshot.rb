class MastraWorkflowSnapshot < ApplicationRecord
  self.table_name = "mastra_workflow_snapshot"
  self.primary_key = %w[workflow_name run_id]

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :workflow_name, presence: true, uniqueness: { scope: :run_id }
  validates :run_id, presence: true
  validates :snapshot, presence: true
end
