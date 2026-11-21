defmodule Trigger.SessionRun do
  @moduledoc """
  Historical record of every run a Session has owned. Append-only —
  rows are inserted on each `ensureRunForSession` claim, never updated.
  Lets us reconstruct the run timeline of a chat for debugging /
  dashboard surfaces. The relation cascades on Session delete (tied to
  the session lifecycle) but `runId` is a plain string column with no
  FK to TaskRun so run pruning is independent.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          run_id: String.t(),
          reason: String.t(),
          triggered_at: DateTime.t(),
          session: Trigger.Session.t() | nil
        }

  schema "SessionRun" do
    field(:id, :string, primary_key: true)
    # TaskRun.id (no FK — runs may be archived independently of session history)
    field(:run_id, :string, source: :runId)
    # One of: "initial" | "continuation" | "upgrade" | "manual".
    # Plain string for forward-compat with future trigger reasons.
    field(:reason, :string)
    field(:triggered_at, Trigger.PrismaDateTime, autogenerate: true, source: :triggeredAt)
    field(:session_id, :string, source: :sessionId)
    belongs_to(:session, Trigger.Session, foreign_key: :session_id, define_field: false, type: :string)
  end
end