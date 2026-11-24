# triggerPayload:: [WorkflowTriggerRunPayload]
# data:: [WorkflowRunData]
# callerIp:: The caller's IP for a run started by a Webhook Secret URL call; shown in run detail.
# contentDeletedAt:: When the run retention job cleared the run's content (ADR 0044): payload contents, run data, step
#   input, output and errors, the run error and the caller IP. Null while the content is kept.
# contentRetentionDays:: The retention period in days that applied when the content was cleared ("Details deleted after N days").
# testRunDefinition:: The draft definition a Test run executed (ADR 0049), so its graph can be drawn after the draft
#   changes. Null for live runs (they point at `workflowVersionId`) and for Test runs saved before it
#   existed. Configuration like a version's definition, not run content: retention leaves it alone.
#   [WorkflowExecutableDefinition]
class WorkflowRun < ApplicationRecord
  self.table_name = "WorkflowRun"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid2.call }
  attribute :isDryRun, default: false
  attribute :attempt, default: 0
  attribute :data, default: -> { JSON.parse("{}") }

  enum :status, { queued: "queued", running: "running", completed: "completed", failed: "failed", canceled: "canceled" },
       default: :queued, validate: true

  validates :triggerType, presence: true
  validates :idempotencyKey, uniqueness: { scope: :workflowId, allow_nil: true }
  validates :triggerPayload, exclusion: { in: [ nil ] }

  belongs_to :workflow, foreign_key: [ :workflowId, :workspaceId ], primary_key: [ :id, :workspaceId ],
                        inverse_of: :runs
  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :workflow_runs
  belongs_to :workflow_version, foreign_key: "workflowVersionId", optional: true, inverse_of: :runs
  belongs_to :response, foreign_key: "responseId", optional: true, inverse_of: :workflow_runs
  has_many :logs, class_name: "WorkflowRunLog", foreign_key: "runId", inverse_of: :run, dependent: :destroy
end
