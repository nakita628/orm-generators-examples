# One immutable snapshot of an evaluator; the Evaluator row is always the latest.
#
# authorKind:: "user", "api_key" or "system": who wrote the version when authorId is null.
class EvaluatorVersion < ApplicationRecord
  self.table_name = "EvaluatorVersion"
  self.inheritance_column = nil

  alias_attribute :created_at, :createdAt

  validates :projectId, presence: true
  validates :evaluatorId, presence: true, uniqueness: { scope: :version }
  validates :version, presence: true
  validates :name, presence: true
  validates :type, presence: true
  validates :config, exclusion: { in: [ nil ] }
end
