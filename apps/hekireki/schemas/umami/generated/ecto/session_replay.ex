defmodule Umami.SessionReplay do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Umami.PrismaDateTime, autogenerate: {Umami.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          session_id: Ecto.UUID.t(),
          visit_id: Ecto.UUID.t(),
          chunk_index: integer(),
          events: binary(),
          event_count: integer(),
          started_at: DateTime.t(),
          ended_at: DateTime.t(),
          inserted_at: DateTime.t() | nil,
          website: Umami.Website.t() | nil
        }

  schema "session_replay" do
    field(:id, Ecto.UUID, primary_key: true, source: :replay_id)
    field(:session_id, Ecto.UUID)
    field(:visit_id, Ecto.UUID)
    field(:chunk_index, :integer)
    field(:events, :binary)
    field(:event_count, :integer)
    field(:started_at, Umami.PrismaDateTime)
    field(:ended_at, Umami.PrismaDateTime)
    belongs_to(:website, Umami.Website, foreign_key: :website_id, type: Ecto.UUID)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end