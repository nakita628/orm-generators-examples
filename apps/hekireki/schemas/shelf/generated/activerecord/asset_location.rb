class AssetLocation < ApplicationRecord
  self.table_name = "AssetLocation"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :quantity, default: 1
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  belongs_to :asset, foreign_key: "assetId", inverse_of: :asset_locations
  belongs_to :location, foreign_key: "locationId", inverse_of: :asset_locations
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :asset_locations
  belongs_to :asset_kit, foreign_key: "assetKitId", optional: true, inverse_of: :asset_locations
end
