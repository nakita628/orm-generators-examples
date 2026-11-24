# Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
# its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
# a token is now bound to the resource it was approved for instead of to a static allow-list.
# `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
# What one user allows one MCP client to do (GLOSSARY.md "Grant", ADR-0041): its organization & workspace
# access and an access level per permission. One per OAuth connection, keyed by user and client rather
# than by `oauthConsent` row, because the provider can write duplicate consent rows for one connection
# and settles them later. Every MCP request reads it; without one the request is refused. The client's
# OAuth scopes stay on the consent and cap it.
#
# access:: Which organizations and workspaces the grant reaches (ENG-7153): All access, only the organizations
#   in `organizations` (and every workspace in them, later ones included), or only the workspaces in
#   `workspaces`. Enforced in SpiceDB through `can()` (ADR-0042); a value no row backs reaches nothing.
# permissions:: The access level per permission of `apps/web/modules/mcp/permissions.ts`: `{ "surveys": "write" }`.
#   A permission that is absent has No access.
# organizations:: The selections behind `organizations` and `workspaces` access (ENG-7153).
class McpGrant < ApplicationRecord
  self.table_name = "McpGrant"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :permissions, default: -> { JSON.parse("{}") }

  enum :access, { all: "all", organizations: "organizations", workspaces: "workspaces" }, default: :all, prefix: true,
                                                                                          validate: true

  validates :userId, uniqueness: { scope: :clientId }

  belongs_to :user, foreign_key: "userId", inverse_of: :mcp_grants
  belongs_to :client, class_name: "OauthClient", foreign_key: "clientId", primary_key: "clientId",
                      inverse_of: :mcp_grants
  has_many :organizations, class_name: "McpGrantOrganization", foreign_key: "grantId", inverse_of: :grant,
                           dependent: :destroy
  has_many :workspaces, class_name: "McpGrantWorkspace", foreign_key: "grantId", inverse_of: :grant, dependent: :destroy
end
