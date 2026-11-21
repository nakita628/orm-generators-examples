package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Defines org-scoped data store overrides (e.g. dedicated ClickHouse for HIPAA orgs).
 * Multiple organizations can share a single data store row via organizationIds.
 */
class OrganizationDataStoreEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OrganizationDataStoreEntity>(OrganizationDataStoreTable)

    /** Human-readable unique key (e.g. "hipaa-clickhouse-us-east") */
    var key by OrganizationDataStoreTable.key

    /** Organization IDs that use this data store */
    var organizationIds by OrganizationDataStoreTable.organizationIds
    var kind by OrganizationDataStoreTable.kind

    /** Versioned config JSON. Structure is discriminated by the top-level `version` field. */
    var config by OrganizationDataStoreTable.config
    var createdAt by OrganizationDataStoreTable.createdAt
    var updatedAt by OrganizationDataStoreTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(OrganizationDataStoreTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
