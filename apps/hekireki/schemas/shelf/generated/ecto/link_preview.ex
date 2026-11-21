defmodule Shelf.LinkPreview do
  @moduledoc """
  What Shelf knows about one web link used in a workspace's custom fields,
  keyed by its `LinkKey` from @shelf/links. Two writers on disjoint columns:
  the link preview job (fetch columns) and people (customTitle).
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          url: String.t(),
          page_title: String.t() | nil,
          outcome: atom() | nil,
          attempts: integer(),
          fetched_at: DateTime.t() | nil,
          next_fetch_at: DateTime.t() | nil,
          custom_title: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Shelf.Organization.t() | nil
        }

  schema "LinkPreview" do
    field(:id, :string, primary_key: true)
    # `LinkKey` from @shelf/links, never a raw field value. At most 2048
    # characters (CHECK constraint in the migration).
    field(:url, :string)
    field(:page_title, :string, source: :pageTitle)
    # Null until the first fetch finishes.
    field(:outcome, Ecto.Enum, values: [:TITLE, :NO_TITLE, :SIGN_IN, :BLOCKED, :UNREACHABLE])
    field(:attempts, :integer, default: 0)
    field(:fetched_at, Shelf.PrismaDateTime, source: :fetchedAt)
    # When the link may be fetched next; null means never again. A claim moves
    # it a few minutes ahead, which is also the lease on a fetch in flight.
    field(:next_fetch_at, Shelf.PrismaDateTime, autogenerate: true, source: :nextFetchAt)
    # The name a person gave the link in this workspace; wins over pageTitle.
    field(:custom_title, :string, source: :customTitle)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end