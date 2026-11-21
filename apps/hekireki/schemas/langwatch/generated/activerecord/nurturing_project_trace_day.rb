# Nurturing's own trace counts from trace's span_received and trace_failed (ARCHITECTURE §9), a
# root span being one trace, one row per project and UTC day of append time, incremented in place
# by each batch. A redelivered batch can count a trace twice; accepted (issue 9378). Kept 30 days.
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
