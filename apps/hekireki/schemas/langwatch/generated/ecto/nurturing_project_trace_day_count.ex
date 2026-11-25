defmodule LangWatch.NurturingProjectTraceDayCount do
  @moduledoc """
  Nurturing's own trace counts (ARCHITECTURE §9), one row per project, UTC day of append time and
  source: `spans` from trace's span_received peer fold (a customer root span is one trace),
  `failures` from trace_failed's. Each fold owns its rows, so the two never overwrite each other;
  readers sum a day's sources. Rebuilt by a projection replay. Kept 30 days.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          project_id: String.t(),
          day: String.t(),
          source: String.t(),
          trace_count: integer(),
          error_count: integer(),
          last_trace_at: DateTime.t() | nil,
          last_failed_trace_at: DateTime.t() | nil,
          applied_event_ids: [String.t()],
          projection_version: String.t()
        }

  schema "NurturingProjectTraceDayCount" do
    field(:project_id, :string, primary_key: true, source: :projectId)
    # UTC `YYYY-MM-DD` of the facts' append time.
    field(:day, :string, primary_key: true)
    # The peer fold that wrote the row: `spans` or `failures`.
    field(:source, :string, primary_key: true)
    field(:trace_count, :integer, default: 0, source: :traceCount)
    field(:error_count, :integer, default: 0, source: :errorCount)
    # Append time of the newest root span.
    field(:last_trace_at, LangWatch.PrismaDateTime, source: :lastTraceAt)
    # Append time of the newest trace_failed.
    field(:last_failed_trace_at, LangWatch.PrismaDateTime, source: :lastFailedTraceAt)
    # The event ids of the last delivery folded in, so a redelivery folds once.
    field(:applied_event_ids, {:array, :string}, source: :appliedEventIds)
    # A row of another version reads as undecodable and is left out of the read.
    field(:projection_version, :string, source: :projectionVersion)
  end
end