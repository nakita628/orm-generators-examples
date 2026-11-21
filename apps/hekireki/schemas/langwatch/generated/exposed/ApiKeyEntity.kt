package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class ApiKeyEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ApiKeyEntity>(ApiKeyTable)

    var name by ApiKeyTable.name
    var description by ApiKeyTable.description
    var lookupId by ApiKeyTable.lookupId
    var hashedSecret by ApiKeyTable.hashedSecret
    var permissionMode by ApiKeyTable.permissionMode
    var userId by ApiKeyTable.userId
    var createdByUserId by ApiKeyTable.createdByUserId
    var organizationId by ApiKeyTable.organizationId
    var expiresAt by ApiKeyTable.expiresAt
    var lastUsedAt by ApiKeyTable.lastUsedAt
    var revokedAt by ApiKeyTable.revokedAt

    /**
     * Why the key was revoked, set with `revokedAt`: "user" when a person
     * revoked it (the API-keys page, the REST API, the devices tab), "rotation"
     * when a mint replaced it, "session" when the CLI login key it was parented
     * to was revoked, "expired" when that session ran out, "offboarded" when the
     * person's membership of the organization ended. "cap" survives on rows the
     * retired per-tool cap revoked. The CLI reads it to decide whether a dead
     * ingest key may be re-minted: a key a person revoked stays dead.
     * NULL for live keys and for keys revoked before the cause was recorded.
     */
    var revocationCause by ApiKeyTable.revocationCause

    /**
     * The CLI login key of the device session that minted this ingestion key.
     * Set by the CLI's personal mint, NULL for every other key. Revoking the
     * parent (logout, the devices tab, a re-login from the same device, session
     * expiry) revokes the children with it. Column ref only, no FK and no
     * index: lookups filter by the indexed `userId` first and match the parent
     * in memory over a person's few live ingest keys.
     */
    var parentApiKeyId by ApiKeyTable.parentApiKeyId

    /**
     * Non-null marks this key as an "ingestion key": a project-scoped,
     * ingest-only ApiKey the unified `langwatch <tool>` CLI (and template
     * installs) mint for OTLP Path B. The slug (claude_code / codex / gemini /
     * opencode / claude_cowork) is stamped as `langwatch.source` provenance on
     * every span the key authorizes. NULL for normal API / service keys.
     */
    var ingestSourceType by ApiKeyTable.ingestSourceType

    /**
     * Set only for ingestion keys derived from an admin IngestionTemplate
     * (e.g. claude_cowork). Drives receiver OTTL application + the
     * `langwatch.template.id` provenance stamp. NULL for unified-CLI tools
     * (which have no template) and for normal keys. Column ref only (no FK,
     * per the app-layer-integrity convention); resolved by the ingest-key service.
     */
    var ingestionTemplateId by ApiKeyTable.ingestionTemplateId

    /**
     * Human label of the CLI device session that minted this ingestion key
     * (e.g. "Rogerio's MacBook Pro"), snapshotted at mint time from the device
     * flow's `client_info`. Denormalized so the org-wide API-keys settings page
     * can show which device a key came from without resolving another user's
     * Redis-stored sessions. NULL for non-ingestion keys and for keys minted by
     * a CLI old enough to predate device metadata.
     */
    var createdByDeviceLabel by ApiKeyTable.createdByDeviceLabel

    /**
     * True for a key LangWatch minted for itself (a run, the gateway's trace
     * export, a Langy session). Hides it from listings, refuses customer edits
     * and revokes, and lets an ownerless run key act as the system. Never set
     * by a customer create; rows before the column stay false.
     */
    var isSystemManaged by ApiKeyTable.isSystemManaged

    /**
     * The browser session that minted a docs session key (ADR-185), so signing
     * out of that session revokes it. NULL for every other key. Unindexed:
     * sign-out filters by the indexed `userId` first.
     */
    var createdByBrowserSessionId by ApiKeyTable.createdByBrowserSessionId
    var createdAt by ApiKeyTable.createdAt
    var updatedAt by ApiKeyTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ApiKeyTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
