package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Per-(organization, tool) override for which `langwatch <tool>` paths the
 * CLI may use: `allowVk` gates the gateway / virtual-key path (Path A) and
 * `allowOtelDirect` gates the direct-OTLP ingestion path (Path B). A missing
 * row resolves to the hardcoded defaults (both true for claude/codex/gemini/
 * opencode; cursor allows the gateway path only), so the policy is purely
 * additive: disabling a path is an explicit admin choice. The CLI caches the
 * resolved map at login and gates path selection on it.
 * See specs/ai-gateway/governance/cli-tool-mode-policy.feature.
 */
object PlatformToolPolicyTable : IdTable<String>("\"PlatformToolPolicy\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")
    val toolSlug = text("toolSlug")
    val allowVk = bool("allowVk").default(true)
    val allowOtelDirect = bool("allowOtelDirect").default(true)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"PlatformToolPolicy_pkey\"")

    init {
        index("\"PlatformToolPolicy_organizationId_idx\"", false, organizationId)
        uniqueIndex("\"PlatformToolPolicy_organizationId_toolSlug_key\"", organizationId, toolSlug)
    }
}
