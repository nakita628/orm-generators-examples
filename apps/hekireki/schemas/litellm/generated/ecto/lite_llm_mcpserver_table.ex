defmodule LiteLLM.LiteLLMMCPServerTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :server_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          server_name: String.t() | nil,
          alias: String.t() | nil,
          description: String.t() | nil,
          instructions: String.t() | nil,
          url: String.t() | nil,
          spec_path: String.t() | nil,
          transport: String.t(),
          auth_type: String.t() | nil,
          credentials: map() | list() | String.t() | number() | boolean() | nil,
          created_by: String.t() | nil,
          updated_by: String.t() | nil,
          mcp_info: map() | list() | String.t() | number() | boolean() | nil,
          mcp_access_groups: [String.t()],
          allowed_tools: [String.t()],
          tool_name_to_display_name: map() | list() | String.t() | number() | boolean() | nil,
          tool_name_to_description: map() | list() | String.t() | number() | boolean() | nil,
          pinned_tools: map() | list() | String.t() | number() | boolean() | nil,
          extra_headers: [String.t()],
          static_headers: map() | list() | String.t() | number() | boolean() | nil,
          env_vars: map() | list() | String.t() | number() | boolean() | nil,
          status: String.t() | nil,
          last_health_check: DateTime.t() | nil,
          health_check_error: String.t() | nil,
          command: String.t() | nil,
          args: [String.t()],
          env: map() | list() | String.t() | number() | boolean() | nil,
          issuer: String.t() | nil,
          authorization_url: String.t() | nil,
          token_url: String.t() | nil,
          registration_url: String.t() | nil,
          oauth2_flow: String.t() | nil,
          token_exchange_endpoint: String.t() | nil,
          audience: String.t() | nil,
          subject_token_type: String.t() | nil,
          token_exchange_profile: String.t() | nil,
          allow_all_keys: boolean(),
          available_on_public_internet: boolean(),
          delegate_auth_to_upstream: boolean(),
          oauth_passthrough: boolean(),
          dcr_bridge: boolean() | nil,
          per_server_oauth_discovery: boolean(),
          is_byok: boolean(),
          byok_description: [String.t()],
          byok_api_key_help_url: String.t() | nil,
          source_url: String.t() | nil,
          timeout: float() | nil,
          max_concurrent_requests: integer() | nil,
          rpm: integer() | nil,
          events: map() | list() | String.t() | number() | boolean() | nil,
          approval_status: String.t() | nil,
          submitted_by: String.t() | nil,
          submitted_at: DateTime.t() | nil,
          reviewed_at: DateTime.t() | nil,
          review_notes: String.t() | nil,
          inserted_at: DateTime.t() | nil,
          updated_at: DateTime.t() | nil
        }

  schema "LiteLLM_MCPServerTable" do
    field(:server_name, :string)
    field(:alias, :string)
    field(:description, :string)
    field(:instructions, :string)
    field(:url, :string)
    field(:spec_path, :string)
    field(:transport, :string, default: "sse")
    field(:auth_type, :string)
    field(:credentials, LiteLLM.PrismaJson, default: %{})
    field(:created_by, :string)
    field(:updated_by, :string)
    field(:mcp_info, LiteLLM.PrismaJson, default: %{})
    field(:mcp_access_groups, {:array, :string})
    field(:allowed_tools, {:array, :string})
    field(:tool_name_to_display_name, LiteLLM.PrismaJson, default: %{})
    field(:tool_name_to_description, LiteLLM.PrismaJson, default: %{})
    field(:pinned_tools, LiteLLM.PrismaJson, default: %{})
    field(:extra_headers, {:array, :string})
    field(:static_headers, LiteLLM.PrismaJson, default: %{})
    field(:env_vars, LiteLLM.PrismaJson, default: [])
    field(:status, :string, default: "unknown")
    field(:last_health_check, LiteLLM.PrismaDateTime)
    field(:health_check_error, :string)
    field(:command, :string)
    field(:args, {:array, :string})
    field(:env, LiteLLM.PrismaJson, default: %{})
    field(:issuer, :string)
    field(:authorization_url, :string)
    field(:token_url, :string)
    field(:registration_url, :string)
    field(:oauth2_flow, :string)
    field(:token_exchange_endpoint, :string)
    field(:audience, :string)
    field(:subject_token_type, :string)
    field(:token_exchange_profile, :string)
    field(:allow_all_keys, :boolean, default: false)
    field(:available_on_public_internet, :boolean, default: true)
    field(:delegate_auth_to_upstream, :boolean, default: false)
    field(:oauth_passthrough, :boolean, default: false)
    field(:dcr_bridge, :boolean)
    field(:per_server_oauth_discovery, :boolean, default: false)
    field(:is_byok, :boolean, default: false)
    field(:byok_description, {:array, :string})
    field(:byok_api_key_help_url, :string)
    field(:source_url, :string)
    field(:timeout, :float)
    field(:max_concurrent_requests, :integer)
    field(:rpm, :integer)
    field(:events, LiteLLM.PrismaJson)
    field(:approval_status, :string, default: "active")
    field(:submitted_by, :string)
    field(:submitted_at, LiteLLM.PrismaDateTime)
    field(:reviewed_at, LiteLLM.PrismaDateTime)
    field(:review_notes, :string)
    timestamps(inserted_at_source: :created_at)
  end
end