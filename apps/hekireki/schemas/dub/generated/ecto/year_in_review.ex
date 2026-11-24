defmodule Dub.YearInReview do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Dub.PrismaDateTime, autogenerate: {Dub.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          year: integer(),
          total_links: integer(),
          total_clicks: integer(),
          top_links: map() | list() | String.t() | number() | boolean(),
          top_countries: map() | list() | String.t() | number() | boolean(),
          sent_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          workspace: Dub.Project.t() | nil
        }

  schema "YearInReview" do
    field(:id, :string, primary_key: true)
    field(:year, :integer)
    field(:total_links, :integer, source: :totalLinks)
    field(:total_clicks, :integer, source: :totalClicks)
    field(:top_links, Dub.PrismaJson, source: :topLinks)
    field(:top_countries, Dub.PrismaJson, source: :topCountries)
    field(:sent_at, Dub.PrismaDateTime, source: :sentAt)
    field(:workspace_id, :string, source: :workspaceId)
    belongs_to(:workspace, Dub.Project, foreign_key: :workspace_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end