package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Admin/platform-curated catalog row for upstream OTLP normalization.
 * Owns the OTTL transform that reshapes a specific upstream tool's
 * span shape into canonical gen_ai conventions. Users connect via an
 * ingestion key (an ingest-only ApiKey carrying ingestionTemplateId) to
 * land traces in their personal project.
 *
 * `organizationId` NULL = platform-published default (always-visible
 * across every org). NOT NULL = org-authored (visible only within
 * that org). v1 ships only platform-published; org-authored authoring
 * UI lands v2 (the column ships now to avoid a future migration).
 *
 * Spec: specs/ai-gateway/governance/ingestion-templates-catalog.feature
 */
object IngestionTemplateTable : IdTable<String>("\"IngestionTemplate\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId").nullable()

    /**
     * Stable slug - unique within (organizationId, slug) scope. Generated
     * from displayName + nanoid(6) on create for org-authored; canonical
     * values for platform-published (`claude_code` / `cursor` /
     * `claude_cowork` / `raw_otlp_advanced`).
     */
    val slug = text("slug")

    /**
     * Mirrors the slug for v1 platform-published rows. Kept as a separate
     * column so a future v2 may publish two templates that share an
     * upstream sourceType (e.g. claude_code-personal vs claude_code-team)
     * without a migration. String not enum - extensibility.
     */
    val sourceType = text("sourceType")
    val displayName = text("displayName")
    val description = text("description").nullable()

    /**
     * Icon source - single column, prefix-discriminated, mirrors
     * AiToolEntry.iconAsset:
     * `preset:<kind>`                 (built-in slot)
     * `data:image/svg+xml;base64,…`   (admin-uploaded for org-authored)
     * null                            → fall back to slug-default in UI
     */
    val iconAsset = text("iconAsset").nullable()

    /**
     * Discriminator for what (if any) extra credential metadata the
     * install drawer captures from the user:
     * null              → otlp_token-only (drawer auto-issues the
     * ingestion key, no input form)
     * `static_api_key`  → drawer captures the user's upstream tool API key
     * `agent_id`        → drawer captures an agent identifier (e.g.
     * Cursor agentId)
     * String not enum so v2 shapes (oauth_redirect / s3_bucket) ship
     * without a migration.
     */
    val credentialSchema = text("credentialSchema").nullable()

    /**
     * Admin/platform-authored OTTL statements. Applied at receive time
     * AFTER the receiver re-stamps the binding's authoritative principal
     * + provenance keys (see `protectedTemplateAttributeKeys` 19-key
     * closed list). Templates write canonical OTTL outputs (gen_ai.*,
     * cost.usd, response.model) - they CANNOT write attribution or
     * provenance keys; the receiver guard restores those post-OTTL.
     */
    val ottlRules = text("ottlRules").default("")

    /**
     * True for the v1 platform-published default templates
     * (organizationId IS NULL). Future org-authored rows write false.
     * Drives the admin UI "READ-ONLY view + Request a custom template"
     * surface - admins consume + opt-out only in v1.
     */
    val platformPublished = bool("platformPublished").default(false)

    /**
     * Admin/platform enable. False = hidden from /me catalog. Existing
     * ingestion keys continue to land traces; new installs are blocked.
     */
    val enabled = bool("enabled").default(true)
    val archivedAt = pgTimestamp("archivedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val createdById = text("createdById").nullable()
    val updatedById = text("updatedById").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"IngestionTemplate_pkey\"")

    init {
        index("\"IngestionTemplate_organizationId_archivedAt_idx\"", false, organizationId, archivedAt)
        index("\"IngestionTemplate_sourceType_idx\"", false, sourceType)
        uniqueIndex("\"IngestionTemplate_organizationId_slug_key\"", organizationId, slug)
    }
}
