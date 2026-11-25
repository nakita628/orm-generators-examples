# A project's ISO week, folded from its owners' events. The id is the fold's aggregate key.
#
# state:: The fold state: counts, sums, a fixed-bucket latency histogram, capped model and evaluator maps.
class DigestProjectWeek < ApplicationRecord
  self.table_name = "DigestProjectWeek"

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :projectId, presence: true
  validates :organizationId, presence: true
  validates :isoWeek, presence: true
  validates :state, exclusion: { in: [ nil ] }
  validates :occurredAt, presence: true
  validates :lastEventId, presence: true
  validates :acceptedAt, presence: true
  validates :projectionVersion, presence: true
end
