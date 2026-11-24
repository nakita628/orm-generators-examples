defmodule Plunk.WorkflowStepExecution do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Plunk.PrismaDateTime, autogenerate: {Plunk.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          status: atom(),
          scheduled_for: DateTime.t() | nil,
          execute_after: DateTime.t() | nil,
          output: map() | list() | String.t() | number() | boolean() | nil,
          error: String.t() | nil,
          started_at: DateTime.t() | nil,
          completed_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          execution: Plunk.WorkflowExecution.t() | nil,
          step: Plunk.WorkflowStep.t() | nil,
          emails: [Plunk.Email.t()]
        }

  schema "workflow_step_executions" do
    field(:status, Ecto.Enum, values: [:PENDING, :SCHEDULED, :WAITING, :RUNNING, :COMPLETED, :SKIPPED, :FAILED], default: :PENDING)
    field(:scheduled_for, Plunk.PrismaDateTime, source: :scheduledFor)
    field(:execute_after, Plunk.PrismaDateTime, source: :executeAfter)
    field(:output, Plunk.PrismaJson)
    field(:error, :string)
    field(:started_at, Plunk.PrismaDateTime, source: :startedAt)
    field(:completed_at, Plunk.PrismaDateTime, source: :completedAt)
    field(:execution_id, :string, source: :executionId)
    field(:step_id, :string, source: :stepId)
    belongs_to(:execution, Plunk.WorkflowExecution, foreign_key: :execution_id, define_field: false)
    belongs_to(:step, Plunk.WorkflowStep, foreign_key: :step_id, define_field: false)
    has_many(:emails, Plunk.Email, foreign_key: :workflow_step_execution_id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end