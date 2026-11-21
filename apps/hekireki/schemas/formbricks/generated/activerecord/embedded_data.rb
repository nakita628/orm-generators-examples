# Defines one Embedded Data field: the extra context stored on a response that isn't an
# answer to a question (brand, plan, page type, a computed score, ...).
#
# A field is either **local** — used by a single survey, which is how existing variables and
# hidden fields are migrated — or **shared**, meaning it is part of the workspace library and can
# be linked by many surveys. There is no stored flag for that: `surveyId` is the owner and `key`
# is the library name, so a local field has `surveyId` set and `key` null, and a shared field the
# other way round. Deriving it rather than storing a boolean removes a column that could drift
# out of step with the two fields it stands for.
#
# @property key - Library name, unique per workspace and immutable. Null for local fields
# @property name - Human-readable display label
# @property source - Where the value comes from (computed | ingested). `reserved` is part of the
# shared enum but never stored as a row: those fields are a static catalog in code
# @property dataType - Declared value type, used for coercion at ingest
# @property defaultValue - Initial value (computed) or fallback when nothing arrives (ingested)
# @property locked - Ingested only: ignore external writes and use the default instead
# @property survey - The owning survey for a local field; cascades so nothing is orphaned
#
# defaultValue:: [EmbeddedDataDefaultValue]
class EmbeddedData < ApplicationRecord
  self.table_name = "EmbeddedData"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :locked, default: false

  enum :source, { computed: "computed", ingested: "ingested", reserved: "reserved" }, validate: true
  enum :dataType, { string: "string", number: "number", boolean: "boolean", date: "date" }, default: :string,
                                                                                            validate: true

  validates :id, uniqueness: { scope: :workspaceId }
  validates :key, uniqueness: { scope: :workspaceId, allow_nil: true }
  validates :name, presence: true

  belongs_to :survey, foreign_key: [ :surveyId, :workspaceId ], primary_key: [ :id, :workspaceId ], optional: true,
                      inverse_of: :embedded_data
  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :embedded_data
  has_many :survey_links, class_name: "SurveyEmbeddedData", foreign_key: [ :embeddedDataId, :workspaceId ],
                          primary_key: [ :id, :workspaceId ], inverse_of: :embedded_data, dependent: :destroy
end
