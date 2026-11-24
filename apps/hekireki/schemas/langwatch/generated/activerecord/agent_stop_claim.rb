# The one open stop an agent may have: the api takes this row before it sends a stop's request,
# so two requests at once cannot both open one. A later request takes it over once the stop it
# names has closed.
# `sentAt` is set once the request command was sent: from then on the claim passes on only after
# its stop closed, never on age. `decision` is taken by the one approve or deny that wins.
# `decisionTakenAt` is set while that decision is in flight and cleared once its command was
# sent; a decision left in flight is taken over once abandoned.
class AgentStopClaim < ApplicationRecord
  self.table_name = "AgentStopClaim"
  self.primary_key = %w[organizationId agentId]

  validates :stopId, presence: true
  validates :claimedAt, presence: true
end
