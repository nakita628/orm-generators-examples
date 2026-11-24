defmodule Umami.HeatmapEvent do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Umami.PrismaDateTime, autogenerate: {Umami.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          session_id: Ecto.UUID.t(),
          visit_id: Ecto.UUID.t(),
          url_path: String.t(),
          event_type: integer(),
          x: integer() | nil,
          y: integer() | nil,
          page_x: integer() | nil,
          page_y: integer() | nil,
          page_w: integer() | nil,
          viewport_w: integer() | nil,
          viewport_h: integer() | nil,
          page_h: integer() | nil,
          scroll_pct: integer() | nil,
          inserted_at: DateTime.t(),
          website: Umami.Website.t() | nil
        }

  schema "heatmap_event" do
    field(:id, Ecto.UUID, primary_key: true, source: :heatmap_event_id)
    field(:session_id, Ecto.UUID)
    field(:visit_id, Ecto.UUID)
    field(:url_path, :string)
    field(:event_type, :integer)
    field(:x, :integer)
    field(:y, :integer)
    field(:page_x, :integer)
    field(:page_y, :integer)
    field(:page_w, :integer)
    field(:viewport_w, :integer)
    field(:viewport_h, :integer)
    field(:page_h, :integer)
    field(:scroll_pct, :integer)
    belongs_to(:website, Umami.Website, foreign_key: :website_id, type: Ecto.UUID)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end