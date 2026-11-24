defmodule Trigger.OrganizationDataStore do
  @moduledoc """
  Defines org-scoped data store overrides (e.g. dedicated ClickHouse for HIPAA orgs).
  Multiple organizations can share a single data store row via organizationIds.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          key: String.t(),
          organization_ids: [String.t()],
          kind: atom(),
          config: map() | list() | String.t() | number() | boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "OrganizationDataStore" do
    field(:id, :string, primary_key: true)
    # Human-readable unique key (e.g. "hipaa-clickhouse-us-east")
    field(:key, :string)
    # Organization IDs that use this data store
    field(:organization_ids, {:array, :string}, source: :organizationIds)
    field(:kind, Ecto.Enum, values: [:CLICKHOUSE])
    # Versioned config JSON. Structure is discriminated by the top-level `version` field.
    field(:config, Trigger.PrismaJson)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end