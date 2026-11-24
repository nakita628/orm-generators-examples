defmodule Formbricks.Organization do
  @moduledoc """
  Represents the top-level organizational hierarchy in Formbricks.
  Self-hosting instances typically have a single organization, while Formbricks Cloud
  supports multiple organizations with multi-tenancy.

  @property id - Unique identifier for the organization
  @property name - Display name of the organization
  @property memberships - User memberships within the organization
  @property workspaces - Collection of workspaces owned by the organization
  @property billing - JSON field containing billing information
  @property whitelabel - Whitelabel configuration for the organization
  @property isAISmartToolsEnabled - Controls access to AI smart tools (e.g. translations) that never touch collected data
  @property displayTimeZone - Optional IANA time zone used for human-facing response timestamps (exports, integrations) and for the calendar days behind the analytics date filters (survey summary, charts, dashboards); null means UTC
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          whitelabel: map() | list() | String.t() | number() | boolean(),
          is_aismart_tools_enabled: boolean(),
          display_time_zone: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          billing: Formbricks.OrganizationBilling.t() | nil,
          memberships: [Formbricks.Membership.t()],
          workspaces: [Formbricks.Workspace.t()],
          invites: [Formbricks.Invite.t()],
          teams: [Formbricks.Team.t()],
          api_keys: [Formbricks.ApiKey.t()],
          feedback_directories: [Formbricks.FeedbackDirectory.t()],
          retention_policies: [Formbricks.RetentionPolicy.t()],
          retention_exemptions: [Formbricks.RetentionExemption.t()],
          retention_runs: [Formbricks.RetentionRun.t()],
          retention_notices: [Formbricks.RetentionNotice.t()]
        }

  schema "Organization" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    # [OrganizationWhitelabel]
    field(:whitelabel, Formbricks.PrismaJson, default: %{})
    field(:is_aismart_tools_enabled, :boolean, default: false, source: :isAISmartToolsEnabled)
    field(:display_time_zone, :string, source: :displayTimeZone)
    has_one(:billing, Formbricks.OrganizationBilling, foreign_key: :organization_id, references: :id)
    has_many(:memberships, Formbricks.Membership, foreign_key: :organization_id, references: :id)
    has_many(:workspaces, Formbricks.Workspace, foreign_key: :organization_id, references: :id)
    has_many(:invites, Formbricks.Invite, foreign_key: :organization_id, references: :id)
    has_many(:teams, Formbricks.Team, foreign_key: :organization_id, references: :id)
    has_many(:api_keys, Formbricks.ApiKey, foreign_key: :organization_id, references: :id)
    has_many(:feedback_directories, Formbricks.FeedbackDirectory, foreign_key: :organization_id, references: :id)
    has_many(:retention_policies, Formbricks.RetentionPolicy, foreign_key: :organization_id, references: :id)
    has_many(:retention_exemptions, Formbricks.RetentionExemption, foreign_key: :organization_id, references: :id)
    has_many(:retention_runs, Formbricks.RetentionRun, foreign_key: :organization_id, references: :id)
    has_many(:retention_notices, Formbricks.RetentionNotice, foreign_key: :organization_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end