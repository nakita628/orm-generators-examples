defmodule Umami.SessionReplaySaved do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Umami.PrismaDateTime, autogenerate: {Umami.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          name: String.t(),
          visit_id: Ecto.UUID.t(),
          inserted_at: DateTime.t() | nil,
          updated_at: DateTime.t() | nil,
          website: Umami.Website.t() | nil
        }

  schema "session_replay_saved" do
    field(:id, Ecto.UUID, primary_key: true, source: :saved_replay_id)
    field(:name, :string)
    field(:visit_id, Ecto.UUID)
    belongs_to(:website, Umami.Website, foreign_key: :website_id, type: Ecto.UUID)
    timestamps(inserted_at_source: :created_at)
  end
end