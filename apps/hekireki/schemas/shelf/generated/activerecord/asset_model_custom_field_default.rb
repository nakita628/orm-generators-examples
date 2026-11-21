# A custom field value an AssetModel gives the assets that join it.
#
# `value` has the exact shape of `AssetCustomFieldValue.value`, so filling an
# asset is a plain copy. Rows cascade away with their model or field.
class AssetModelCustomFieldDefault < ApplicationRecord
  self.table_name = "AssetModelCustomFieldDefault"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :value, exclusion: { in: [ nil ] }
  validates :customFieldId, presence: true, uniqueness: { scope: :assetModelId }

  belongs_to :asset_model, foreign_key: "assetModelId", inverse_of: :custom_field_defaults
end
