# Nurturing's own per-project usage: a peer fold over the owners' prompt, evaluator, scenario
# and simulation-run events (ARCHITECTURE §9), one peer fold per owner pipeline, summed per
# organization for the per-email traits. Traces are counted in NurturingProjectTraceDay.
#
# owner:: The owner pipeline whose peer fold wrote the row (prompt, evaluator, scenario, simulation),
#   or `evaluation`: the project's evaluation days, written by a subscriber, at most once a day.
# folded:: The project's usage as the fold folded it (NurturingProjectUsageState).
# appliedEventIds:: The event ids of the last delivery folded in, so a redelivery folds once.
# projectionVersion:: A row of another version reads as undecodable and is re-folded from the log.
# updatedAt:: Event-derived, so a replayed row equals the row it rebuilds.
class NurturingProjectUsageView < ApplicationRecord
  self.table_name = "NurturingProjectUsageView"
  self.primary_key = %w[projectId owner]

  attribute :appliedEventIds, default: -> { [] }
  alias_attribute :updated_at, :updatedAt

  validates :folded, exclusion: { in: [ nil ] }
  validates :projectionVersion, presence: true
end
