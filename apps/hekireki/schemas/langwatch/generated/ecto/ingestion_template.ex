defmodule LangWatch.IngestionTemplate do
  @moduledoc """
  Admin/platform-curated catalog row for upstream OTLP normalization.
  Owns the OTTL transform that reshapes a specific upstream tool's
  span shape into canonical gen_ai conventions. Users connect via an
  ingestion key (an ingest-only ApiKey carrying ingestionTemplateId) to
  land traces in their personal project.

  `organizationId` NULL = platform-published default (always-visible
  across every org). NOT NULL = org-authored (visible only within
  that org). v1 ships only platform-published; org-authored authoring
  UI lands v2 (the column ships now to avoid a future migration).

  Spec: specs/ai-gateway/governance/ingestion-templates-catalog.feature
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          slug: String.t(),
          source_type: String.t(),
          display_name: String.t(),
          description: String.t() | nil,
          icon_asset: String.t() | nil,
          credential_schema: String.t() | nil,
          ottl_rules: String.t(),
          platform_published: boolean(),
          enabled: boolean(),
          archived_at: DateTime.t() | nil,
          created_by_id: String.t() | nil,
          updated_by_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: LangWatch.Organization.t() | nil
        }

  schema "IngestionTemplate" do
    field(:id, :string, primary_key: true)
    # Stable slug - unique within (organizationId, slug) scope. Generated
    # from displayName + nanoid(6) on create for org-authored; canonical
    # values for platform-published (`claude_code` / `cursor` /
    # `claude_cowork` / `raw_otlp_advanced`).
    field(:slug, :string)
    # Mirrors the slug for v1 platform-published rows. Kept as a separate
    # column so a future v2 may publish two templates that share an
    # upstream sourceType (e.g. claude_code-personal vs claude_code-team)
    # without a migration. String not enum - extensibility.
    field(:source_type, :string, source: :sourceType)
    field(:display_name, :string, source: :displayName)
    field(:description, :string)
    # Icon source - single column, prefix-discriminated, mirrors
    # AiToolEntry.iconAsset:
    # `preset:<kind>`                 (built-in slot)
    # `data:image/svg+xml;base64,…`   (admin-uploaded for org-authored)
    # null                            → fall back to slug-default in UI
    field(:icon_asset, :string, source: :iconAsset)
    # Discriminator for what (if any) extra credential metadata the
    # install drawer captures from the user:
    # null              → otlp_token-only (drawer auto-issues the
    # ingestion key, no input form)
    # `static_api_key`  → drawer captures the user's upstream tool API key
    # `agent_id`        → drawer captures an agent identifier (e.g.
    # Cursor agentId)
    # String not enum so v2 shapes (oauth_redirect / s3_bucket) ship
    # without a migration.
    field(:credential_schema, :string, source: :credentialSchema)
    # Admin/platform-authored OTTL statements. Applied at receive time
    # AFTER the receiver re-stamps the binding's authoritative principal
    # + provenance keys (see `protectedTemplateAttributeKeys` 19-key
    # closed list). Templates write canonical OTTL outputs (gen_ai.*,
    # cost.usd, response.model) - they CANNOT write attribution or
    # provenance keys; the receiver guard restores those post-OTTL.
    field(:ottl_rules, :string, default: "", source: :ottlRules)
    # True for the v1 platform-published default templates
    # (organizationId IS NULL). Future org-authored rows write false.
    # Drives the admin UI "READ-ONLY view + Request a custom template"
    # surface - admins consume + opt-out only in v1.
    field(:platform_published, :boolean, default: false, source: :platformPublished)
    # Admin/platform enable. False = hidden from /me catalog. Existing
    # ingestion keys continue to land traces; new installs are blocked.
    field(:enabled, :boolean, default: true)
    field(:archived_at, LangWatch.PrismaDateTime, source: :archivedAt)
    field(:created_by_id, :string, source: :createdById)
    field(:updated_by_id, :string, source: :updatedById)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end