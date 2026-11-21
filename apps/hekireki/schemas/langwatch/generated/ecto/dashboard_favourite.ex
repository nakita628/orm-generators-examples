defmodule LangWatch.DashboardFavourite do
  @moduledoc """
  A member's star, in their own order within a project: a board or a From
  LangWatch template. Personal: one member's stars never show to another. No
  foreign key (per the database rule): the dashboard service removes a board's
  favourites when it deletes the board, so a star never outlives its dashboard.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          user_id: String.t(),
          dashboard_id: String.t() | nil,
          template_id: String.t() | nil,
          project_id: String.t(),
          position: integer(),
          inserted_at: DateTime.t()
        }

  schema "DashboardFavourite" do
    field(:id, :string, primary_key: true)
    field(:user_id, :string, source: :userId)
    # A DashboardsV2 board's id (copied boards kept their Dashboard id); null when
    # the star is a template.
    field(:dashboard_id, :string, source: :dashboardId)
    # The starred From LangWatch template id; null when the star is a board.
    field(:template_id, :string, source: :templateId)
    field(:project_id, :string, source: :projectId)
    # The member's drag order among their own stars in the project.
    field(:position, :integer)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end