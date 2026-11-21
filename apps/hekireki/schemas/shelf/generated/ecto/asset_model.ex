defmodule Shelf.AssetModel do
  @moduledoc """
  *
  * A template/grouping entity for assets.
  * Groups individual assets of the same type (e.g., "Dell Latitude 5550")
  * and provides default field values when creating new assets from the model.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          image: String.t() | nil,
          thumbnail_image: String.t() | nil,
          default_valuation: float() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          default_category: Shelf.Category.t() | nil,
          organization: Shelf.Organization.t() | nil,
          created_by: Shelf.User.t() | nil,
          custom_field_defaults: [Shelf.AssetModelCustomFieldDefault.t()],
          assets: [Shelf.Asset.t()],
          model_requests: [Shelf.BookingModelRequest.t()]
        }

  schema "AssetModel" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    # Public URL (bucket `files`) of the model's cover image. Rendered by any
    # asset of this model that has no image of its own. Public — and therefore
    # never-expiring and cacheable — like Location.imageUrl; see
    # resolveAssetImage() for the cascade.
    field(:image, :string)
    # Public URL of the 108px thumbnail generated alongside `image`. Used at
    # list sizes so inheriting assets never render the full-size original.
    field(:thumbnail_image, :string, source: :thumbnailImage)
    field(:default_valuation, :float, source: :defaultValuation)
    field(:default_category_id, :string, source: :defaultCategoryId)
    field(:organization_id, :string, source: :organizationId)
    field(:user_id, :string, source: :userId)
    belongs_to(:default_category, Shelf.Category, foreign_key: :default_category_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:created_by, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    # Default custom field values. Like the category and valuation defaults,
    # each one fills the matching field only when it is empty, at the moment
    # an asset joins this model. Nothing stays linked afterwards.
    has_many(:custom_field_defaults, Shelf.AssetModelCustomFieldDefault, foreign_key: :asset_model_id, references: :id)
    has_many(:assets, Shelf.Asset, foreign_key: :asset_model_id, references: :id)
    has_many(:model_requests, Shelf.BookingModelRequest, foreign_key: :asset_model_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end