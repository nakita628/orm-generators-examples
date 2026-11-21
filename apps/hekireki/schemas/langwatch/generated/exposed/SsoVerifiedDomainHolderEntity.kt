package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

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
class SsoVerifiedDomainHolderEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<SsoVerifiedDomainHolderEntity>(SsoVerifiedDomainHolderTable)

    val domain by SsoVerifiedDomainHolderTable.domain
    val connectionId by SsoVerifiedDomainHolderTable.connectionId
    var organizationId by SsoVerifiedDomainHolderTable.organizationId
}
