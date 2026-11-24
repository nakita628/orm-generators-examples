defmodule Plunk.WorkflowStep do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Plunk.PrismaDateTime, autogenerate: {Plunk.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          type: atom(),
          name: String.t(),
          position: map() | list() | String.t() | number() | boolean(),
          config: map() | list() | String.t() | number() | boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          workflow: Plunk.Workflow.t() | nil,
          template: Plunk.Template.t() | nil,
          outgoing_transitions: [Plunk.WorkflowTransition.t()],
          incoming_transitions: [Plunk.WorkflowTransition.t()],
          executions: [Plunk.WorkflowStepExecution.t()],
          current_workflow_executions: [Plunk.WorkflowExecution.t()]
        }

  schema "workflow_steps" do
    field(:type, Ecto.Enum, values: [:TRIGGER, :SEND_EMAIL, :DELAY, :WAIT_FOR_EVENT, :CONDITION, :EXIT, :WEBHOOK, :UPDATE_CONTACT])
    field(:name, :string)
    field(:position, Plunk.PrismaJson)
    field(:config, Plunk.PrismaJson)
    field(:workflow_id, :string, source: :workflowId)
    field(:template_id, :string, source: :templateId)
    belongs_to(:workflow, Plunk.Workflow, foreign_key: :workflow_id, define_field: false)
    belongs_to(:template, Plunk.Template, foreign_key: :template_id, define_field: false)
    has_many(:outgoing_transitions, Plunk.WorkflowTransition, foreign_key: :from_step_id)
    has_many(:incoming_transitions, Plunk.WorkflowTransition, foreign_key: :to_step_id)
    has_many(:executions, Plunk.WorkflowStepExecution, foreign_key: :step_id)
    has_many(:current_workflow_executions, Plunk.WorkflowExecution, foreign_key: :current_step_id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end