defmodule LangWatch.AiToolEntry do
  @moduledoc """
  Org-scoped catalog of AI tools surfaced on the personal home portal
  (/me) and managed by org admins at /governance/tool-catalog.

  Three first-class tile types via `type` discriminator:
  - "coding_assistant" - click expands to setup helper (e.g. `langwatch claude`)
  - "model_provider"   - click expands to inline personal-VK creation
  - "external_tool"    - click expands to admin-attached markdown + external link

  Scoping: organization-default entries can be OVERRIDDEN at the team
  tier (same `slug` + `scope=team` + matching `scopeId`) so admins can
  hide or re-skin a globally-listed tool for a specific team. The
  user-facing `list` endpoint resolves this hierarchy at read time.

  `config` is a discriminated JSON union per `type`:
  coding_assistant: { setupCommand, setupDocsUrl, helperText? }
  model_provider:   { providerKey, suggestedRoutingPolicyId?, defaultLabel?, projectSuggestionText? }
  external_tool:    { descriptionMarkdown, linkUrl, ctaLabel? }

  Spec: specs/ai-governance/personal-portal/tool-catalog-*.feature
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          scope: String.t(),
          scope_id: String.t(),
          type: String.t(),
          display_name: String.t(),
          slug: String.t(),
          icon_key: String.t() | nil,
          icon_asset: String.t() | nil,
          order: integer(),
          enabled: boolean(),
          config: map() | list() | String.t() | number() | boolean(),
          archived_at: DateTime.t() | nil,
          created_by_id: String.t() | nil,
          updated_by_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: LangWatch.Organization.t() | nil,
          teams: [LangWatch.AiToolEntryTeam.t()],
          departments: [LangWatch.AiToolEntryDepartment.t()]
        }

  schema "AiToolEntry" do
    field(:id, :string, primary_key: true)
    # "organization" | "team". String not enum so new scopes ship
    # without migrations.
    field(:scope, :string)
    # organizationId or teamId, narrowing what scope this entry applies
    # to. Always equals organizationId when scope=organization.
    field(:scope_id, :string, source: :scopeId)
    # "coding_assistant" | "model_provider" | "external_tool".
    # Discriminator drives the per-tile UX shape and validates
    # the `config` JSON column at the app layer.
    field(:type, :string)
    field(:display_name, :string, source: :displayName)
    # Stable slug for team-overrides-org resolution (a team-scoped
    # entry with the same slug as an org-scoped one hides the org
    # default for users in that team). Server-owned: generated from
    # displayName + nanoid(6) on create, immutable on update.
    field(:slug, :string)
    # @deprecated Use `iconAsset` instead. Kept for back-compat with
    # pre-refactor rows; new writes should populate iconAsset.
    field(:icon_key, :string, source: :iconKey)
    # Icon source - single column, prefix-discriminated:
    # "preset:claude_code" / "preset:codex" / ...   (built-in)
    # "data:image/svg+xml;base64,..."               (admin-uploaded)
    # null                                          → fall back to
    # type-default in UI
    field(:icon_asset, :string, source: :iconAsset)
    # Admin-curated sort within the user-facing portal grid.
    field(:order, :integer, default: 0)
    # When false the entry is hidden from the user-facing list but
    # remains in the admin list (so admins can re-enable without
    # losing per-entry config). Disabled is "soft-hide", archived is
    # "soft-delete" (preserves audit trail + scope binding history).
    field(:enabled, :boolean, default: true)
    # Per-type config JSON. Validated via zod discriminated union
    # at the service layer (see AiToolEntryService).
    field(:config, LangWatch.PrismaJson, default: %{})
    field(:archived_at, LangWatch.PrismaDateTime, source: :archivedAt)
    field(:created_by_id, :string, source: :createdById)
    field(:updated_by_id, :string, source: :updatedById)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    # @deprecated Multi-team scope. Superseded by `departments` - the
    # "Visible to" model is now WHOLE ORGANIZATION or a set of
    # DEPARTMENTS. Rows are no longer written here (the table is kept
    # for now to avoid a destructive drop); readers prefer `departments`.
    has_many(:teams, LangWatch.AiToolEntryTeam, foreign_key: :entry_id, references: :id)
    # Department scope. Empty = org-wide (visible to every member).
    # Non-empty = visible only to members whose `departmentId` is in
    # this set. Mirrored onto the legacy scope/scopeId pair for
    # back-compat reads (scope='organization' when empty,
    # scope='department' otherwise).
    has_many(:departments, LangWatch.AiToolEntryDepartment, foreign_key: :entry_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end