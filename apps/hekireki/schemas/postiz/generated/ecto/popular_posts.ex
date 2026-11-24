defmodule Postiz.PopularPosts do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Postiz.PrismaDateTime, autogenerate: {Postiz.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          category: String.t(),
          topic: String.t(),
          content: String.t(),
          hook: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "PopularPosts" do
    field(:category, :string)
    field(:topic, :string)
    field(:content, :string)
    field(:hook, :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end