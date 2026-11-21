package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Org-scoped catalog of AI tools surfaced on the personal home portal
 * (/me) and managed by org admins at /governance/tool-catalog.
 *
 * Three first-class tile types via `type` discriminator:
 * - "coding_assistant" - click expands to setup helper (e.g. `langwatch claude`)
 * - "model_provider"   - click expands to inline personal-VK creation
 * - "external_tool"    - click expands to admin-attached markdown + external link
 *
 * Scoping: organization-default entries can be OVERRIDDEN at the team
 * tier (same `slug` + `scope=team` + matching `scopeId`) so admins can
 * hide or re-skin a globally-listed tool for a specific team. The
 * user-facing `list` endpoint resolves this hierarchy at read time.
 *
 * `config` is a discriminated JSON union per `type`:
 * coding_assistant: { setupCommand, setupDocsUrl, helperText? }
 * model_provider:   { providerKey, suggestedRoutingPolicyId?, defaultLabel?, projectSuggestionText? }
 * external_tool:    { descriptionMarkdown, linkUrl, ctaLabel? }
 *
 * Spec: specs/ai-governance/personal-portal/tool-catalog-*.feature
 */
object AiToolEntryTable : IdTable<String>("\"AiToolEntry\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")

    /**
     * "organization" | "team". String not enum so new scopes ship
     * without migrations.
     */
    val scope = text("scope")

    /**
     * organizationId or teamId, narrowing what scope this entry applies
     * to. Always equals organizationId when scope=organization.
     */
    val scopeId = text("scopeId")

    /**
     * "coding_assistant" | "model_provider" | "external_tool".
     * Discriminator drives the per-tile UX shape and validates
     * the `config` JSON column at the app layer.
     */
    val type = text("type")
    val displayName = text("displayName")

    /**
     * Stable slug for team-overrides-org resolution (a team-scoped
     * entry with the same slug as an org-scoped one hides the org
     * default for users in that team). Server-owned: generated from
     * displayName + nanoid(6) on create, immutable on update.
     */
    val slug = text("slug")

    /**
     * @deprecated Use `iconAsset` instead. Kept for back-compat with
     * pre-refactor rows; new writes should populate iconAsset.
     */
    val iconKey = text("iconKey").nullable()

    /**
     * Icon source - single column, prefix-discriminated:
     * "preset:claude_code" / "preset:codex" / ...   (built-in)
     * "data:image/svg+xml;base64,..."               (admin-uploaded)
     * null                                          → fall back to
     * type-default in UI
     */
    val iconAsset = text("iconAsset").nullable()

    /** Admin-curated sort within the user-facing portal grid. */
    val order = integer("order").default(0)

    /**
     * When false the entry is hidden from the user-facing list but
     * remains in the admin list (so admins can re-enable without
     * losing per-entry config). Disabled is "soft-hide", archived is
     * "soft-delete" (preserves audit trail + scope binding history).
     */
    val enabled = bool("enabled").default(true)

    /**
     * Per-type config JSON. Validated via zod discriminated union
     * at the service layer (see AiToolEntryService).
     */
    val config = jsonb<String>("config", { it }, { it }).default("{}")
    val archivedAt = pgTimestamp("archivedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val createdById = text("createdById").nullable()
    val updatedById = text("updatedById").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"AiToolEntry_pkey\"")

    init {
        index("\"AiToolEntry_organizationId_scope_scopeId_idx\"", false, organizationId, scope, scopeId)
        index("\"AiToolEntry_organizationId_enabled_archivedAt_idx\"", false, organizationId, enabled, archivedAt)
        index("\"AiToolEntry_organizationId_type_idx\"", false, organizationId, type)
    }
}
