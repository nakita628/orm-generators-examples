defmodule LangWatch.NurturingProjectTraceDay do
  @moduledoc """
  Nurturing's own trace counts from trace's span_received and trace_failed (ARCHITECTURE §9), a
  root span being one trace, one row per project and UTC day of append time, incremented in place
  by each batch. A redelivered batch can count a trace twice; accepted (issue 9378). Kept 30 days.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          project_id: String.t(),
          day: String.t(),
          trace_count: integer(),
          error_count: integer(),
          last_trace_at: DateTime.t() | nil,
          last_failed_trace_at: DateTime.t() | nil
        }

  schema "NurturingProjectTraceDay" do
    field(:project_id, :string, primary_key: true, source: :projectId)
    # UTC `YYYY-MM-DD` of the root spans' append time.
    field(:day, :string, primary_key: true)
    field(:trace_count, :integer, default: 0, source: :traceCount)
    # Traces trace recorded failed (lw.trace.trace_failed).
    field(:error_count, :integer, default: 0, source: :errorCount)
    # Append time of the newest root span.
    field(:last_trace_at, LangWatch.PrismaDateTime, source: :lastTraceAt)
    # Append time of the newest trace_failed.
    field(:last_failed_trace_at, LangWatch.PrismaDateTime, source: :lastFailedTraceAt)
  end
end