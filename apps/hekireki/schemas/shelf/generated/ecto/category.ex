defmodule Shelf.Category do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          color: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil,
          assets: [Shelf.Asset.t()],
          kits: [Shelf.Kit.t()],
          asset_model_defaults: [Shelf.AssetModel.t()],
          custom_fields: [Shelf.CustomField.t()]
        }

  schema "Category" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    field(:color, :string)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    has_many(:assets, Shelf.Asset, foreign_key: :category_id, references: :id)
    has_many(:kits, Shelf.Kit, foreign_key: :category_id, references: :id)
    has_many(:asset_model_defaults, Shelf.AssetModel, foreign_key: :default_category_id, references: :id)
    many_to_many(:custom_fields, Shelf.CustomField, join_through: "_CategoryToCustomField", join_keys: [A: :id, B: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end