defmodule Postiz.Announcement do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Postiz.PrismaDateTime, autogenerate: {Postiz.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          title: String.t(),
          description: String.t(),
          color: atom(),
          inserted_at: DateTime.t()
        }

  schema "Announcement" do
    field(:title, :string)
    field(:description, :string)
    field(:color, Ecto.Enum, values: [:INFO, :WARNING, :ERROR], default: :INFO)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end