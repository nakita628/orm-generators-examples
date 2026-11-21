defmodule LangWatch.ProcessManagerOutboxAttempt do
  @moduledoc """
  Why an outbox delivery failed, one row per FAILED attempt
  (specs/ops/dead-letter-recovery.feature). Successes write nothing, so the
  table grows with trouble, not with traffic. Rows die with their message:
  the cascade delete is the retention story, riding the existing sweep.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          project_id: String.t(),
          attempt: integer(),
          occurred_at: DateTime.t(),
          outcome: atom(),
          error_type: String.t(),
          error_message: String.t(),
          retry_after_ms: integer() | nil,
          outbox: LangWatch.ProcessManagerOutbox.t() | nil
        }

  schema "ProcessManagerOutboxAttempt" do
    field(:id, :string, primary_key: true)
    field(:project_id, :string, source: :projectId)
    # 1-based attempt number, as the dispatcher counts it.
    field(:attempt, :integer)
    field(:occurred_at, LangWatch.PrismaDateTime, source: :occurredAt)
    field(:outcome, Ecto.Enum, values: [:retry_scheduled, :dead])
    field(:error_type, :string, source: :errorType)
    # The safe failure diagnostic the dispatcher already computes for its
    # span — never a raw provider body.
    field(:error_message, :string, source: :errorMessage)
    field(:retry_after_ms, :integer, source: :retryAfterMs)
    field(:outbox_id, :string, source: :outboxId)
    belongs_to(:outbox, LangWatch.ProcessManagerOutbox, foreign_key: :outbox_id, define_field: false, type: :string)
  end
end