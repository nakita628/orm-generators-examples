defmodule Shelf.Image do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          content_type: String.t(),
          alt_text: String.t() | nil,
          blob: binary(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          owner_org: Shelf.Organization.t() | nil,
          user: Shelf.User.t() | nil,
          location: Shelf.Location.t() | nil,
          organization: Shelf.Organization.t() | nil
        }

  schema "Image" do
    field(:id, :string, primary_key: true)
    field(:content_type, :string, source: :contentType)
    field(:alt_text, :string, source: :altText)
    field(:blob, :binary)
    field(:owner_org_id, :string, source: :ownerOrgId)
    field(:user_id, :string, source: :userId)
    belongs_to(:owner_org, Shelf.Organization, foreign_key: :owner_org_id, define_field: false, type: :string)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    has_one(:location, Shelf.Location, foreign_key: :image_id, references: :id)
    has_one(:organization, Shelf.Organization, foreign_key: :image_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end