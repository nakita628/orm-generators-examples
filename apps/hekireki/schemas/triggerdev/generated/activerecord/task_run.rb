# idempotencyKeyOptions:: Stores the user-provided key and scope: { key: string, scope: "run" | "attempt" | "global" }
# debounce:: Debounce options: { key: string, delay: string, createdAt: Date }
# masterQueue:: The main queue that this run is part of
# region:: User-facing geo region, stamped at trigger; workerQueue is where it actually ran.
# secondaryMasterQueue:: @deprecated
# attemptNumber:: From engine v2+ this will be defined after a run has been dequeued (starting at 1)
# runTags:: Denormized column that holds the raw tags
# taskVersion:: Denormalized version of the background worker task
# startedAt:: startedAt marks the point at which a run is dequeued from MarQS
# executedAt:: executedAt is set when the first attempt is about to execute
# priorityMs:: The "priority" of the run. This is just a negative offset in ms for the queue timestamp
#   E.g. a value of 60_000 would put the run into the queue 60s ago.
# gates:: Gates for this run: other queues it must also hold a concurrency slot in while
#   executing. Shape: [{ queue: string, concurrencyKey?: string }]
# oneTimeUseToken:: optional token that can be used to authenticate the task run
# associated_waitpoint:: When this run is finished, the waitpoint will be marked as completed
# blocked_by_waitpoints:: If there are any blocked waitpoints, the run won't be executed
# connected_waitpoints:: All waitpoints that blocked this run at some point, used for display purposes
# taskEventStore:: Where the logs are stored
# root_task_run:: This represents the original task that that was triggered outside of a Trigger.dev task
# descendant_runs:: The root run will have a list of all the descendant runs, children, grand children, etc.
# parent_task_run:: The immediate parent run of this task run
# child_runs:: The immediate child runs of this task run
# parent_task_run_attempt:: The immediate parent attempt of this task run
# batch:: The batch run that this task run is a part of
# resumeParentOnCompletion:: whether or not the task run was created because of a triggerAndWait for batchTriggerAndWait
# depth:: The depth of this task run in the task run hierarchy
# parentSpanId:: The span ID of the "trigger" span in the parent task run
# runChainState:: Holds the state of the run chain for deadlock detection
# seedMetadata:: seed run metadata
# metadata:: Run metadata
# annotations:: Structured annotations: triggerSource, triggerAction, rootTriggerSource, rootScheduleId
# isWarmStart:: Whether the latest attempt was a warm start. Null until first attempt starts.
# output:: Run output
# error:: Run error
# planType:: Organization's billing plan type (cached for fallback when billing API fails)
# realtimeStreamsVersion:: The version of the realtime streams implementation used by the run
# realtimeStreams:: Store the stream keys that are being used by the run
# streamBasinName:: S2 basin where this run's realtime streams live. Stamped at create
#   time from `Organization.streamBasinName` so reads can resolve the
#   basin without joining org. Null when the org has no per-org basin
#   (OSS, or pre-backfill); reads fall back to the global basin.
class TaskRun < ApplicationRecord
  self.table_name = "TaskRun"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :number, default: 0
  attribute :isTest, default: false
  attribute :payloadType, default: "application/json"
  attribute :masterQueue, default: "main"
  attribute :runTags, default: -> { [] }
  attribute :usageDurationMs, default: 0
  attribute :costInCents, default: 0
  attribute :baseCostInCents, default: 0
  attribute :priorityMs, default: 0
  attribute :taskEventStore, default: "taskEvent"
  attribute :bulkActionGroupIds, default: -> { [] }
  attribute :resumeParentOnCompletion, default: false
  attribute :depth, default: 0
  attribute :seedMetadataType, default: "application/json"
  attribute :metadataType, default: "application/json"
  attribute :metadataVersion, default: 1
  attribute :outputType, default: "application/json"
  attribute :realtimeStreamsVersion, default: "v1"
  attribute :realtimeStreams, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :engine, { v1: "V1", v2: "V2" }, default: :v1, validate: true
  enum :status, { delayed: "DELAYED", pending: "PENDING", pending_version: "PENDING_VERSION", waiting_for_deploy: "WAITING_FOR_DEPLOY", dequeued: "DEQUEUED", executing: "EXECUTING", waiting_to_resume: "WAITING_TO_RESUME", retrying_after_failure: "RETRYING_AFTER_FAILURE", paused: "PAUSED", canceled: "CANCELED", interrupted: "INTERRUPTED", completed_successfully: "COMPLETED_SUCCESSFULLY", completed_with_errors: "COMPLETED_WITH_ERRORS", system_failure: "SYSTEM_FAILURE", crashed: "CRASHED", expired: "EXPIRED", timed_out: "TIMED_OUT" },
       default: :pending, validate: true
  enum :environmentType, { production: "PRODUCTION", staging: "STAGING", development: "DEVELOPMENT", preview: "PREVIEW" },
       validate: { allow_nil: true }

  validates :friendlyId, presence: true, uniqueness: true
  validates :taskIdentifier, presence: true,
                             uniqueness: { scope: [ :runtimeEnvironmentId, :idempotencyKey ], conditions: -> { where.not(idempotencyKey: nil) } }
  validates :payload, presence: true
  validates :traceId, presence: true
  validates :spanId, presence: true
  validates :queue, presence: true
  validates :oneTimeUseToken, uniqueness: { allow_nil: true }
  validates :oneTimeUseToken, uniqueness: { allow_nil: true }

  belongs_to :runtime_environment, foreign_key: "runtimeEnvironmentId", inverse_of: :task_runs
  belongs_to :project, foreign_key: "projectId", inverse_of: :task_runs
  belongs_to :locked_by, class_name: "BackgroundWorkerTask", foreign_key: "lockedById", optional: true,
                         inverse_of: :runs
  belongs_to :locked_to_version, class_name: "BackgroundWorker", foreign_key: "lockedToVersionId", optional: true,
                                 inverse_of: :locked_runs
  belongs_to :root_task_run, class_name: "TaskRun", foreign_key: "rootTaskRunId", optional: true,
                             inverse_of: :descendant_runs
  belongs_to :parent_task_run, class_name: "TaskRun", foreign_key: "parentTaskRunId", optional: true,
                               inverse_of: :child_runs
  belongs_to :parent_task_run_attempt, class_name: "TaskRunAttempt", foreign_key: "parentTaskRunAttemptId",
                                       optional: true, inverse_of: :child_runs
  belongs_to :batch, class_name: "BatchTaskRun", foreign_key: "batchId", optional: true, inverse_of: :runs
  has_one :associated_waitpoint, class_name: "Waitpoint", foreign_key: "completedByTaskRunId",
                                 inverse_of: :completed_by_task_run, dependent: :nullify
  has_one :dependency, class_name: "TaskRunDependency", foreign_key: "taskRunId", inverse_of: :task_run,
                       dependent: :destroy
  has_many :attempts, class_name: "TaskRunAttempt", foreign_key: "taskRunId", inverse_of: :task_run, dependent: :destroy
  has_many :checkpoints, foreign_key: "runId", inverse_of: :run, dependent: :destroy
  has_many :blocked_by_waitpoints, class_name: "TaskRunWaitpoint", foreign_key: "taskRunId", inverse_of: :task_run,
                                   dependent: :restrict_with_error
  has_many :batch_items, class_name: "BatchTaskRunItem", foreign_key: "taskRunId", inverse_of: :task_run,
                         dependent: :destroy
  has_many :checkpoint_restore_event, foreign_key: "runId", inverse_of: :run, dependent: :destroy
  has_many :execution_snapshots, class_name: "TaskRunExecutionSnapshot", foreign_key: "runId", inverse_of: :run,
                                 dependent: :restrict_with_error
  has_many :source_bulk_action_items, class_name: "BulkActionItem", foreign_key: "sourceRunId", inverse_of: :source_run,
                                      dependent: :destroy
  has_many :destination_bulk_action_items, class_name: "BulkActionItem", foreign_key: "destinationRunId",
                                           inverse_of: :destination_run, dependent: :destroy
  has_many :descendant_runs, class_name: "TaskRun", foreign_key: "rootTaskRunId", inverse_of: :root_task_run,
                             dependent: :nullify
  has_many :child_runs, class_name: "TaskRun", foreign_key: "parentTaskRunId", inverse_of: :parent_task_run,
                        dependent: :nullify
  has_many :playground_conversations, foreign_key: "runId", inverse_of: :run, dependent: :nullify
  has_and_belongs_to_many :connected_waitpoints, class_name: "Waitpoint", join_table: "_WaitpointRunConnections",
                                                 foreign_key: "A", association_foreign_key: "B"
end
