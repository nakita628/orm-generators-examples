# A Waitpoint blocks a run from continuing until it's completed
# If there's a waitpoint blocking a run, it shouldn't be in the queue
#
# idempotencyKey:: If it's an Event type waitpoint, this is the event. It can also be provided for the DATETIME type
# userProvidedIdempotencyKey:: If this is true then we can show it in the dashboard/return it from the SDK
# idempotencyKeyExpiresAt:: If there's a user provided idempotency key, this is the time it expires at
# inactiveIdempotencyKey:: If an idempotencyKey is no longer active, we store it here and generate a new one for the idempotencyKey field.
#   Clearing an idempotencyKey is useful for debounce or cancelling child runs.
#   This is a workaround because Prisma doesn't support partial indexes.
# completedByTaskRunId:: If it's a RUN type waitpoint, this is the associated run
# completedAfter:: If it's a DATETIME type waitpoint, this is the date.
#   If it's a MANUAL waitpoint, this can be set as the `timeout`.
# completedByBatchId:: If it's a BATCH type waitpoint, this is the associated batch
# blocking_task_runs:: The runs this waitpoint is blocking
# connected_runs:: All runs that have ever been blocked by this waitpoint, used for display purposes
# completed_execution_snapshots:: When a waitpoint is complete
# output:: When completed, an output can be stored here
# tags:: Denormized column that holds the raw tags
#   Denormalized column that holds the raw tags
class Waitpoint < ApplicationRecord
  self.table_name = "Waitpoint"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :outputType, default: "application/json"
  attribute :outputIsError, default: false
  attribute :tags, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { run: "RUN", datetime: "DATETIME", manual: "MANUAL", batch: "BATCH" }, validate: true
  enum :status, { pending: "PENDING", completed: "COMPLETED" }, default: :pending, validate: true

  validates :friendlyId, presence: true, uniqueness: true
  validates :idempotencyKey, presence: true, uniqueness: { scope: :environmentId }
  validates :userProvidedIdempotencyKey, inclusion: { in: [ true, false ] }
  validates :completedByTaskRunId, uniqueness: { allow_nil: true }

  belongs_to :completed_by_task_run, class_name: "TaskRun", foreign_key: "completedByTaskRunId", optional: true,
                                     inverse_of: :associated_waitpoint
  belongs_to :completed_by_batch, class_name: "BatchTaskRun", foreign_key: "completedByBatchId", optional: true,
                                  inverse_of: :waitpoints
  belongs_to :project, foreign_key: "projectId", inverse_of: :waitpoints
  belongs_to :environment, class_name: "RuntimeEnvironment", foreign_key: "environmentId", inverse_of: :waitpoints
  has_many :blocking_task_runs, class_name: "TaskRunWaitpoint", foreign_key: "waitpointId", inverse_of: :waitpoint,
                                dependent: :restrict_with_error
  has_and_belongs_to_many :connected_runs, class_name: "TaskRun", join_table: "_WaitpointRunConnections",
                                           foreign_key: "B", association_foreign_key: "A"
  has_and_belongs_to_many :completed_execution_snapshots, class_name: "TaskRunExecutionSnapshot",
                                                          join_table: "_completedWaitpoints", foreign_key: "B",
                                                          association_foreign_key: "A"
end
