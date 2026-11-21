package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class UpdateEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, UpdateEntity>(UpdateTable)

    var title by UpdateTable.title
    var content by UpdateTable.content
    var url by UpdateTable.url
    var imageUrl by UpdateTable.imageUrl
    var publishDate by UpdateTable.publishDate
    var status by UpdateTable.status
    var targetRoles by UpdateTable.targetRoles
    var clickCount by UpdateTable.clickCount
    var viewCount by UpdateTable.viewCount
    val userReads by UserUpdateReadEntity referrersOn UserUpdateReadTable.updateId
    var createdBy by UserEntity referencedOn UpdateTable.createdById
    var createdById by UpdateTable.createdById
    var createdAt by UpdateTable.createdAt
    var updatedAt by UpdateTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(UpdateTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
