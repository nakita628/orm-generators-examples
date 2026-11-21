# Org-scoped catalog of AI tools surfaced on the personal home portal
# (/me) and managed by org admins at /governance/tool-catalog.
#
# Three first-class tile types via `type` discriminator:
# - "coding_assistant" - click expands to setup helper (e.g. `langwatch claude`)
# - "model_provider"   - click expands to inline personal-VK creation
# - "external_tool"    - click expands to admin-attached markdown + external link
#
# Scoping: organization-default entries can be OVERRIDDEN at the team
# tier (same `slug` + `scope=team` + matching `scopeId`) so admins can
# hide or re-skin a globally-listed tool for a specific team. The
# user-facing `list` endpoint resolves this hierarchy at read time.
#
# `config` is a discriminated JSON union per `type`:
# coding_assistant: { setupCommand, setupDocsUrl, helperText? }
# model_provider:   { providerKey, suggestedRoutingPolicyId?, defaultLabel?, projectSuggestionText? }
# external_tool:    { descriptionMarkdown, linkUrl, ctaLabel? }
#
# Spec: specs/ai-governance/personal-portal/tool-catalog-*.feature
#
# scope:: "organization" | "team". String not enum so new scopes ship
#   without migrations.
# scopeId:: organizationId or teamId, narrowing what scope this entry applies
#   to. Always equals organizationId when scope=organization.
# type:: "coding_assistant" | "model_provider" | "external_tool".
#   Discriminator drives the per-tile UX shape and validates
#   the `config` JSON column at the app layer.
# slug:: Stable slug for team-overrides-org resolution (a team-scoped
#   entry with the same slug as an org-scoped one hides the org
#   default for users in that team). Server-owned: generated from
#   displayName + nanoid(6) on create, immutable on update.
# iconKey:: @deprecated Use `iconAsset` instead. Kept for back-compat with
#   pre-refactor rows; new writes should populate iconAsset.
# iconAsset:: Icon source - single column, prefix-discriminated:
#   "preset:claude_code" / "preset:codex" / ...   (built-in)
#   "data:image/svg+xml;base64,..."               (admin-uploaded)
#   null                                          → fall back to
#   type-default in UI
# order:: Admin-curated sort within the user-facing portal grid.
# enabled:: When false the entry is hidden from the user-facing list but
#   remains in the admin list (so admins can re-enable without
#   losing per-entry config). Disabled is "soft-hide", archived is
#   "soft-delete" (preserves audit trail + scope binding history).
# config:: Per-type config JSON. Validated via zod discriminated union
#   at the service layer (see AiToolEntryService).
# teams:: @deprecated Multi-team scope. Superseded by `departments` - the
#   "Visible to" model is now WHOLE ORGANIZATION or a set of
#   DEPARTMENTS. Rows are no longer written here (the table is kept
#   for now to avoid a destructive drop); readers prefer `departments`.
# departments:: Department scope. Empty = org-wide (visible to every member).
#   Non-empty = visible only to members whose `departmentId` is in
#   this set. Mirrored onto the legacy scope/scopeId pair for
#   back-compat reads (scope='organization' when empty,
#   scope='department' otherwise).
class AiToolEntry < ApplicationRecord
  self.table_name = "AiToolEntry"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :order, default: 0
  attribute :enabled, default: true
  attribute :config, default: -> { JSON.parse("{}") }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :scope, presence: true
  validates :scopeId, presence: true
  validates :type, presence: true
  validates :displayName, presence: true
  validates :slug, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :ai_tool_entries
  has_many :teams, class_name: "AiToolEntryTeam", foreign_key: "entryId", inverse_of: :entry, dependent: :destroy
  has_many :departments, class_name: "AiToolEntryDepartment", foreign_key: "entryId", inverse_of: :entry,
                         dependent: :destroy
end
