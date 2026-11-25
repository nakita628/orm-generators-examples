package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Each account nurturing erased, as a SHA-256 of its id and never the id: a gateway key fact
 * naming it, however late it lands, stores no creator or holder. Platform-wide, like the erasure.
 */
object NurturingErasedUserTable : IdTable<String>("\"NurturingErasedUser\"") {
    override val id = text("userIdHash").entityId()
    val erasedAt = pgTimestamp("erasedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"NurturingErasedUser_pkey\"")
}
