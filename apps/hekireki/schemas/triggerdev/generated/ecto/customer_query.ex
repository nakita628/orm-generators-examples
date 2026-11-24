defmodule Trigger.CustomerQuery do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          query: String.t(),
          scope: atom(),
          stats: map() | list() | String.t() | number() | boolean(),
          title: String.t() | nil,
          source: atom(),
          filter_period: String.t() | nil,
          filter_from: DateTime.t() | nil,
          filter_to: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          organization: Trigger.Organization.t() | nil,
          project: Trigger.Project.t() | nil,
          environment: Trigger.RuntimeEnvironment.t() | nil,
          user: Trigger.User.t() | nil
        }

  schema "CustomerQuery" do
    field(:id, :string, primary_key: true)
    # The TSQL query text that was executed
    field(:query, :string)
    # The scope of the query (determines which tenant IDs were used)
    field(:scope, Ecto.Enum, values: [:ORGANIZATION, :PROJECT, :ENVIRONMENT])
    # Query execution statistics from ClickHouse
    field(:stats, Trigger.PrismaJson)
    # AI-generated title summarizing the query
    field(:title, :string)
    # Where the query originated from
    field(:source, Ecto.Enum, values: [:DASHBOARD, :API], default: :DASHBOARD)
    # Time filter settings used when the query was executed
    # Period like "7d", "24h", etc.
    field(:filter_period, :string, source: :filterPeriod)
    # Custom start date (ISO 8601)
    field(:filter_from, Trigger.PrismaDateTime, source: :filterFrom)
    # Custom end date (ISO 8601)
    field(:filter_to, Trigger.PrismaDateTime, source: :filterTo)
    field(:organization_id, :string, source: :organizationId)
    field(:project_id, :string, source: :projectId)
    field(:environment_id, :string, source: :environmentId)
    field(:user_id, :string, source: :userId)
    belongs_to(:organization, Trigger.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:project, Trigger.Project, foreign_key: :project_id, define_field: false, type: :string)
    belongs_to(:environment, Trigger.RuntimeEnvironment, foreign_key: :environment_id, define_field: false, type: :string)
    # Optional user who executed the query (null for API calls)
    belongs_to(:user, Trigger.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end