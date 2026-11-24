# Whether a target was already stopped before a stop first touched it, recorded once before the
# stop's platform call, so a retried call after a success still knows the change is the stop's own.
class AgentStopTargetPriorState < ApplicationRecord
  self.table_name = "AgentStopTargetPriorState"
  self.primary_key = %w[organizationId stopId targetId]

  validates :wasStopped, inclusion: { in: [ true, false ] }
  validates :recordedAt, presence: true
end
