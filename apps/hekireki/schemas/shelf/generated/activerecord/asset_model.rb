# *
# * A template/grouping entity for assets.
# * Groups individual assets of the same type (e.g., "Dell Latitude 5550")
# * and provides default field values when creating new assets from the model.
#
# image:: Public URL (bucket `files`) of the model's cover image. Rendered by any
#   asset of this model that has no image of its own. Public — and therefore
#   never-expiring and cacheable — like Location.imageUrl; see
#   resolveAssetImage() for the cascade.
# thumbnailImage:: Public URL of the 108px thumbnail generated alongside `image`. Used at
#   list sizes so inheriting assets never render the full-size original.
# custom_field_defaults:: Default custom field values. Like the category and valuation defaults,
#   each one fills the matching field only when it is empty, at the moment
#   an asset joins this model. Nothing stays linked afterwards.
class AssetModel < ApplicationRecord
  self.table_name = "AssetModel"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :name, presence: true

  belongs_to :default_category, class_name: "Category", foreign_key: "defaultCategoryId", optional: true,
                                inverse_of: :asset_model_defaults
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :asset_models
  belongs_to :created_by, class_name: "User", foreign_key: "userId", inverse_of: :asset_models
  has_many :custom_field_defaults, class_name: "AssetModelCustomFieldDefault", foreign_key: "assetModelId",
                                   inverse_of: :asset_model, dependent: :destroy
  has_many :assets, foreign_key: "assetModelId", inverse_of: :asset_model, dependent: :nullify
  has_many :model_requests, class_name: "BookingModelRequest", foreign_key: "assetModelId", inverse_of: :asset_model,
                            dependent: :destroy
end
