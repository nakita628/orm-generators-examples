defmodule LangWatch.Topic do
  @moduledoc """
  PROJECTION of the topic-clustering event stream (ADR-051): topics are
  `topics_recorded` events; this table is written ONLY by
  TopicModelFoldProjection (cursor: TopicModelProjection) and is
  rebuildable by replay. Kept named `Topic` because it is the domain
  entity every surface reads — the `*Projection` suffix is reserved for
  cursor/read-model tables.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          embeddings_model: String.t(),
          centroid: map(),
          p95_distance: float(),
          automatically_generated: boolean(),
          last_event_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          project: LangWatch.Project.t() | nil,
          parent: LangWatch.Topic.t() | nil,
          subtopics: [LangWatch.Topic.t()]
        }

  schema "Topic" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:embeddings_model, :string)
    field(:centroid, :map)
    field(:p95_distance, :float, source: :p95Distance)
    field(:automatically_generated, :boolean, default: true, source: :automaticallyGenerated)
    # The `topics_recorded` event that recorded this topic (its id in the
    # event log). Null means the row predates event ownership and has not
    # been seeded onto the stream yet; the projection sets it on every row
    # it writes.
    field(:last_event_id, :string, source: :lastEventId)
    field(:project_id, :string, source: :projectId)
    field(:parent_id, :string, source: :parentId)
    belongs_to(:project, LangWatch.Project, foreign_key: :project_id, define_field: false, type: :string)
    belongs_to(:parent, LangWatch.Topic, foreign_key: :parent_id, define_field: false, type: :string)
    has_many(:subtopics, LangWatch.Topic, foreign_key: :parent_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end