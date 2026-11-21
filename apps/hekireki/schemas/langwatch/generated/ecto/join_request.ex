defmodule LangWatch.JoinRequest do
  @moduledoc """
  The join-request pipeline's projection (ADR-117, D12 - see
  specs/identity/join-requests.feature). A pure event-truth head like
  `Identifier` and `SsoConnection`: every column is fold-written, rows are
  never deleted (a terminal state is a tombstone the admin panel stops
  listing, not an absence), and a replay rebuilds it whole-row from the
  event log.

  The fold cursor lives on the row rather than in a sibling table, because
  one row IS one aggregate here - a request.

  No relation to Organization or User on purpose, matching the two heads
  above: a projection is rebuilt from events, and a foreign key would make a
  replay depend on relational state the log does not carry.

  Membership is NOT here. An approval dispatches an attach on the grants
  ledger, exactly as accepting an invitation does, and that ledger fact
  carries `source: "join-request"` as its provenance - which is what puts a
  join approval on the customer's audit page beside every other membership
  change rather than in a history of its own.

  The requester's address is not here either: the DOMAIN is the fact, and it
  is the only part of an address a join decision may depend on.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          user_id: String.t(),
          organization_id: String.t(),
          domain: String.t(),
          state: String.t(),
          matched_via: String.t(),
          origin: String.t(),
          expires_at: DateTime.t() | nil,
          resolved_at: DateTime.t() | nil,
          resolved_by_type: String.t() | nil,
          resolved_by_id: String.t() | nil,
          withdrawal_cause: String.t() | nil,
          connection_id: String.t() | nil,
          occurred_at: DateTime.t(),
          last_event_id: String.t(),
          accepted_at: DateTime.t(),
          projection_version: String.t(),
          created_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "JoinRequest" do
    field(:id, :string, primary_key: true)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    # The normalized domain the match was made on.
    field(:domain, :string)
    # PENDING | APPROVED | REJECTED | EXPIRED | WITHDRAWN
    field(:state, :string)
    # How the domain was matched: `verified-identifier-domain` or `sso-connection-domain`.
    field(:matched_via, :string, source: :matchedVia)
    # Where the request was made: `web` or `cli` (ADR-171 v6). A request from
    # the terminal lands a Developer seat whatever the joiner seat says.
    field(:origin, :string, default: "web")
    # While PENDING: when the day-14 wake lapses it. Null once resolved.
    field(:expires_at, LangWatch.PrismaDateTime, source: :expiresAt)
    field(:resolved_at, LangWatch.PrismaDateTime, source: :resolvedAt)
    # user | policy | invite - who ended it. `policy` is domain auto-join's
    # own principal, deliberately not the one SSO auto-join uses.
    field(:resolved_by_type, :string, source: :resolvedByType)
    # The admin's user id, `domain-auto`, or the invitation id.
    field(:resolved_by_id, :string, source: :resolvedById)
    # user | invite-accepted, on a withdrawal.
    field(:withdrawal_cause, :string, source: :withdrawalCause)
    # The SSO connection a single sign-on arrival was admitted through; null otherwise.
    field(:connection_id, :string, source: :connectionId)
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