package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * A connection currently entitled to use an organization-owned verified
 * domain. Usually there is one holder. During an Auth0 migration there are
 * exactly two: the grandfathered predecessor and its direct replacement.
 *
 * The projection repository enforces that relationship before inserting a
 * second holder. The flattened migration also installs partial unique live
 * legacy/direct slots on `SsoConnection`, closing concurrent third-writer
 * races that Prisma cannot describe declaratively.
 */
object SsoVerifiedDomainHolderTable : CompositeIdTable("\"SsoVerifiedDomainHolder\"") {
    val domain = text("domain").entityId()
    val connectionId = text("connectionId").entityId()
    val organizationId = text("organizationId")
    override val primaryKey = PrimaryKey(domain, connectionId, name = "\"SsoVerifiedDomainHolder_pkey\"")

    init {
        index("\"SsoVerifiedDomainHolder_connectionId_idx\"", false, connectionId)
        index("\"SsoVerifiedDomainHolder_organizationId_idx\"", false, organizationId)
    }
}
