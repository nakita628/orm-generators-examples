defmodule Shelf.Update do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          title: String.t(),
          content: String.t(),
          url: String.t() | nil,
          image_url: String.t() | nil,
          publish_date: DateTime.t(),
          status: atom(),
          target_roles: atom(),
          click_count: integer(),
          view_count: integer(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          created_by: Shelf.User.t() | nil,
          user_reads: [Shelf.UserUpdateRead.t()]
        }

  schema "Update" do
    field(:id, :string, primary_key: true)
    field(:title, :string)
    field(:content, :string)
    field(:url, :string)
    field(:image_url, :string, source: :imageUrl)
    field(:publish_date, Shelf.PrismaDateTime, source: :publishDate)
    field(:status, Ecto.Enum, values: [:DRAFT, :PUBLISHED], default: :DRAFT)
    field(:target_roles, {:array, Ecto.Enum}, values: [:ADMIN, :BASE, :OWNER, :SELF_SERVICE, :MANAGER], source: :targetRoles)
    field(:click_count, :integer, default: 0, source: :clickCount)
    field(:view_count, :integer, default: 0, source: :viewCount)
    field(:created_by_id, :string, source: :createdById)
    belongs_to(:created_by, Shelf.User, foreign_key: :created_by_id, define_field: false, type: :string)
    has_many(:user_reads, Shelf.UserUpdateRead, foreign_key: :update_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end