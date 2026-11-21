# version:: The number the version is addressed and displayed by. Numbered rows hold
#   a gap-free sequence. The rolling autosave row rides on the counter
#   instead, which is what keeps it clear of that sequence.
# counterVersion:: `Experiment.workbenchVersion` at the moment this row was last written.
#   The rolling row is rewritten in place, so this is the only column that
#   says how recent its content is, and the version list sorts and pages by
#   it. It also joins a row to the live state: the row whose `counterVersion`
#   equals the experiment's counter holds what the workbench shows now.
# authorLabel:: Who wrote this version: "user", "langy" or "api".
# runId:: The run that wrote this version, when a run wrote it.
#
#   A run writes its cells into the workbench state, which advances the
#   counter. Without this, the page that started that run reads the bump as
#   somebody else's write, refuses to save, and asks the reader to reload
#   over their own unsaved edits. Named here, the page adopts the version and
#   carries on.
# schemaVersion:: Schema version of `state`, so a later shape change can be read forward.
class ExperimentVersion < ApplicationRecord
  self.table_name = "ExperimentVersion"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :autoSaved, default: false
  attribute :schemaVersion, default: "1"
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :experimentId, presence: true, uniqueness: { scope: :version }
  validates :projectId, presence: true
  validates :version, presence: true
  validates :counterVersion, presence: true
  validates :authorLabel, presence: true
  validates :state, exclusion: { in: [ nil ] }
end
