defmodule LangWatch.InsightDailyScheduleProjection do
  @moduledoc """
  One person's daily insights run on one board in a project, folded from its
  `insight_daily_schedule` aggregate (modules/insight). A read model: rebuildable by
  replaying the event log. Epochs in milliseconds.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          project_id: String.t(),
          user_id: String.t(),
          board_kind: String.t(),
          board_id: String.t(),
          board_name: String.t(),
          state: String.t(),
          off_reason: String.t() | nil,
          hour: integer() | nil,
          timezone: String.t() | nil,
          max_insights: integer() | nil,
          last_run_id: String.t() | nil,
          last_run_at: float() | nil,
          last_run_outcome: String.t() | nil,
          last_run_reason: String.t() | nil,
          last_run_filed: integer() | nil,
          last_run_renewed: integer() | nil,
          last_run_conversation_id: String.t() | nil,
          consecutive_skips: integer(),
          skip_streak_slot: float() | nil,
          created_at: float(),
          updated_at: float(),
          occurred_at: float(),
          accepted_at: float(),
          last_event_id: String.t(),
          projection_version: String.t()
        }

  schema "InsightDailyScheduleProjection" do
    # The schedule id, which is also its aggregate id: one per project, person and board.
    field(:id, :string, primary_key: true)
    field(:project_id, :string, source: :projectId)
    field(:user_id, :string, source: :userId)
    # dashboard | template. A pointer only, never a relation: a run checks the board is there.
    field(:board_kind, :string, source: :boardKind)
    field(:board_id, :string, source: :boardId)
    field(:board_name, :string, source: :boardName)
    # on | off | undecided.
    field(:state, :string, default: "undecided")
    # board_deleted | skipped_days | member_removed | project_archived. Why the system turned
    # it off. NULL while it is on, and when the person turned it off.
    field(:off_reason, :string, source: :offReason)
    field(:hour, :integer)
    field(:timezone, :string)
    field(:max_insights, :integer, source: :maxInsights)
    # The last run that recorded an outcome. NULL until one did.
    field(:last_run_id, :string, source: :lastRunId)
    field(:last_run_at, :float, source: :lastRunAt)
    # filed | nothing | failed | skipped
    field(:last_run_outcome, :string, source: :lastRunOutcome)
    field(:last_run_reason, :string, source: :lastRunReason)
    field(:last_run_filed, :integer, source: :lastRunFiled)
    # How many insights the run found still true and renewed.
    field(:last_run_renewed, :integer, source: :lastRunRenewed)
    field(:last_run_conversation_id, :string, source: :lastRunConversationId)
    # The scheduled days in a row whose run ended skipped, and the slot of the newest counted.
    field(:consecutive_skips, :integer, default: 0, source: :consecutiveSkips)
    field(:skip_streak_slot, :float, source: :skipStreakSlot)
    field(:created_at, :float, source: :createdAt)
    field(:updated_at, :float, source: :updatedAt)
    field(:occurred_at, :float, source: :occurredAt)
    field(:accepted_at, :float, source: :acceptedAt)
    field(:last_event_id, :string, source: :lastEventId)
    field(:projection_version, :string, source: :projectionVersion)
  end
end