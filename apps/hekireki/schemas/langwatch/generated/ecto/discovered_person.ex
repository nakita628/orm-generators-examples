defmodule LangWatch.DiscoveredPerson do
  @moduledoc """
  A name-text seen on a provider's cost or audit rows: "the provider told us
  someone called m.silva exists". One row per provider-scoped identity.

  This is the unit of GDPR erasure (ADR-128 §9), and deliberately so: most of
  these people have no LangWatch login at all - contractors, seat holders who
  never signed in, anyone whose email a provider put on a cost row - so there
  is no platform user to key on and a user-deletion event can never fire for
  them.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          provider: String.t(),
          raw_actor_id: String.t(),
          display_text: String.t(),
          kind: String.t(),
          department: String.t() | nil,
          first_seen_at: DateTime.t(),
          last_seen_at: DateTime.t(),
          erased_at: DateTime.t() | nil,
          money_rows_pending_at: DateTime.t() | nil,
          money_rebuild_since: String.t() | nil,
          suspended_at: DateTime.t() | nil,
          suspended_reason: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "DiscoveredPerson" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    # "anthropic" | "openai" | "databricks" | ... - the provider that named them.
    field(:provider, :string)
    # What the provider calls them, verbatim. The join key onto the money rows'
    # RawActorId. Replaced by a deterministic pseudonym on erasure, which keeps
    # the row addressable for spend attribution without keeping the identifier.
    field(:raw_actor_id, :string, source: :rawActorId)
    # Name or email text as seen. Pseudonymized alongside rawActorId on erasure.
    field(:display_text, :string, source: :displayText)
    # person | service_account - determined deterministically per ADR-128 §10
    # (a bare UUID in Databricks `executed_by` is a service account, never an
    # agent), never guessed from the shape of a name.
    field(:kind, :string)
    # The department the provider's directory filed this person under, verbatim.
    # An OBSERVATION, not the `Department` entity: it is free text a tenant owns
    # on their side, it exists for people who hold no LangWatch account, and
    # nothing rolls spend up by it. Spend attribution stays on `Department`, via
    # the membership the match engine proves (directoryDepartmentSync.service).
    # Null means no directory has named one - never "no department", because a
    # blank field in the tenant's directory and a person no directory lists are
    # the same null here.
    field(:department, :string)
    field(:first_seen_at, LangWatch.PrismaDateTime, source: :firstSeenAt)
    field(:last_seen_at, LangWatch.PrismaDateTime, source: :lastSeenAt)
    # Set when this row's identity has been erased. The row survives so its
    # spend stays attributable to *someone*; the identifier does not.
    field(:erased_at, LangWatch.PrismaDateTime, source: :erasedAt)
    # Set alongside `erasedAt` and cleared once the daily cost rows have been
    # removed and their rebuild requested. Non-null means the ClickHouse half
    # of the erasure did not finish — a crash between the delete and the
    # replay leaves days short by the erased amount, and without this marker a
    # re-run reads `erasedAt`, short-circuits, and reports a clean erasure that
    # did no work. The re-run resumes at that step instead.
    field(:money_rows_pending_at, LangWatch.PrismaDateTime, source: :moneyRowsPendingAt)
    # The earliest day to rebuild from, recorded BEFORE the rows are deleted.
    # Once they are gone there is nothing left to ask which days they were on,
    # so a resume with no record of this could only replay everything or
    # nothing. Null means nothing was rebuildable — either no day carried the
    # identifier, or every one of them is older than the event log keeps.
    field(:money_rebuild_since, :string, source: :moneyRebuildSince)
    # ADR-128 §12 conflict rule: automatic linking has stopped for this
    # identity because evidence pointed at two candidates, or contradicted an
    # existing link. Stored HERE rather than on the suggestion rows, because a
    # halt on auto-linking that a recompute or a restart clears is not a halt.
    field(:suspended_at, LangWatch.PrismaDateTime, source: :suspendedAt)
    # What tripped the suspension, for the human who reviews it.
    field(:suspended_reason, :string, source: :suspendedReason)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end