defmodule Formbricks.DashboardWidget do
  @moduledoc """
  Represents a chart widget on a dashboard.
  Widgets are positioned using a grid layout system.

  @property id - Unique identifier for the widget
  @property dashboard - The dashboard this widget belongs to
  @property chart - The chart displayed in this widget (chart name is used for display)
  @property layout - Grid layout configuration (x, y, width, height)
  @property order - Display order within the dashboard
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          layout: map() | list() | String.t() | number() | boolean(),
          order: integer(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          dashboard: Formbricks.Dashboard.t() | nil,
          chart: Formbricks.Chart.t() | nil
        }

  schema "DashboardWidget" do
    field(:id, :string, primary_key: true)
    # [WidgetLayout] - Grid layout: { x, y, w, h }
    field(:layout, Formbricks.PrismaJson, default: %{"x" => 0, "y" => 0, "w" => 4, "h" => 3})
    field(:order, :integer, default: 0)
    field(:dashboard_id, :string, source: :dashboardId)
    field(:chart_id, :string, source: :chartId)
    belongs_to(:dashboard, Formbricks.Dashboard, foreign_key: :dashboard_id, define_field: false, type: :string)
    belongs_to(:chart, Formbricks.Chart, foreign_key: :chart_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end