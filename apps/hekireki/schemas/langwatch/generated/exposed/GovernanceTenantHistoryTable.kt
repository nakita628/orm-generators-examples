package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Every ClickHouse `TenantId` this organization has ever written governance
 * rows under, appended on first use and never pruned.
 *
 * The live `resolveGovProjectId` lookup cannot be this translation: it filters
 * `archivedAt: null` and so returns null forever once someone archives the
 * governance project, while the write path re-reads the same project by slug
 * with no such filter and keeps landing rows under the old tenant. The
 * organization would map to zero tenants on read and one on write - a
 * permanent split-brain in which the cost screen reports "no governance data"
 * and an erasure job erases nothing and reports success.
 *
 * Reads resolve against the whole history; erasure walks all of it, because
 * personal data does not stop existing in a tenant that stopped being current.
 */
object GovernanceTenantHistoryTable : IdTable<String>("\"GovernanceTenantHistory\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")

    /** A hidden governance project id this organization has written rows under. */
    val tenantId = text("tenantId")
    val firstUsedAt = pgTimestamp("firstUsedAt", 3)
    val lastUsedAt = pgTimestamp("lastUsedAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"GovernanceTenantHistory_pkey\"")

    init {
        index("\"GovernanceTenantHistory_organizationId_idx\"", false, organizationId)
        uniqueIndex("\"GovernanceTenantHistory_organizationId_tenantId_key\"", organizationId, tenantId)
    }
}
