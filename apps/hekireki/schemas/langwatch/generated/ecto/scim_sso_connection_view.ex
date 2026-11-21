defmodule LangWatch.ScimSsoConnectionView do
  @moduledoc """
  SCIM's own copy of an organization's SSO connections: a peer fold over
  identity's connection facts (ARCHITECTURE §9), folded with identity-contract's
  reducer, one row per connection. SCIM reads this instead of asking identity.
  Rebuilt from identity's event log by a projection replay, so no relation:
  a foreign key would make a replay depend on state the log does not carry.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          folded: map(),
          applied_event_ids: [String.t()],
          projection_version: String.t(),
          created_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "ScimSsoConnectionView" do
    # The aggregate: identity's connection id.
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    # The connection as the reducer folded it (SsoConnectionState + LastEventOccurredAt).
    field(:folded, :map)
    # The event ids of the last delivery folded in, so a redelivery folds once.
    field(:applied_event_ids, {:array, :string}, source: :appliedEventIds)
    # A row of another version reads as undecodable and is re-folded from the log.
    field(:projection_version, :string, source: :projectionVersion)
    # Event-derived, so a replayed row equals the row it rebuilds.
    field(:created_at, LangWatch.PrismaDateTime, source: :createdAt)
    field(:updated_at, LangWatch.PrismaDateTime, source: :updatedAt)
  end
end