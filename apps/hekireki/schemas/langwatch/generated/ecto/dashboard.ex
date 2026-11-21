defmodule LangWatch.Dashboard do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          order: integer(),
          description: String.t() | nil,
          created_by_id: String.t() | nil,
          scope: atom(),
          organization_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          project: LangWatch.Project.t() | nil,
          graphs: [LangWatch.CustomGraph.t()]
        }

  schema "Dashboard" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    # Legacy board sort position. The Dashboards v2 app orders boards by a
    # member's favourites and never writes this; kept because the legacy
    # Analytics reports section and the REST reorder endpoint still write it,
    # and a rolling deploy's older image still reads it.
    field(:order, :integer, default: 0)
    # Free text under the board's name; null until a member writes one.
    # Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    field(:description, :string)
    # The member who created it; null for boards that predate the column.
    # Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    field(:created_by_id, :string, source: :createdById)
    # Who sees it. Only the author changes it; a board with no author stays PROJECT.
    # Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    field(:scope, Ecto.Enum, values: [:PRIVATE, :PROJECT, :ORGANIZATION], default: :PROJECT)
    # The owning project's organization, stamped when the board is first set to
    # ORGANIZATION; null before. No relation, per the database rule.
    # Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    field(:organization_id, :string, source: :organizationId)
    # The project that owns the board; only there can it be edited.
    field(:project_id, :string, source: :projectId)
    belongs_to(:project, LangWatch.Project, foreign_key: :project_id, define_field: false, type: :string)
    has_many(:graphs, LangWatch.CustomGraph, foreign_key: :dashboard_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end