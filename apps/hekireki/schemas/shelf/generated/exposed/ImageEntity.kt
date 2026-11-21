package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class ImageEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ImageEntity>(ImageTable)

    var contentType by ImageTable.contentType
    var altText by ImageTable.altText
    var blob by ImageTable.blob
    var createdAt by ImageTable.createdAt
    var updatedAt by ImageTable.updatedAt
    val location by LocationEntity optionalBackReferencedOn LocationTable.imageId
    val organization by OrganizationEntity optionalBackReferencedOn OrganizationTable.imageId
    var ownerOrg by OrganizationEntity referencedOn ImageTable.ownerOrgId
    var ownerOrgId by ImageTable.ownerOrgId
    var user by UserEntity referencedOn ImageTable.userId
    var userId by ImageTable.userId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ImageTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
