defmodule LangWatch.GovernanceTenantHistory do
  @moduledoc """
  Every ClickHouse `TenantId` this organization has ever written governance
  rows under, appended on first use and never pruned.

  The live `resolveGovProjectId` lookup cannot be this translation: it filters
  `archivedAt: null` and so returns null forever once someone archives the
  governance project, while the write path re-reads the same project by slug
  with no such filter and keeps landing rows under the old tenant. The
  organization would map to zero tenants on read and one on write - a
  permanent split-brain in which the cost screen reports "no governance data"
  and an erasure job erases nothing and reports success.

  Reads resolve against the whole history; erasure walks all of it, because
  personal data does not stop existing in a tenant that stopped being current.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          tenant_id: String.t(),
          first_used_at: DateTime.t(),
          last_used_at: DateTime.t()
        }

  schema "GovernanceTenantHistory" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    # A hidden governance project id this organization has written rows under.
    field(:tenant_id, :string, source: :tenantId)
    field(:first_used_at, LangWatch.PrismaDateTime, source: :firstUsedAt)
    field(:last_used_at, LangWatch.PrismaDateTime, source: :lastUsedAt)
  end
end