defmodule Shelf.AssetIndexSettings do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          mode: atom(),
          columns: map(),
          freeze_column: boolean(),
          show_asset_image: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil
        }

  schema "AssetIndexSettings" do
    field(:id, :string, primary_key: true)
    field(:mode, Ecto.Enum, values: [:SIMPLE, :ADVANCED], default: :SIMPLE)
    field(:columns, :map)
    field(:freeze_column, :boolean, default: true, source: :freezeColumn)
    field(:show_asset_image, :boolean, default: true, source: :showAssetImage)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end