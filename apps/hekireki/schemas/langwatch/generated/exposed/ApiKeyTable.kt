package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object ApiKeyTable : IdTable<String>("\"ApiKey\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val name = text("name")
    val description = text("description").nullable()
    val lookupId = text("lookupId").uniqueIndex("\"ApiKey_lookupId_key\"")
    val hashedSecret = text("hashedSecret")
    val permissionMode = text("permissionMode").default("all")
    val userId = text("userId").nullable()
    val createdByUserId = text("createdByUserId").nullable()
    val organizationId = text("organizationId")
    val expiresAt = pgTimestamp("expiresAt", 3).nullable()
    val lastUsedAt = pgTimestamp("lastUsedAt", 3).nullable()
    val revokedAt = pgTimestamp("revokedAt", 3).nullable()

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
    val revocationCause = text("revocationCause").nullable()

    /**
     * The CLI login key of the device session that minted this ingestion key.
     * Set by the CLI's personal mint, NULL for every other key. Revoking the
     * parent (logout, the devices tab, a re-login from the same device, session
     * expiry) revokes the children with it. Column ref only, no FK and no
     * index: lookups filter by the indexed `userId` first and match the parent
     * in memory over a person's few live ingest keys.
     */
    val parentApiKeyId = text("parentApiKeyId").nullable()

    /**
     * Non-null marks this key as an "ingestion key": a project-scoped,
     * ingest-only ApiKey the unified `langwatch <tool>` CLI (and template
     * installs) mint for OTLP Path B. The slug (claude_code / codex / gemini /
     * opencode / claude_cowork) is stamped as `langwatch.source` provenance on
     * every span the key authorizes. NULL for normal API / service keys.
     */
    val ingestSourceType = text("ingestSourceType").nullable()

    /**
     * Set only for ingestion keys derived from an admin IngestionTemplate
     * (e.g. claude_cowork). Drives receiver OTTL application + the
     * `langwatch.template.id` provenance stamp. NULL for unified-CLI tools
     * (which have no template) and for normal keys. Column ref only (no FK,
     * per the app-layer-integrity convention); resolved by the ingest-key service.
     */
    val ingestionTemplateId = text("ingestionTemplateId").nullable()

    /**
     * Human label of the CLI device session that minted this ingestion key
     * (e.g. "Rogerio's MacBook Pro"), snapshotted at mint time from the device
     * flow's `client_info`. Denormalized so the org-wide API-keys settings page
     * can show which device a key came from without resolving another user's
     * Redis-stored sessions. NULL for non-ingestion keys and for keys minted by
     * a CLI old enough to predate device metadata.
     */
    val createdByDeviceLabel = text("createdByDeviceLabel").nullable()

    /**
     * True for a key LangWatch minted for itself (a run, the gateway's trace
     * export, a Langy session). Hides it from listings, refuses customer edits
     * and revokes, and lets an ownerless run key act as the system. Never set
     * by a customer create; rows before the column stay false.
     */
    val isSystemManaged = bool("isSystemManaged").default(false)

    /**
     * The browser session that minted a docs session key (ADR-185), so signing
     * out of that session revokes it. NULL for every other key. Unindexed:
     * sign-out filters by the indexed `userId` first.
     */
    val createdByBrowserSessionId = text("createdByBrowserSessionId").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"ApiKey_pkey\"")

    init {
        index("\"ApiKey_userId_idx\"", false, userId)
        index("\"ApiKey_organizationId_idx\"", false, organizationId)
        index("\"ApiKey_lookupId_idx\"", false, lookupId)
        index("\"ApiKey_ingestSourceType_idx\"", false, ingestSourceType)
        index("\"ApiKey_ingestionTemplateId_idx\"", false, ingestionTemplateId)
    }
}
