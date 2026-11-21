package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AnnouncementEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AnnouncementEntity>(AnnouncementTable)

    var name by AnnouncementTable.name
    var content by AnnouncementTable.content
    var link by AnnouncementTable.link
    var linkText by AnnouncementTable.linkText
    var published by AnnouncementTable.published
    var createdAt by AnnouncementTable.createdAt
    var updatedAt by AnnouncementTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AnnouncementTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
