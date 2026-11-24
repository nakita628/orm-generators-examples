# Per-call data kept beside a workflow (ADR 0048): one row per source, `trigger` for the Webhook
# trigger's last payload (plus the last call's time and IP), or an HTTP request step id for its last
# Test request response. Written on every call, so it lives here and never bumps `Workflow.updatedAt`.
#
# sample:: [WorkflowSampleData]
class WorkflowSample < ApplicationRecord
  self.table_name = "WorkflowSample"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid2.call }

  validates :source, presence: true, uniqueness: { scope: :workflowId }
  validates :sample, exclusion: { in: [ nil ] }

  belongs_to :workflow, foreign_key: [ :workflowId, :workspaceId ], primary_key: [ :id, :workspaceId ],
                        inverse_of: :samples
end
