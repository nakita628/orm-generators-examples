defmodule Plunk.WorkflowTransition do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Plunk.PrismaDateTime, autogenerate: {Plunk.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          condition: map() | list() | String.t() | number() | boolean() | nil,
          priority: integer(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          from_step: Plunk.WorkflowStep.t() | nil,
          to_step: Plunk.WorkflowStep.t() | nil
        }

  schema "workflow_transitions" do
    field(:condition, Plunk.PrismaJson)
    field(:priority, :integer, default: 0)
    field(:from_step_id, :string, source: :fromStepId)
    field(:to_step_id, :string, source: :toStepId)
    belongs_to(:from_step, Plunk.WorkflowStep, foreign_key: :from_step_id, define_field: false)
    belongs_to(:to_step, Plunk.WorkflowStep, foreign_key: :to_step_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end