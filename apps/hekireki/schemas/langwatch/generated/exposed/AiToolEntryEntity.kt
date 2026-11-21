package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class AiToolEntryEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AiToolEntryEntity>(AiToolEntryTable)

    var organizationId by AiToolEntryTable.organizationId

    /**
     * "organization" | "team". String not enum so new scopes ship
     * without migrations.
     */
    var scope by AiToolEntryTable.scope

    /**
     * organizationId or teamId, narrowing what scope this entry applies
     * to. Always equals organizationId when scope=organization.
     */
    var scopeId by AiToolEntryTable.scopeId

    /**
     * "coding_assistant" | "model_provider" | "external_tool".
     * Discriminator drives the per-tile UX shape and validates
     * the `config` JSON column at the app layer.
     */
    var type by AiToolEntryTable.type
    var displayName by AiToolEntryTable.displayName

    /**
     * Stable slug for team-overrides-org resolution (a team-scoped
     * entry with the same slug as an org-scoped one hides the org
     * default for users in that team). Server-owned: generated from
     * displayName + nanoid(6) on create, immutable on update.
     */
    var slug by AiToolEntryTable.slug

    /**
     * @deprecated Use `iconAsset` instead. Kept for back-compat with
     * pre-refactor rows; new writes should populate iconAsset.
     */
    var iconKey by AiToolEntryTable.iconKey

    /**
     * Icon source - single column, prefix-discriminated:
     * "preset:claude_code" / "preset:codex" / ...   (built-in)
     * "data:image/svg+xml;base64,..."               (admin-uploaded)
     * null                                          → fall back to
     * type-default in UI
     */
    var iconAsset by AiToolEntryTable.iconAsset

    /** Admin-curated sort within the user-facing portal grid. */
    var order by AiToolEntryTable.order

    /**
     * When false the entry is hidden from the user-facing list but
     * remains in the admin list (so admins can re-enable without
     * losing per-entry config). Disabled is "soft-hide", archived is
     * "soft-delete" (preserves audit trail + scope binding history).
     */
    var enabled by AiToolEntryTable.enabled

    /**
     * Per-type config JSON. Validated via zod discriminated union
     * at the service layer (see AiToolEntryService).
     */
    var config by AiToolEntryTable.config
    var archivedAt by AiToolEntryTable.archivedAt
    var createdAt by AiToolEntryTable.createdAt
    var updatedAt by AiToolEntryTable.updatedAt
    var createdById by AiToolEntryTable.createdById
    var updatedById by AiToolEntryTable.updatedById

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AiToolEntryTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
