defmodule Umami.Annotation do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Umami.PrismaDateTime, autogenerate: {Umami.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          user_id: Ecto.UUID.t() | nil,
          date: DateTime.t(),
          all_day: boolean(),
          note: String.t(),
          inserted_at: DateTime.t() | nil,
          updated_at: DateTime.t() | nil,
          website: Umami.Website.t() | nil
        }

  schema "annotation" do
    field(:id, Ecto.UUID, primary_key: true, source: :annotation_id)
    field(:user_id, Ecto.UUID)
    field(:date, Umami.PrismaDateTime)
    field(:all_day, :boolean, default: true)
    field(:note, :string)
    belongs_to(:website, Umami.Website, foreign_key: :website_id, type: Ecto.UUID)
    timestamps(inserted_at_source: :created_at)
  end
end