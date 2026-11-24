defmodule Trigger.EnvironmentVariableIgnore do
  @moduledoc """
  A reason the environment variables "Needs attention" page shows a key in one environment, which
  the project's team chose to ignore. An ignored EMPTY goes when the value is next saved.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          reason: atom(),
          inserted_at: DateTime.t(),
          project: Trigger.Project.t() | nil,
          environment: Trigger.RuntimeEnvironment.t() | nil,
          variable: Trigger.EnvironmentVariable.t() | nil,
          created_by: Trigger.User.t() | nil
        }

  schema "EnvironmentVariableIgnore" do
    field(:id, :string, primary_key: true)
    field(:reason, Ecto.Enum, values: [:EMPTY, :MISSING, :OIDC])
    field(:project_id, :string, source: :projectId)
    field(:environment_id, :string, source: :environmentId)
    field(:variable_id, :string, source: :variableId)
    field(:created_by_id, :string, source: :createdById)
    belongs_to(:project, Trigger.Project, foreign_key: :project_id, define_field: false, type: :string)
    belongs_to(:environment, Trigger.RuntimeEnvironment, foreign_key: :environment_id, define_field: false, type: :string)
    # For MISSING, a variable of the project with no value in this environment.
    belongs_to(:variable, Trigger.EnvironmentVariable, foreign_key: :variable_id, define_field: false, type: :string)
    belongs_to(:created_by, Trigger.User, foreign_key: :created_by_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end