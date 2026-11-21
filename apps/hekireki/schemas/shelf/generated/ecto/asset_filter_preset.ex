defmodule Shelf.AssetFilterPreset do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          query: String.t(),
          starred: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Shelf.Organization.t() | nil,
          owner: Shelf.User.t() | nil
        }

  schema "AssetFilterPreset" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:query, :string)
    field(:starred, :boolean, default: false)
    field(:organization_id, :string, source: :organizationId)
    field(:owner_id, :string, source: :ownerId)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:owner, Shelf.User, foreign_key: :owner_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end