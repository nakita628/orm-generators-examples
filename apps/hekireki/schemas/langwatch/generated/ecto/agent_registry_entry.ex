defmodule LangWatch.AgentRegistryEntry do
  @moduledoc """
  One agent in the organisation's governance registry, from any source governance reads
  (enterprise/modules/governance/specs/agent-registry.feature). `entry` is the latest
  machine-readable snapshot; `version` grows by one on every material change.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          source_key: String.t(),
          source: String.t(),
          status: String.t(),
          version: integer(),
          fingerprint: String.t(),
          entry: map() | list() | String.t() | number() | boolean(),
          first_seen_at: DateTime.t(),
          last_seen_at: DateTime.t() | nil,
          removed_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "AgentRegistryEntry" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    # Where the agent came from and its id there, e.g. `langwatch:<agentId>`.
    field(:source_key, :string, source: :sourceKey)
    field(:source, :string)
    field(:status, :string)
    field(:version, :integer)
    field(:fingerprint, :string)
    field(:entry, LangWatch.PrismaJson)
    field(:first_seen_at, LangWatch.PrismaDateTime, source: :firstSeenAt)
    field(:last_seen_at, LangWatch.PrismaDateTime, source: :lastSeenAt)
    field(:removed_at, LangWatch.PrismaDateTime, source: :removedAt)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end