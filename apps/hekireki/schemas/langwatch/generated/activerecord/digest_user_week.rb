# A person's ISO week of coding-agent activity in one organization, folded from the agents' facts.
#
# state:: Sessions (a capped set of ids) and the agents' metric series, converged last-write-wins.
class DigestUserWeek < ApplicationRecord
  self.table_name = "DigestUserWeek"

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :userId, presence: true
  validates :organizationId, presence: true
  validates :isoWeek, presence: true
  validates :state, exclusion: { in: [ nil ] }
  validates :occurredAt, presence: true
  validates :lastEventId, presence: true
  validates :acceptedAt, presence: true
  validates :projectionVersion, presence: true
end
