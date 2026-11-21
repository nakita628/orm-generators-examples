defmodule LangWatch.LangyConversationTurnProjection do
  @moduledoc """
  Rebuildable per-turn operational render document.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          project_id: String.t(),
          conversation_id: String.t(),
          turn_id: String.t(),
          status: String.t(),
          question_parts: map(),
          answer_parts: map(),
          tool_calls: map(),
          plan: map() | nil,
          error: String.t() | nil,
          started_at: float() | nil,
          ended_at: float() | nil,
          created_at: float(),
          updated_at: float(),
          occurred_at: float(),
          accepted_at: float(),
          last_event_id: String.t(),
          projection_version: String.t()
        }

  schema "LangyConversationTurnProjection" do
    field(:id, :string, primary_key: true)
    field(:project_id, :string, source: :projectId)
    field(:conversation_id, :string, source: :conversationId)
    field(:turn_id, :string, source: :turnId)
    # LangyConversationTurnStatus. Deliberately TEXT, not a Postgres enum: this
    # is a PROJECTION — one writer (the fold), rebuildable from the event log,
    # and nothing queries or indexes on it. An enum bought no selectivity here,
    # only typo protection, and charged a migration for every new status that
    # Postgres then could not remove. The one definition lives in
    # `LANGY_CONVERSATION_TURN_STATUS`; the repository validates against it.
    field(:status, :string)
    # LangyMessagePart[] — the user question that opened the turn.
    field(:question_parts, :map, source: :questionParts)
    # LangyMessagePart[] — the agent's whole final answer.
    field(:answer_parts, :map, source: :answerParts)
    # LangyTurnToolCall[] — tool-call lifecycle audit in initiation order.
    field(:tool_calls, :map, source: :toolCalls)
    # LangyPlanItemData[] | null — null means the turn never kept a plan.
    field(:plan, :map)
    field(:error, :string)
    field(:started_at, :float, source: :startedAt)
    field(:ended_at, :float, source: :endedAt)
    field(:created_at, :float, source: :createdAt)
    field(:updated_at, :float, source: :updatedAt)
    field(:occurred_at, :float, source: :occurredAt)
    field(:accepted_at, :float, source: :acceptedAt)
    field(:last_event_id, :string, source: :lastEventId)
    field(:projection_version, :string, source: :projectionVersion)
  end
end