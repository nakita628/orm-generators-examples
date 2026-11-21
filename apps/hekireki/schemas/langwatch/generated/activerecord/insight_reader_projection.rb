# One reader's own state on one insight (seen, done, kept), folded from the same stream.
# An insight has one reader today, its owner; a later forward may add more.
class InsightReaderProjection < ApplicationRecord
  self.table_name = "InsightReaderProjection"

  validates :projectId, presence: true, uniqueness: { scope: [ :insightId, :userId ] }
  validates :insightId, presence: true
  validates :userId, presence: true
  validates :createdAt, presence: true
  validates :updatedAt, presence: true
  validates :occurredAt, presence: true
  validates :acceptedAt, presence: true
  validates :lastEventId, presence: true
  validates :projectionVersion, presence: true
end
