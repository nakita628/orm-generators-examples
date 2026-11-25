# One source's presence fact for one discovered agent. An agent is gone only when every
# source that listed it marked it gone; an agent no source listed (older sightings) never is.
#
# goneAt:: Set when this source's listing stopped naming the agent; cleared when it names it again.
class DiscoveredAgentListing < ApplicationRecord
  self.table_name = "DiscoveredAgentListing"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true
  validates :ingestionSourceId, presence: true, uniqueness: { scope: :discoveredAgentId }
  validates :discoveredAgentId, presence: true
  validates :lastListedAt, presence: true
end
