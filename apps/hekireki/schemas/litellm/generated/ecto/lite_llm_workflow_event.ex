defmodule LiteLLM.LiteLLMWorkflowEvent do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :event_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          event_type: String.t(),
          step_name: String.t(),
          sequence_number: integer(),
          data: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          run: LiteLLM.LiteLLMWorkflowRun.t() | nil
        }

  schema "LiteLLM_WorkflowEvent" do
    field(:event_type, :string)
    field(:step_name, :string)
    field(:sequence_number, :integer)
    field(:data, LiteLLM.PrismaJson)
    belongs_to(:run, LiteLLM.LiteLLMWorkflowRun, foreign_key: :run_id)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end