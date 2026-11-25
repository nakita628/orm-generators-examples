# retired: unread since v3.28.0. Replaced by NurturingProjectTraceDayCount, which the trace-day
# peer folds write and a replay rebuilds; an operator step may drop it once old writers are gone.
#
# day:: UTC `YYYY-MM-DD` of the root spans' append time.
# errorCount:: Traces trace recorded failed (lw.trace.trace_failed).
# lastTraceAt:: Append time of the newest root span.
# lastFailedTraceAt:: Append time of the newest trace_failed.
class NurturingProjectTraceDay < ApplicationRecord
  self.table_name = "NurturingProjectTraceDay"
  self.primary_key = %w[projectId day]

  attribute :traceCount, default: 0
  attribute :errorCount, default: 0
end
