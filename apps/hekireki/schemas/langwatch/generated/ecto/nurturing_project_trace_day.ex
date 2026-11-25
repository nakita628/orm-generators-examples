defmodule LangWatch.NurturingProjectTraceDay do
  @moduledoc """
  retired: unread since v3.28.0. Replaced by NurturingProjectTraceDayCount, which the trace-day
  peer folds write and a replay rebuilds; an operator step may drop it once old writers are gone.
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