# PROJECTION of the topic-clustering event stream (ADR-051): topics are
# `topics_recorded` events; this table is written ONLY by
# TopicModelFoldProjection (cursor: TopicModelProjection) and is
# rebuildable by replay. Kept named `Topic` because it is the domain
# entity every surface reads — the `*Projection` suffix is reserved for
# cursor/read-model tables.
#
# lastEventId:: The `topics_recorded` event that recorded this topic (its id in the
#   event log). Null means the row predates event ownership and has not
#   been seeded onto the stream yet; the projection sets it on every row
#   it writes.
class Topic < ApplicationRecord
  self.table_name = "Topic"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :automaticallyGenerated, default: true
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :name, presence: true
  validates :embeddings_model, presence: true
  validates :centroid, exclusion: { in: [ nil ] }
  validates :p95Distance, presence: true

  belongs_to :project, foreign_key: "projectId", inverse_of: :topics
  belongs_to :parent, class_name: "Topic", foreign_key: "parentId", optional: true, inverse_of: :subtopics
  has_many :subtopics, class_name: "Topic", foreign_key: "parentId", inverse_of: :parent,
                       dependent: :restrict_with_error
end
