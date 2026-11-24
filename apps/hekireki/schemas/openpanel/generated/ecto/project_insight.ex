defmodule OpenPanel.ProjectInsight do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: OpenPanel.PrismaDateTime, autogenerate: {OpenPanel.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          project_id: String.t(),
          module_key: String.t(),
          dimension_key: String.t(),
          window_kind: String.t(),
          state: atom(),
          title: String.t(),
          summary: String.t() | nil,
          display_name: String.t(),
          payload: map() | list() | String.t() | number() | boolean(),
          direction: String.t() | nil,
          impact_score: float(),
          severity_band: String.t() | nil,
          relevance_score: float() | nil,
          ai_summary: String.t() | nil,
          ai_category: String.t() | nil,
          email_worthy: boolean(),
          reference_worthy: boolean(),
          enriched_at: DateTime.t() | nil,
          enrich_version: integer() | nil,
          version: integer(),
          thread_id: Ecto.UUID.t(),
          window_start: DateTime.t() | nil,
          window_end: DateTime.t() | nil,
          first_detected_at: DateTime.t(),
          last_seen_at: DateTime.t(),
          updated_at: DateTime.t(),
          events: [OpenPanel.InsightEvent.t()]
        }

  schema "project_insights" do
    field(:id, Ecto.UUID, primary_key: true, read_after_writes: true)
    field(:project_id, :string, source: :projectId)
    field(:module_key, :string, source: :moduleKey)
    field(:dimension_key, :string, source: :dimensionKey)
    field(:window_kind, :string, source: :windowKind)
    field(:state, Ecto.Enum, values: [:active, :suppressed, :closed], default: :active)
    field(:title, :string)
    field(:summary, :string)
    field(:display_name, :string, default: "", source: :displayName)
    # [IPrismaProjectInsightPayload]
    field(:payload, OpenPanel.PrismaJson, default: %{})
    field(:direction, :string)
    field(:impact_score, :float, default: 0.0, source: :impactScore)
    field(:severity_band, :string, source: :severityBand)
    field(:relevance_score, :float, source: :relevanceScore)
    field(:ai_summary, :string, source: :aiSummary)
    field(:ai_category, :string, source: :aiCategory)
    field(:email_worthy, :boolean, default: false, source: :emailWorthy)
    field(:reference_worthy, :boolean, default: false, source: :referenceWorthy)
    field(:enriched_at, OpenPanel.PrismaDateTime, source: :enrichedAt)
    field(:enrich_version, :integer, source: :enrichVersion)
    field(:version, :integer, default: 1)
    field(:thread_id, Ecto.UUID, read_after_writes: true, source: :threadId)
    field(:window_start, OpenPanel.PrismaDateTime, source: :windowStart)
    field(:window_end, OpenPanel.PrismaDateTime, source: :windowEnd)
    field(:first_detected_at, OpenPanel.PrismaDateTime, autogenerate: true, source: :firstDetectedAt)
    field(:last_seen_at, OpenPanel.PrismaDateTime, autogenerate: true, source: :lastSeenAt)
    has_many(:events, OpenPanel.InsightEvent, foreign_key: :insight_id, references: :id)
    timestamps(inserted_at: false, updated_at_source: :lastUpdatedAt)
  end
end