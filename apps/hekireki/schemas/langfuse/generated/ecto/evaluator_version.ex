defmodule Langfuse.EvaluatorVersion do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Langfuse.PrismaDateTime, autogenerate: {Langfuse.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          version: integer(),
          prompt: String.t() | nil,
          prompt_messages: map() | list() | String.t() | number() | boolean() | nil,
          partner: String.t() | nil,
          model: String.t() | nil,
          provider: String.t() | nil,
          model_params: map() | list() | String.t() | number() | boolean() | nil,
          vars: [String.t()],
          variable_mapping: map() | list() | String.t() | number() | boolean() | nil,
          output_definition: map() | list() | String.t() | number() | boolean() | nil,
          source_code: String.t() | nil,
          source_code_language: atom() | nil,
          questions: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          evaluator: Langfuse.Evaluator.t() | nil,
          created_by_user: Langfuse.User.t() | nil
        }

  schema "evaluator_versions" do
    field(:id, :string, primary_key: true)
    field(:version, :integer)
    field(:prompt, :string)
    field(:prompt_messages, Langfuse.PrismaJson)
    field(:partner, :string)
    field(:model, :string)
    field(:provider, :string)
    field(:model_params, Langfuse.PrismaJson)
    field(:vars, {:array, :string})
    field(:variable_mapping, Langfuse.PrismaJson)
    field(:output_definition, Langfuse.PrismaJson)
    field(:source_code, :string)
    field(:source_code_language, Ecto.Enum, values: [:PYTHON, :TYPESCRIPT])
    field(:questions, Langfuse.PrismaJson)
    belongs_to(:evaluator, Langfuse.Evaluator, foreign_key: :evaluator_id, type: :string)
    belongs_to(:created_by_user, Langfuse.User, foreign_key: :created_by_user_id, type: :string)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end