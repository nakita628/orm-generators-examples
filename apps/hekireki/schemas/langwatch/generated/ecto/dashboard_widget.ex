defmodule LangWatch.DashboardWidget do
  @moduledoc """
  A Dashboards v2 widget: a srcdoc definition placed on a DashboardsV2 board, or
  unplaced. Widgets made before the split were copied from CustomGraph with their
  ids. No relation, per the database rule.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          project_id: String.t(),
          dashboard_id: String.t() | nil,
          name: String.t(),
          definition: map(),
          grid_column: integer(),
          grid_row: integer(),
          col_span: integer(),
          row_span: integer(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "DashboardWidget" do
    field(:id, :string, primary_key: true)
    field(:project_id, :string, source: :projectId)
    # The DashboardsV2 board it sits on; null while unplaced.
    field(:dashboard_id, :string, source: :dashboardId)
    field(:name, :string)
    # The widget's srcdoc definition.
    field(:definition, :map)
    field(:grid_column, :integer, default: 0, source: :gridColumn)
    field(:grid_row, :integer, default: 0, source: :gridRow)
    field(:col_span, :integer, default: 1, source: :colSpan)
    field(:row_span, :integer, default: 1, source: :rowSpan)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end