package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class ReportFoundEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ReportFoundEntity>(ReportFoundTable)

    var email by ReportFoundTable.email
    var content by ReportFoundTable.content
    var createdAt by ReportFoundTable.createdAt
    var updatedAt by ReportFoundTable.updatedAt
    var asset by AssetEntity optionalReferencedOn ReportFoundTable.assetId
    var assetId by ReportFoundTable.assetId
    var kit by KitEntity optionalReferencedOn ReportFoundTable.kitId
    var kitId by ReportFoundTable.kitId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ReportFoundTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
