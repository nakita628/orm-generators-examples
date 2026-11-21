package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class CustomerQueryEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, CustomerQueryEntity>(CustomerQueryTable)

    /** The TSQL query text that was executed */
    var query by CustomerQueryTable.query

    /** The scope of the query (determines which tenant IDs were used) */
    var scope by CustomerQueryTable.scope

    /** Query execution statistics from ClickHouse */
    var stats by CustomerQueryTable.stats

    /** AI-generated title summarizing the query */
    var title by CustomerQueryTable.title

    /** Where the query originated from */
    var source by CustomerQueryTable.source1
    var organization by OrganizationEntity referencedOn CustomerQueryTable.organizationId
    var organizationId by CustomerQueryTable.organizationId
    var project by ProjectEntity optionalReferencedOn CustomerQueryTable.projectId
    var projectId by CustomerQueryTable.projectId
    var environment by RuntimeEnvironmentEntity optionalReferencedOn CustomerQueryTable.environmentId
    var environmentId by CustomerQueryTable.environmentId

    /** Optional user who executed the query (null for API calls) */
    var user by UserEntity optionalReferencedOn CustomerQueryTable.userId
    var userId by CustomerQueryTable.userId

    /**
     * Time filter settings used when the query was executed
     * Period like "7d", "24h", etc.
     */
    var filterPeriod by CustomerQueryTable.filterPeriod

    /** Custom start date (ISO 8601) */
    var filterFrom by CustomerQueryTable.filterFrom

    /** Custom end date (ISO 8601) */
    var filterTo by CustomerQueryTable.filterTo
    var createdAt by CustomerQueryTable.createdAt
}
