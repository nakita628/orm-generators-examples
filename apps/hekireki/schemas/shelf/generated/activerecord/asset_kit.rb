# Pivot linking an `Asset` to the `Kit` it belongs to.
#
# `@@unique([assetId])` enforces "at most one kit per asset" today.
# The structure already allows multi-row-per-asset; if multi-kit
# allocation lands later, drop the `assetId` unique constraint and
# add a type-aware single-row trigger + sum-within-total CONSTRAINT
# TRIGGER (parallel to the Custody / BookingAsset pivots).
class AssetKit < ApplicationRecord
  self.table_name = "AssetKit"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :quantity, default: 1
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :assetId, uniqueness: { scope: :kitId }

  belongs_to :asset, foreign_key: "assetId", inverse_of: :asset_kits
  belongs_to :kit, foreign_key: "kitId", inverse_of: :asset_kits
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :asset_kits
  has_many :asset_locations, foreign_key: "assetKitId", inverse_of: :asset_kit, dependent: :destroy
end
