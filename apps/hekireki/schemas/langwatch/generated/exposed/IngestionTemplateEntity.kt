package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class IngestionTemplateEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, IngestionTemplateEntity>(IngestionTemplateTable)

    var organizationId by IngestionTemplateTable.organizationId

    /**
     * Stable slug - unique within (organizationId, slug) scope. Generated
     * from displayName + nanoid(6) on create for org-authored; canonical
     * values for platform-published (`claude_code` / `cursor` /
     * `claude_cowork` / `raw_otlp_advanced`).
     */
    var slug by IngestionTemplateTable.slug

    /**
     * Mirrors the slug for v1 platform-published rows. Kept as a separate
     * column so a future v2 may publish two templates that share an
     * upstream sourceType (e.g. claude_code-personal vs claude_code-team)
     * without a migration. String not enum - extensibility.
     */
    var sourceType by IngestionTemplateTable.sourceType
    var displayName by IngestionTemplateTable.displayName
    var description by IngestionTemplateTable.description

    /**
     * Icon source - single column, prefix-discriminated, mirrors
     * AiToolEntry.iconAsset:
     * `preset:<kind>`                 (built-in slot)
     * `data:image/svg+xml;base64,…`   (admin-uploaded for org-authored)
     * null                            → fall back to slug-default in UI
     */
    var iconAsset by IngestionTemplateTable.iconAsset

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
    var credentialSchema by IngestionTemplateTable.credentialSchema

    /**
     * Admin/platform-authored OTTL statements. Applied at receive time
     * AFTER the receiver re-stamps the binding's authoritative principal
     * + provenance keys (see `protectedTemplateAttributeKeys` 19-key
     * closed list). Templates write canonical OTTL outputs (gen_ai.*,
     * cost.usd, response.model) - they CANNOT write attribution or
     * provenance keys; the receiver guard restores those post-OTTL.
     */
    var ottlRules by IngestionTemplateTable.ottlRules

    /**
     * True for the v1 platform-published default templates
     * (organizationId IS NULL). Future org-authored rows write false.
     * Drives the admin UI "READ-ONLY view + Request a custom template"
     * surface - admins consume + opt-out only in v1.
     */
    var platformPublished by IngestionTemplateTable.platformPublished

    /**
     * Admin/platform enable. False = hidden from /me catalog. Existing
     * ingestion keys continue to land traces; new installs are blocked.
     */
    var enabled by IngestionTemplateTable.enabled
    var archivedAt by IngestionTemplateTable.archivedAt
    var createdAt by IngestionTemplateTable.createdAt
    var updatedAt by IngestionTemplateTable.updatedAt
    var createdById by IngestionTemplateTable.createdById
    var updatedById by IngestionTemplateTable.updatedById

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(IngestionTemplateTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
