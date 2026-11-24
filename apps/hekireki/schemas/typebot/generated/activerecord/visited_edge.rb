class VisitedEdge < ApplicationRecord
  self.table_name = "VisitedEdge"
  self.primary_key = %w[resultId index]

  validates :edgeId, presence: true
  validates :index, presence: true, uniqueness: { scope: :resultId }

  belongs_to :result, foreign_key: "resultId", inverse_of: :edges
end
