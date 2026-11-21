#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// Org-scoped catalog of AI tools surfaced on the personal home portal
/// (/me) and managed by org admins at /governance/tool-catalog.
///
/// Three first-class tile types via `type` discriminator:
/// - "coding_assistant" - click expands to setup helper (e.g. `langwatch claude`)
/// - "model_provider"   - click expands to inline personal-VK creation
/// - "external_tool"    - click expands to admin-attached markdown + external link
///
/// Scoping: organization-default entries can be OVERRIDDEN at the team
/// tier (same `slug` + `scope=team` + matching `scopeId`) so admins can
/// hide or re-skin a globally-listed tool for a specific team. The
/// user-facing `list` endpoint resolves this hierarchy at read time.
///
/// `config` is a discriminated JSON union per `type`:
/// coding_assistant: { setupCommand, setupDocsUrl, helperText? }
/// model_provider:   { providerKey, suggestedRoutingPolicyId?, defaultLabel?, projectSuggestionText? }
/// external_tool:    { descriptionMarkdown, linkUrl, ctaLabel? }
///
/// Spec: specs/ai-governance/personal-portal/tool-catalog-*.feature
/// </summary>
public partial class AiToolEntry
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    /// <summary>
    /// "organization" | "team". String not enum so new scopes ship
    /// without migrations.
    /// </summary>
    public string Scope { get; set; } = null!;

    /// <summary>
    /// organizationId or teamId, narrowing what scope this entry applies
    /// to. Always equals organizationId when scope=organization.
    /// </summary>
    public string ScopeId { get; set; } = null!;

    /// <summary>
    /// "coding_assistant" | "model_provider" | "external_tool".
    /// Discriminator drives the per-tile UX shape and validates
    /// the `config` JSON column at the app layer.
    /// </summary>
    public string Type { get; set; } = null!;

    public string DisplayName { get; set; } = null!;

    /// <summary>
    /// Stable slug for team-overrides-org resolution (a team-scoped
    /// entry with the same slug as an org-scoped one hides the org
    /// default for users in that team). Server-owned: generated from
    /// displayName + nanoid(6) on create, immutable on update.
    /// </summary>
    public string Slug { get; set; } = null!;

    /// <summary>
    /// @deprecated Use `iconAsset` instead. Kept for back-compat with
    /// pre-refactor rows; new writes should populate iconAsset.
    /// </summary>
    public string? IconKey { get; set; }

    /// <summary>
    /// Icon source - single column, prefix-discriminated:
    /// "preset:claude_code" / "preset:codex" / ...   (built-in)
    /// "data:image/svg+xml;base64,..."               (admin-uploaded)
    /// null                                          → fall back to
    /// type-default in UI
    /// </summary>
    public string? IconAsset { get; set; }

    /// <summary>
    /// Admin-curated sort within the user-facing portal grid.
    /// </summary>
    public int Order { get; set; } = 0;

    /// <summary>
    /// When false the entry is hidden from the user-facing list but
    /// remains in the admin list (so admins can re-enable without
    /// losing per-entry config). Disabled is "soft-hide", archived is
    /// "soft-delete" (preserves audit trail + scope binding history).
    /// </summary>
    public bool Enabled { get; set; } = true;

    /// <summary>
    /// Per-type config JSON. Validated via zod discriminated union
    /// at the service layer (see AiToolEntryService).
    /// </summary>
    public string Config { get; set; } = "{}";

    public DateTime? ArchivedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string? CreatedById { get; set; }

    public string? UpdatedById { get; set; }

    /// <summary>
    /// @deprecated Multi-team scope. Superseded by `departments` - the
    /// "Visible to" model is now WHOLE ORGANIZATION or a set of
    /// DEPARTMENTS. Rows are no longer written here (the table is kept
    /// for now to avoid a destructive drop); readers prefer `departments`.
    /// </summary>
    public virtual ICollection<AiToolEntryTeam> Teams { get; set; } = new List<AiToolEntryTeam>();

    /// <summary>
    /// Department scope. Empty = org-wide (visible to every member).
    /// Non-empty = visible only to members whose `departmentId` is in
    /// this set. Mirrored onto the legacy scope/scopeId pair for
    /// back-compat reads (scope='organization' when empty,
    /// scope='department' otherwise).
    /// </summary>
    public virtual ICollection<AiToolEntryDepartment> Departments { get; set; } = new List<AiToolEntryDepartment>();
}
