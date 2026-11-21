package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class PlatformToolPolicyEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, PlatformToolPolicyEntity>(PlatformToolPolicyTable)

    var organizationId by PlatformToolPolicyTable.organizationId
    var toolSlug by PlatformToolPolicyTable.toolSlug
    var allowVk by PlatformToolPolicyTable.allowVk
    var allowOtelDirect by PlatformToolPolicyTable.allowOtelDirect
    var createdAt by PlatformToolPolicyTable.createdAt
    var updatedAt by PlatformToolPolicyTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(PlatformToolPolicyTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
