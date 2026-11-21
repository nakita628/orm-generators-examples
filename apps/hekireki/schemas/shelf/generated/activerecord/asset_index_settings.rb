class AssetIndexSettings < ApplicationRecord
  self.table_name = "AssetIndexSettings"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :columns,
            default: -> { JSON.parse("[ {\"name\": \"id\", \"visible\": true, \"position\": 0}, {\"name\": \"status\", \"visible\": true, \"position\": 1}, {\"name\": \"description\", \"visible\": true, \"position\": 2}, {\"name\": \"valuation\", \"visible\": true, \"position\": 3}, {\"name\": \"createdAt\", \"visible\": true, \"position\": 4}, {\"name\": \"category\", \"visible\": true, \"position\": 5}, {\"name\": \"tags\", \"visible\": true, \"position\": 6}, {\"name\": \"location\", \"visible\": true, \"position\": 7}, {\"name\": \"kit\", \"visible\": true, \"position\": 8}, {\"name\": \"custody\", \"visible\": true, \"position\": 9} ]") }
  attribute :freezeColumn, default: true
  attribute :showAssetImage, default: true
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :mode, { simple: "SIMPLE", advanced: "ADVANCED" }, default: :simple, validate: true

  validates :userId, uniqueness: { scope: :organizationId }

  belongs_to :user, foreign_key: "userId", inverse_of: :asset_index_settings
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :asset_index_settings
end
