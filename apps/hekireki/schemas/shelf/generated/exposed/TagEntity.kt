package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class TagEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TagEntity>(TagTable)

    var name by TagTable.name
    var description by TagTable.description
    var color by TagTable.color
    var useFor by TagTable.useFor
    var assets by AssetEntity via AssetToTagTable
    var userId by TagTable.userId
    var user by UserEntity referencedOn TagTable.userId
    var organization by OrganizationEntity referencedOn TagTable.organizationId
    var organizationId by TagTable.organizationId
    var bookings by BookingEntity via BookingToTagTable
    var createdAt by TagTable.createdAt
    var updatedAt by TagTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(TagTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
