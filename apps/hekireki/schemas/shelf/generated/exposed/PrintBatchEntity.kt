package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class PrintBatchEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, PrintBatchEntity>(PrintBatchTable)

    var name by PrintBatchTable.name
    var printed by PrintBatchTable.printed
    val qrCodes by QrEntity optionalReferrersOn QrTable.batchId
    var createdAt by PrintBatchTable.createdAt
    var updatedAt by PrintBatchTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(PrintBatchTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
