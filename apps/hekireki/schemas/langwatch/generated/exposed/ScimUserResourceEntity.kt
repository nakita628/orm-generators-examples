package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import java.time.Instant
import java.time.temporal.ChronoUnit

/** Directory profile and lifecycle are tenant facts, independent of the shared account. */
class ScimUserResourceEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<ScimUserResourceEntity>(ScimUserResourceTable)

    val organizationId by ScimUserResourceTable.organizationId
    val userId by ScimUserResourceTable.userId
    var userName by ScimUserResourceTable.userName
    var name by ScimUserResourceTable.name
    var active by ScimUserResourceTable.active
    var deletedAt by ScimUserResourceTable.deletedAt
    var createdAt by ScimUserResourceTable.createdAt
    var updatedAt by ScimUserResourceTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ScimUserResourceTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
