defmodule LangWatch.SavedView do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          filters: map(),
          query: String.t() | nil,
          period: map() | nil,
          order: integer(),
          kind: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          project: LangWatch.Project.t() | nil,
          user: LangWatch.User.t() | nil
        }

  schema "SavedView" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:filters, :map, default: %{})
    field(:query, :string)
    field(:period, :map)
    field(:order, :integer, default: 0)
    # Discriminator for the two storage shapes that share this table:
    # - "v1-traces-filter" (default) - legacy /messages page filter map saved views
    # - "v2-traces-lens"             - new traces v2 lens configs (columns, addons, grouping, sort, filterText) encoded inside `filters` JSON
    # Each client filters by its own kind on read so neither sees the other's rows.
    field(:kind, :string, default: "v1-traces-filter")
    field(:project_id, :string, source: :projectId)
    field(:user_id, :string, source: :userId)
    belongs_to(:project, LangWatch.Project, foreign_key: :project_id, define_field: false, type: :string)
    belongs_to(:user, LangWatch.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end