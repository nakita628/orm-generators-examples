defmodule LangWatch.LangyConversationProjection do
  @moduledoc """
  Rebuildable operational state folded from Langy's canonical event stream.
  The source of truth remains the ClickHouse event log.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          project_id: String.t(),
          conversation_id: String.t(),
          user_id: String.t(),
          title: String.t() | nil,
          title_source: atom(),
          origin: String.t(),
          status: String.t(),
          is_shared: boolean(),
          shared_at: float() | nil,
          shared_by_id: String.t() | nil,
          message_count: integer(),
          last_activity_at: float() | nil,
          current_turn_id: String.t() | nil,
          last_error: String.t() | nil,
          last_model: String.t() | nil,
          pending_handoff_token: String.t() | nil,
          pending_handoff_turn_id: String.t() | nil,
          run_token: String.t() | nil,
          archived_at: float() | nil,
          created_at: float(),
          updated_at: float(),
          occurred_at: float(),
          accepted_at: float(),
          last_event_id: String.t(),
          projection_version: String.t()
        }

  schema "LangyConversationProjection" do
    field(:id, :string, primary_key: true)
    field(:project_id, :string, source: :projectId)
    field(:conversation_id, :string, source: :conversationId)
    # Owner; first-writer-wins from the first message's userId.
    field(:user_id, :string, source: :userId)
    field(:title, :string)
    # derived | auto | user — governs auto-title precedence (user is sticky).
    field(:title_source, Ecto.Enum, values: [:derived, :auto, :user], source: :titleSource)
    # Who started the conversation: `interactive` for the person in the panel, `run` for a
    # module acting for them with nobody at the keyboard. A row folded before the column
    # existed reads as interactive.
    field(:origin, :string, default: "interactive")
    field(:status, :string)
    field(:is_shared, :boolean, default: false, source: :isShared)
    field(:shared_at, :float, source: :sharedAt)
    field(:shared_by_id, :string, source: :sharedById)
    field(:message_count, :integer, default: 0, source: :messageCount)
    field(:last_activity_at, :float, source: :lastActivityAt)
    # The turn in flight, or null when idle (busy guard reads this).
    field(:current_turn_id, :string, source: :currentTurnId)
    field(:last_error, :string, source: :lastError)
    # The model the latest accepted turn ran on; reopening seeds the picker from it.
    field(:last_model, :string, source: :lastModel)
    # ADR-048 shutdown-handoff resume token (opaque, worker-authored).
    field(:pending_handoff_token, :string, source: :pendingHandoffToken)
    field(:pending_handoff_turn_id, :string, source: :pendingHandoffTurnId)
    # Per-conversation HMAC key for worker stream frames; server-side only.
    field(:run_token, :string, source: :runToken)
    field(:archived_at, :float, source: :archivedAt)
    field(:created_at, :float, source: :createdAt)
    field(:updated_at, :float, source: :updatedAt)
    # Business time of the latest applied event.
    field(:occurred_at, :float, source: :occurredAt)
    # Event-log acceptance time. Together with lastEventId this is the cursor.
    field(:accepted_at, :float, source: :acceptedAt)
    field(:last_event_id, :string, source: :lastEventId)
    field(:projection_version, :string, source: :projectionVersion)
  end
end