defmodule LangWatch.ScimSyncState do
  @moduledoc """
  D08 — one connection's directory sync, as the `scim-sync` pipeline's fold
  writes it.

  Pure event truth: every column is fold-written, rows are never deleted,
  and REVOKED is a tombstone rather than an absence — a dead letter that
  vanished when a connection was torn down would be a removal nobody could
  check afterwards.

  No relation to Organization or SsoConnection, matching `SsoConnection`
  itself: a projection is rebuilt from events, and a foreign key would make
  a replay depend on relational state the log does not carry.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          connection_id: String.t(),
          organization_id: String.t(),
          state: String.t(),
          last_pushed_at: DateTime.t() | nil,
          last_failure: map() | nil,
          dead_letters: map(),
          revoked_cause: String.t() | nil,
          occurred_at: DateTime.t(),
          last_event_id: String.t(),
          accepted_at: DateTime.t(),
          projection_version: String.t(),
          created_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "ScimSyncState" do
    # The aggregate: the connection's id, which IS its sync's id.
    field(:id, :string, primary_key: true)
    field(:connection_id, :string, source: :connectionId)
    field(:organization_id, :string, source: :organizationId)
    # TOKEN_ISSUED | SYNCING | ERROR | REVOKED
    field(:state, :string)
    field(:last_pushed_at, LangWatch.PrismaDateTime, source: :lastPushedAt)
    # The failure standing right now: { op, errorCode, attempts, retiredAtMs,
    # userId, occurredAtMs }. A reason CODE and ids — never a token, never a
    # secret, never an internal hostname.
    field(:last_failure, :map, source: :lastFailure)
    # Every failure retired without ever being applied, newest last. The dead
    # letters an administrator has to see; retiring never drops one.
    field(:dead_letters, :map, source: :deadLetters)
    # revoke | teardown, once the sync has ended.
    field(:revoked_cause, :string, source: :revokedCause)
    field(:occurred_at, LangWatch.PrismaDateTime, source: :occurredAt)
    field(:last_event_id, :string, source: :lastEventId)
    field(:accepted_at, LangWatch.PrismaDateTime, source: :acceptedAt)
    field(:projection_version, :string, source: :projectionVersion)
    # Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    # would make a replayed row differ from the row it rebuilds, which is
    # exactly the parity this projection has to keep.
    field(:created_at, LangWatch.PrismaDateTime, source: :createdAt)
    field(:updated_at, LangWatch.PrismaDateTime, source: :updatedAt)
  end
end