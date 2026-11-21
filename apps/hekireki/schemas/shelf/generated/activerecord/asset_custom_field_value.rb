class AssetCustomFieldValue < ApplicationRecord
  self.table_name = "AssetCustomFieldValue"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :value, exclusion: { in: [ nil ] }

  belongs_to :asset, foreign_key: "assetId", inverse_of: :custom_fields
  belongs_to :custom_field, foreign_key: "customFieldId", inverse_of: :asset_custom_fields_values
end
