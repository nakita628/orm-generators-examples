defmodule LangWatch.InsightProjection do
  @moduledoc """
  One insight in its owner's inbox in a project, folded from its `insight` aggregate
  (modules/insight). A read model: rebuildable by replaying the event log. Epochs in
  milliseconds.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          project_id: String.t(),
          title: String.t(),
          body: String.t(),
          tone: String.t(),
          topic: String.t() | nil,
          valid_days: integer(),
          lwql: String.t() | nil,
          source_conversation_id: String.t() | nil,
          source_message_id: String.t() | nil,
          filed_by_user_id: String.t() | nil,
          filed_at: float(),
          renewed_at: float() | nil,
          created_at: float(),
          updated_at: float(),
          occurred_at: float(),
          accepted_at: float(),
          last_event_id: String.t(),
          projection_version: String.t(),
          board_id: String.t() | nil,
          board_name: String.t() | nil,
          widget_id: String.t() | nil,
          widget_name: String.t() | nil,
          filed_via: String.t(),
          replay_start: float() | nil,
          replay_end: float() | nil,
          replay_granularity_seconds: integer() | nil,
          replay_context: map() | nil,
          owner_user_id: String.t() | nil,
          finding_key: String.t() | nil
        }

  schema "InsightProjection" do
    # The insight id, which is also its aggregate id.
    field(:id, :string, primary_key: true)
    field(:project_id, :string, source: :projectId)
    field(:title, :string)
    field(:body, :string)
    # bad | watch | good
    field(:tone, :string)
    field(:topic, :string)
    field(:valid_days, :integer, source: :validDays)
    field(:lwql, :string)
    field(:source_conversation_id, :string, source: :sourceConversationId)
    field(:source_message_id, :string, source: :sourceMessageId)
    field(:filed_by_user_id, :string, source: :filedByUserId)
    field(:filed_at, :float, source: :filedAt)
    # Set by a later run that found the insight still true.
    field(:renewed_at, :float, source: :renewedAt)
    field(:created_at, :float, source: :createdAt)
    field(:updated_at, :float, source: :updatedAt)
    field(:occurred_at, :float, source: :occurredAt)
    field(:accepted_at, :float, source: :acceptedAt)
    field(:last_event_id, :string, source: :lastEventId)
    field(:projection_version, :string, source: :projectionVersion)
    # The board and widget it came from. A pointer only, never a relation: the row reads the
    # same once either is deleted, so the names are kept as filed.
    field(:board_id, :string, source: :boardId)
    field(:board_name, :string, source: :boardName)
    field(:widget_id, :string, source: :widgetId)
    field(:widget_name, :string, source: :widgetName)
    # chat | run
    field(:filed_via, :string, default: "chat", source: :filedVia)
    # The fixed window `lwql` replays over (epoch milliseconds, half-open) and its step.
    field(:replay_start, :float, source: :replayStart)
    field(:replay_end, :float, source: :replayEnd)
    field(:replay_granularity_seconds, :integer, source: :replayGranularitySeconds)
    # The values in force when it was filed: `{ period, parameters }`.
    field(:replay_context, :map, source: :replayContext)
    # Whose insight it is: the one person who reads it. NULL on a row folded before the
    # column existed, which is owned by whoever filed it (`filedByUserId`).
    field(:owner_user_id, :string, source: :ownerUserId)
    # The stable key a daily run matches a repeated finding by, so it renews the insight
    # and files no twin. NULL for an insight saved from a chat and on a row folded before
    # the column existed.
    field(:finding_key, :string, source: :findingKey)
  end
end