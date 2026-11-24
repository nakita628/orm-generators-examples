defmodule LangWatch.GatewayA2ATarget do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          project_id: String.t() | nil,
          name: String.t(),
          agent_name: String.t(),
          card_url: String.t(),
          upstream_url: String.t(),
          encrypted_upstream_headers: String.t() | nil,
          restrict_to_allow_list: boolean(),
          allowed_virtual_key_ids: [String.t()],
          blocked: boolean(),
          block_reason: String.t() | nil,
          created_by_id: String.t(),
          archived_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "GatewayA2ATarget" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:project_id, :string, source: :projectId)
    field(:name, :string)
    field(:agent_name, :string, source: :agentName)
    field(:card_url, :string, source: :cardUrl)
    field(:upstream_url, :string, source: :upstreamUrl)
    field(:encrypted_upstream_headers, :string, source: :encryptedUpstreamHeaders)
    field(:restrict_to_allow_list, :boolean, default: false, source: :restrictToAllowList)
    field(:allowed_virtual_key_ids, {:array, :string}, source: :allowedVirtualKeyIds)
    field(:blocked, :boolean, default: false)
    field(:block_reason, :string, source: :blockReason)
    field(:created_by_id, :string, source: :createdById)
    field(:archived_at, LangWatch.PrismaDateTime, source: :archivedAt)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end