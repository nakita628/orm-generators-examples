# Represents the top-level organizational hierarchy in Formbricks.
# Self-hosting instances typically have a single organization, while Formbricks Cloud
# supports multiple organizations with multi-tenancy.
#
# @property id - Unique identifier for the organization
# @property name - Display name of the organization
# @property memberships - User memberships within the organization
# @property workspaces - Collection of workspaces owned by the organization
# @property billing - JSON field containing billing information
# @property whitelabel - Whitelabel configuration for the organization
# @property isAISmartToolsEnabled - Controls access to AI smart tools (e.g. translations) that never touch collected data
# @property displayTimeZone - Optional IANA time zone used for human-facing response timestamps (exports, integrations) and for the calendar days behind the analytics date filters (survey summary, charts, dashboards); null means UTC
#
# whitelabel:: [OrganizationWhitelabel]
class Organization < ApplicationRecord
  self.table_name = "Organization"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :whitelabel, default: -> { JSON.parse("{}") }
  attribute :isAISmartToolsEnabled, default: false

  validates :name, presence: true

  has_one :billing, class_name: "OrganizationBilling", inverse_of: :organization, dependent: :destroy
  has_many :memberships, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :workspaces, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :invites, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :teams, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :api_keys, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :feedback_directories, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :retention_policies, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :retention_exemptions, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :retention_runs, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :retention_notices, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :mcp_grants, class_name: "McpGrantOrganization", foreign_key: "organizationId", inverse_of: :organization,
                        dependent: :destroy
end
