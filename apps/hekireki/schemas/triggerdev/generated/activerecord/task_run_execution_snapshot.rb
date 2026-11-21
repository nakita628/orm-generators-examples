# Used by the RunEngine during TaskRun execution
# It has the required information to transactionally progress a run through states,
# and prevent side effects like heartbeats failing a run that has progressed.
# It is optimised for performance and is designed to be cleared at some point,
# so there are no cascading relationships to other models.
#
# engine:: This should always be 2+ (V1 didn't use the run engine or snapshots)
# executionStatus:: The execution status
# description:: For debugging
# isValid:: We store invalid snapshots as a record of the run state when we tried to move
# previousSnapshotId:: The previous snapshot ID
# runId:: Run
# attemptNumber:: This is the current run attempt number. Users can define how many attempts they want for a run.
# environmentId:: Environment
# completed_waitpoints:: Waitpoints that have been completed for this execution
# completedWaitpointOrder:: An array of waitpoint IDs in the correct order, used for batches
# checkpointId:: Checkpoint
# workerId:: Worker
# metadata:: Metadata used by various systems in the run engine
class TaskRunExecutionSnapshot < ApplicationRecord
  self.table_name = "TaskRunExecutionSnapshot"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :isValid, default: true
  attribute :completedWaitpointOrder, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :engine, { v1: "V1", v2: "V2" }, default: :v2, validate: true
  enum :executionStatus, { run_created: "RUN_CREATED", delayed: "DELAYED", queued: "QUEUED", queued_executing: "QUEUED_EXECUTING", pending_executing: "PENDING_EXECUTING", executing: "EXECUTING", executing_with_waitpoints: "EXECUTING_WITH_WAITPOINTS", suspended: "SUSPENDED", pending_cancel: "PENDING_CANCEL", finished: "FINISHED" },
       validate: true
  enum :runStatus, { delayed: "DELAYED", pending: "PENDING", pending_version: "PENDING_VERSION", waiting_for_deploy: "WAITING_FOR_DEPLOY", dequeued: "DEQUEUED", executing: "EXECUTING", waiting_to_resume: "WAITING_TO_RESUME", retrying_after_failure: "RETRYING_AFTER_FAILURE", paused: "PAUSED", canceled: "CANCELED", interrupted: "INTERRUPTED", completed_successfully: "COMPLETED_SUCCESSFULLY", completed_with_errors: "COMPLETED_WITH_ERRORS", system_failure: "SYSTEM_FAILURE", crashed: "CRASHED", expired: "EXPIRED", timed_out: "TIMED_OUT" },
       prefix: true, validate: true
  enum :environmentType, { production: "PRODUCTION", staging: "STAGING", development: "DEVELOPMENT", preview: "PREVIEW" },
       validate: true

  validates :description, presence: true
  validates :environmentId, presence: true
  validates :projectId, presence: true
  validates :organizationId, presence: true

  belongs_to :run, class_name: "TaskRun", foreign_key: "runId", inverse_of: :execution_snapshots
  belongs_to :checkpoint, class_name: "TaskRunCheckpoint", foreign_key: "checkpointId", optional: true,
                          inverse_of: :execution_snapshot
  has_and_belongs_to_many :completed_waitpoints, class_name: "Waitpoint", join_table: "_completedWaitpoints",
                                                 foreign_key: "A", association_foreign_key: "B"
end
