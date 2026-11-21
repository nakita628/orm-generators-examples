package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class SavedViewEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SavedViewEntity>(SavedViewTable)

    var projectId by SavedViewTable.projectId
    var userId by SavedViewTable.userId
    var name by SavedViewTable.name
    var filters by SavedViewTable.filters
    var query by SavedViewTable.query
    var period by SavedViewTable.period
    var order by SavedViewTable.order

    /**
     * Discriminator for the two storage shapes that share this table:
     * - "v1-traces-filter" (default) - legacy /messages page filter map saved views
     * - "v2-traces-lens"             - new traces v2 lens configs (columns, addons, grouping, sort, filterText) encoded inside `filters` JSON
     * Each client filters by its own kind on read so neither sees the other's rows.
     */
    var kind by SavedViewTable.kind
    var createdAt by SavedViewTable.createdAt
    var updatedAt by SavedViewTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SavedViewTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
