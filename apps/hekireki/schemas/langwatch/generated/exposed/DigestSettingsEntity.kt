package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/** The digest's one settings row. `armed` is the weekly send's master switch; it starts off. */
class DigestSettingsEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DigestSettingsEntity>(DigestSettingsTable)

    var armed by DigestSettingsTable.armed
    var lastPlannedWeek by DigestSettingsTable.lastPlannedWeek
    var updatedAt by DigestSettingsTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(DigestSettingsTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
