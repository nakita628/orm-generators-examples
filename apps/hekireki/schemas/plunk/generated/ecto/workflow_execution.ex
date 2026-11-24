defmodule Plunk.WorkflowExecution do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Plunk.PrismaDateTime, autogenerate: {Plunk.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          status: atom(),
          exit_reason: String.t() | nil,
          context: map() | list() | String.t() | number() | boolean() | nil,
          started_at: DateTime.t(),
          completed_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          workflow: Plunk.Workflow.t() | nil,
          contact: Plunk.Contact.t() | nil,
          current_step: Plunk.WorkflowStep.t() | nil,
          step_executions: [Plunk.WorkflowStepExecution.t()],
          emails: [Plunk.Email.t()]
        }

  schema "workflow_executions" do
    field(:status, Ecto.Enum, values: [:RUNNING, :WAITING, :COMPLETED, :EXITED, :FAILED, :CANCELLED], default: :RUNNING)
    field(:exit_reason, :string, source: :exitReason)
    field(:context, Plunk.PrismaJson)
    field(:started_at, Plunk.PrismaDateTime, autogenerate: true, source: :startedAt)
    field(:completed_at, Plunk.PrismaDateTime, source: :completedAt)
    field(:workflow_id, :string, source: :workflowId)
    field(:contact_id, :string, source: :contactId)
    field(:current_step_id, :string, source: :currentStepId)
    belongs_to(:workflow, Plunk.Workflow, foreign_key: :workflow_id, define_field: false)
    belongs_to(:contact, Plunk.Contact, foreign_key: :contact_id, define_field: false)
    belongs_to(:current_step, Plunk.WorkflowStep, foreign_key: :current_step_id, define_field: false)
    has_many(:step_executions, Plunk.WorkflowStepExecution, foreign_key: :execution_id)
    has_many(:emails, Plunk.Email, foreign_key: :workflow_execution_id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end