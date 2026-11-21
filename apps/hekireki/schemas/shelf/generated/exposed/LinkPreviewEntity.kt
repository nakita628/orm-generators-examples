package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * What Shelf knows about one web link used in a workspace's custom fields,
 * keyed by its `LinkKey` from @shelf/links. Two writers on disjoint columns:
 * the link preview job (fetch columns) and people (customTitle).
 */
class LinkPreviewEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, LinkPreviewEntity>(LinkPreviewTable)

    /**
     * `LinkKey` from @shelf/links, never a raw field value. At most 2048
     * characters (CHECK constraint in the migration).
     */
    var url by LinkPreviewTable.url
    var pageTitle by LinkPreviewTable.pageTitle

    /** Null until the first fetch finishes. */
    var outcome by LinkPreviewTable.outcome
    var attempts by LinkPreviewTable.attempts
    var fetchedAt by LinkPreviewTable.fetchedAt

    /**
     * When the link may be fetched next; null means never again. A claim moves
     * it a few minutes ahead, which is also the lease on a fetch in flight.
     */
    var nextFetchAt by LinkPreviewTable.nextFetchAt

    /** The name a person gave the link in this workspace; wins over pageTitle. */
    var customTitle by LinkPreviewTable.customTitle
    var createdAt by LinkPreviewTable.createdAt
    var updatedAt by LinkPreviewTable.updatedAt
    var organization by OrganizationEntity referencedOn LinkPreviewTable.organizationId
    var organizationId by LinkPreviewTable.organizationId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(LinkPreviewTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
