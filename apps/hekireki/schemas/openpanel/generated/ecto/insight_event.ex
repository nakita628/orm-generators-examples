defmodule OpenPanel.InsightEvent do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: OpenPanel.PrismaDateTime, autogenerate: {OpenPanel.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          event_kind: String.t(),
          change_from: map() | list() | String.t() | number() | boolean() | nil,
          change_to: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          insight: OpenPanel.ProjectInsight.t() | nil
        }

  schema "insight_events" do
    field(:id, Ecto.UUID, primary_key: true, read_after_writes: true)
    field(:event_kind, :string, source: :eventKind)
    field(:change_from, OpenPanel.PrismaJson, source: :changeFrom)
    field(:change_to, OpenPanel.PrismaJson, source: :changeTo)
    field(:insight_id, Ecto.UUID, source: :insightId)
    belongs_to(:insight, OpenPanel.ProjectInsight, foreign_key: :insight_id, define_field: false, type: Ecto.UUID)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end