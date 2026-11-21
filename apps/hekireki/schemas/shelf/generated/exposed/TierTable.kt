package shelf.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object TierTable : IdTable<TierId>("\"Tier\"") {
    override val id = pgEnum("id", "\"TierId\"", TierId.entries, TierId::dbName).entityId()
    val name = text("name")
    val tierLimitId =
        optReference(
            "tierLimitId",
            TierLimitTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Tier_tierLimitId_fkey\"",
        ).uniqueIndex("\"Tier_tierLimitId_key\"")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Tier_pkey\"")
}
