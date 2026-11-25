package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/** The digest's one settings row. `armed` is the weekly send's master switch; it starts off. */
object DigestSettingsTable : IdTable<String>("\"DigestSettings\"") {
    override val id = text("id").default("singleton").entityId()
    val armed = bool("armed").default(false)
    val lastPlannedWeek = text("lastPlannedWeek").nullable()
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"DigestSettings_pkey\"")
}
