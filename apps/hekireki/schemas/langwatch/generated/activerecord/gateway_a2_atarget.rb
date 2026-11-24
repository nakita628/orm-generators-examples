class GatewayA2ATarget < ApplicationRecord
  self.table_name = "GatewayA2ATarget"

  attribute :restrictToAllowList, default: false
  attribute :allowedVirtualKeyIds, default: -> { [] }
  attribute :blocked, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true
  validates :name, presence: true
  validates :agentName, presence: true
  validates :cardUrl, presence: true
  validates :upstreamUrl, presence: true
  validates :createdById, presence: true
end
