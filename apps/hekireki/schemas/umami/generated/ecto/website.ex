defmodule Umami.Website do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Umami.PrismaDateTime, autogenerate: {Umami.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          name: String.t(),
          domain: String.t() | nil,
          reset_at: DateTime.t() | nil,
          deleted_at: DateTime.t() | nil,
          recorder_enabled: boolean(),
          replay_config: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t() | nil,
          updated_at: DateTime.t() | nil,
          user: Umami.User.t() | nil,
          create_user: Umami.User.t() | nil,
          team: Umami.Team.t() | nil,
          annotations: [Umami.Annotation.t()],
          event_data: [Umami.EventData.t()],
          reports: [Umami.Report.t()],
          revenue: [Umami.Revenue.t()],
          segments: [Umami.Segment.t()],
          session_data: [Umami.SessionData.t()],
          session_replays: [Umami.SessionReplay.t()],
          session_replays_saved: [Umami.SessionReplaySaved.t()],
          heatmap_events: [Umami.HeatmapEvent.t()]
        }

  schema "website" do
    field(:id, Ecto.UUID, primary_key: true, source: :website_id)
    field(:name, :string)
    field(:domain, :string)
    field(:reset_at, Umami.PrismaDateTime)
    field(:deleted_at, Umami.PrismaDateTime)
    field(:recorder_enabled, :boolean, default: false)
    field(:replay_config, Umami.PrismaJson)
    belongs_to(:user, Umami.User, foreign_key: :user_id, type: Ecto.UUID)
    belongs_to(:create_user, Umami.User, foreign_key: :created_by, type: Ecto.UUID)
    belongs_to(:team, Umami.Team, foreign_key: :team_id, type: Ecto.UUID)
    has_many(:annotations, Umami.Annotation, foreign_key: :website_id, references: :id)
    has_many(:event_data, Umami.EventData, foreign_key: :website_id, references: :id)
    has_many(:reports, Umami.Report, foreign_key: :website_id, references: :id)
    has_many(:revenue, Umami.Revenue, foreign_key: :website_id, references: :id)
    has_many(:segments, Umami.Segment, foreign_key: :website_id, references: :id)
    has_many(:session_data, Umami.SessionData, foreign_key: :website_id, references: :id)
    has_many(:session_replays, Umami.SessionReplay, foreign_key: :website_id, references: :id)
    has_many(:session_replays_saved, Umami.SessionReplaySaved, foreign_key: :website_id, references: :id)
    has_many(:heatmap_events, Umami.HeatmapEvent, foreign_key: :website_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end