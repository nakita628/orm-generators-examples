class AssetFilterPreset < ApplicationRecord
  self.table_name = "AssetFilterPreset"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :starred, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :name, presence: true, uniqueness: { scope: [ :organizationId, :ownerId ] }
  validates :query, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :asset_filter_presets
  belongs_to :owner, class_name: "User", foreign_key: "ownerId", inverse_of: :asset_filter_presets
end
