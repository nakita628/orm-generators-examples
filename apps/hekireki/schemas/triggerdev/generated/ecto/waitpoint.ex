defmodule Trigger.Waitpoint do
  @moduledoc """
  A Waitpoint blocks a run from continuing until it's completed
  If there's a waitpoint blocking a run, it shouldn't be in the queue
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          friendly_id: String.t(),
          type: atom(),
          status: atom(),
          completed_at: DateTime.t() | nil,
          idempotency_key: String.t(),
          user_provided_idempotency_key: boolean(),
          idempotency_key_expires_at: DateTime.t() | nil,
          inactive_idempotency_key: String.t() | nil,
          completed_after: DateTime.t() | nil,
          output: String.t() | nil,
          output_type: String.t(),
          output_is_error: boolean(),
          tags: [String.t()],
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          completed_by_task_run: Trigger.TaskRun.t() | nil,
          completed_by_batch: Trigger.BatchTaskRun.t() | nil,
          project: Trigger.Project.t() | nil,
          environment: Trigger.RuntimeEnvironment.t() | nil,
          blocking_task_runs: [Trigger.TaskRunWaitpoint.t()],
          connected_runs: [Trigger.TaskRun.t()],
          completed_execution_snapshots: [Trigger.TaskRunExecutionSnapshot.t()]
        }

  schema "Waitpoint" do
    field(:id, :string, primary_key: true)
    field(:friendly_id, :string, source: :friendlyId)
    field(:type, Ecto.Enum, values: [:RUN, :DATETIME, :MANUAL, :BATCH])
    field(:status, Ecto.Enum, values: [:PENDING, :COMPLETED], default: :PENDING)
    field(:completed_at, Trigger.PrismaDateTime, source: :completedAt)
    # If it's an Event type waitpoint, this is the event. It can also be provided for the DATETIME type
    field(:idempotency_key, :string, source: :idempotencyKey)
    # If this is true then we can show it in the dashboard/return it from the SDK
    field(:user_provided_idempotency_key, :boolean, source: :userProvidedIdempotencyKey)
    # If there's a user provided idempotency key, this is the time it expires at
    field(:idempotency_key_expires_at, Trigger.PrismaDateTime, source: :idempotencyKeyExpiresAt)
    # If an idempotencyKey is no longer active, we store it here and generate a new one for the idempotencyKey field.
    # Clearing an idempotencyKey is useful for debounce or cancelling child runs.
    # This is a workaround because Prisma doesn't support partial indexes.
    field(:inactive_idempotency_key, :string, source: :inactiveIdempotencyKey)
    # If it's a DATETIME type waitpoint, this is the date.
    # If it's a MANUAL waitpoint, this can be set as the `timeout`.
    field(:completed_after, Trigger.PrismaDateTime, source: :completedAfter)
    # When completed, an output can be stored here
    field(:output, :string)
    field(:output_type, :string, default: "application/json", source: :outputType)
    field(:output_is_error, :boolean, default: false, source: :outputIsError)
    # Denormized column that holds the raw tags
    # Denormalized column that holds the raw tags
    field(:tags, {:array, :string})
    # If it's a RUN type waitpoint, this is the associated run
    field(:completed_by_task_run_id, :string, source: :completedByTaskRunId)
    # If it's a BATCH type waitpoint, this is the associated batch
    field(:completed_by_batch_id, :string, source: :completedByBatchId)
    field(:project_id, :string, source: :projectId)
    field(:environment_id, :string, source: :environmentId)
    belongs_to(:completed_by_task_run, Trigger.TaskRun, foreign_key: :completed_by_task_run_id, define_field: false, type: :string)
    belongs_to(:completed_by_batch, Trigger.BatchTaskRun, foreign_key: :completed_by_batch_id, define_field: false, type: :string)
    belongs_to(:project, Trigger.Project, foreign_key: :project_id, define_field: false, type: :string)
    belongs_to(:environment, Trigger.RuntimeEnvironment, foreign_key: :environment_id, define_field: false, type: :string)
    # The runs this waitpoint is blocking
    has_many(:blocking_task_runs, Trigger.TaskRunWaitpoint, foreign_key: :waitpoint_id, references: :id)
    # All runs that have ever been blocked by this waitpoint, used for display purposes
    many_to_many(:connected_runs, Trigger.TaskRun, join_through: "_WaitpointRunConnections", join_keys: [B: :id, A: :id])
    # When a waitpoint is complete
    many_to_many(:completed_execution_snapshots, Trigger.TaskRunExecutionSnapshot, join_through: "_completedWaitpoints", join_keys: [B: :id, A: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end