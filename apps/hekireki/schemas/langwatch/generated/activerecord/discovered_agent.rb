# The same shape as DiscoveredPerson for non-human actors - Genie spaces,
# Copilot bots, OpenAI projects - plus whatever provider-native metadata the
# source published. Discovering an agent grants it nothing either.
#
# rawAgentId:: Genie space / Copilot bot / OpenAI project id, verbatim.
# metadata:: Workspace, plan, and kin - provider-native descriptive fields only. Never
#   credentials: this row is read by admin screens, not by any pull path.
class DiscoveredAgent < ApplicationRecord
  self.table_name = "DiscoveredAgent"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true, uniqueness: { scope: [ :provider, :rawAgentId ] }
  validates :provider, presence: true
  validates :rawAgentId, presence: true
  validates :displayText, presence: true
  validates :metadata, exclusion: { in: [ nil ] }
  validates :firstSeenAt, presence: true
  validates :lastSeenAt, presence: true
end
