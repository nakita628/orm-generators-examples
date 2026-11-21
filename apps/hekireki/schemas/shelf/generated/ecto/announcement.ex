defmodule Shelf.Announcement do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          content: String.t(),
          link: String.t() | nil,
          link_text: String.t() | nil,
          published: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "Announcement" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:content, :string)
    field(:link, :string)
    field(:link_text, :string, source: :linkText)
    field(:published, :boolean, default: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end