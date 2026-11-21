# scenario:: Deleting the case takes its history with it. Under relationMode="prisma"
#   the cascade is the client's, so a raw SQL delete still has to clear these
#   rows itself.
# projectId:: The project of the parent case, carried here so every read filters by
#   tenant without a join. The service is what keeps the pair honest: no
#   composite foreign key can, under relationMode="prisma".
# authorLabel:: Who wrote this version: "user", "api", "cli" or "langy".
# snapshot:: Envelope { schemaVersion, fields, changedFields } holding the editable
#   field set as saved and the names of the fields this save changed.
# schemaVersion:: Schema version of `snapshot`, so a later shape change can be read forward.
class ScenarioVersion < ApplicationRecord
  self.table_name = "ScenarioVersion"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :schemaVersion, default: 1
  alias_attribute :created_at, :createdAt

  validates :projectId, presence: true
  validates :version, presence: true, uniqueness: { scope: :scenarioId }
  validates :authorLabel, presence: true
  validates :snapshot, exclusion: { in: [ nil ] }

  belongs_to :scenario, foreign_key: "scenarioId", inverse_of: :versions
end
