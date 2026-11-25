# Nurturing's own trace counts (ARCHITECTURE §9), one row per project, UTC day of append time and
# source: `spans` from trace's span_received peer fold (a customer root span is one trace),
# `failures` from trace_failed's. Each fold owns its rows, so the two never overwrite each other;
# readers sum a day's sources. Rebuilt by a projection replay. Kept 30 days.
#
# day:: UTC `YYYY-MM-DD` of the facts' append time.
# source:: The peer fold that wrote the row: `spans` or `failures`.
# lastTraceAt:: Append time of the newest root span.
# lastFailedTraceAt:: Append time of the newest trace_failed.
# appliedEventIds:: The event ids of the last delivery folded in, so a redelivery folds once.
# projectionVersion:: A row of another version reads as undecodable and is left out of the read.
class NurturingProjectTraceDayCount < ApplicationRecord
  self.table_name = "NurturingProjectTraceDayCount"
  self.primary_key = %w[projectId day source]

  attribute :traceCount, default: 0
  attribute :errorCount, default: 0
  attribute :appliedEventIds, default: -> { [] }

  validates :projectionVersion, presence: true
end
