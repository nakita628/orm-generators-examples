defmodule Formbricks.McpGrant do
  @moduledoc """
  Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
  its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
  a token is now bound to the resource it was approved for instead of to a static allow-list.
  `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
  What one user allows one MCP client to do (GLOSSARY.md "Grant", ADR-0041): its organization & workspace
  access and an access level per permission. One per OAuth connection, keyed by user and client rather
  than by `oauthConsent` row, because the provider can write duplicate consent rows for one connection
  and settles them later. Every MCP request reads it; without one the request is refused. The client's
  OAuth scopes stay on the consent and cap it.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          access: atom(),
          permissions: map() | list() | String.t() | number() | boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Formbricks.User.t() | nil,
          client: Formbricks.OauthClient.t() | nil,
          organizations: [Formbricks.McpGrantOrganization.t()],
          workspaces: [Formbricks.McpGrantWorkspace.t()]
        }

  schema "McpGrant" do
    field(:id, :string, primary_key: true)
    # Which organizations and workspaces the grant reaches (ENG-7153): All access, only the organizations
    # in `organizations` (and every workspace in them, later ones included), or only the workspaces in
    # `workspaces`. Enforced in SpiceDB through `can()` (ADR-0042); a value no row backs reaches nothing.
    field(:access, Ecto.Enum, values: [:all, :organizations, :workspaces], default: :all)
    # The access level per permission of `apps/web/modules/mcp/permissions.ts`: `{ "surveys": "write" }`.
    # A permission that is absent has No access.
    field(:permissions, Formbricks.PrismaJson, default: %{})
    field(:user_id, :string, source: :userId)
    field(:client_id, :string, source: :clientId)
    belongs_to(:user, Formbricks.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:client, Formbricks.OauthClient, foreign_key: :client_id, define_field: false, type: :string, references: :client_id)
    # The selections behind `organizations` and `workspaces` access (ENG-7153).
    has_many(:organizations, Formbricks.McpGrantOrganization, foreign_key: :grant_id, references: :id)
    has_many(:workspaces, Formbricks.McpGrantWorkspace, foreign_key: :grant_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end