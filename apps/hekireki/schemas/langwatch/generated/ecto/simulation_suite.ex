defmodule LangWatch.SimulationSuite do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          slug: String.t(),
          description: String.t() | nil,
          scenario_ids: [String.t()],
          targets: map() | list() | String.t() | number() | boolean(),
          repeat_count: integer(),
          labels: [String.t()],
          simulator_model: String.t() | nil,
          judge_model: String.t() | nil,
          kind: String.t(),
          scope: map() | list() | String.t() | number() | boolean() | nil,
          fields: map() | list() | String.t() | number() | boolean() | nil,
          evaluators: map() | list() | String.t() | number() | boolean() | nil,
          archived_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          project: LangWatch.Project.t() | nil,
          test_suite_scenarios: [LangWatch.Scenario.t()]
        }

  schema "SimulationSuite" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:slug, :string)
    field(:description, :string)
    field(:scenario_ids, {:array, :string}, source: :scenarioIds)
    field(:targets, LangWatch.PrismaJson)
    field(:repeat_count, :integer, default: 1, source: :repeatCount)
    field(:labels, {:array, :string})
    field(:simulator_model, :string, source: :simulatorModel)
    field(:judge_model, :string, source: :judgeModel)
    field(:kind, :string, default: "run_plan")
    field(:scope, LangWatch.PrismaJson)
    field(:fields, LangWatch.PrismaJson)
    field(:evaluators, LangWatch.PrismaJson)
    field(:archived_at, LangWatch.PrismaDateTime, source: :archivedAt)
    field(:project_id, :string, source: :projectId)
    belongs_to(:project, LangWatch.Project, foreign_key: :project_id, define_field: false, type: :string)
    has_many(:test_suite_scenarios, LangWatch.Scenario, foreign_key: :test_suite_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end