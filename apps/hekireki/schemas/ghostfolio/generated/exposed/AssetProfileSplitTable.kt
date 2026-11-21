package ghostfolio.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit
import java.util.UUID

/**
 * A stock split of an asset profile. The numerator is the number of shares
 * held after the split, the denominator the number of shares held before, for
 * example 4 and 1 for a 4:1 split or 1 and 10 for a 1:10 reverse split.
 *
 * The ratio is stored as two integers instead of the resulting factor, so that
 * it stays exact, for example 1/3 for a 1:3 reverse split.
 *
 * Activities before the date are adjusted by the ratio: the quantity is
 * multiplied by it and the unit price divided by it, so that the total value
 * of an activity stays the same.
 *
 * Market data is already split-adjusted by the data providers and must not be
 * adjusted again.
 */
object AssetProfileSplitTable : IdTable<String>("\"AssetProfileSplit\"") {
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val date = pgTimestamp("date", 3)
    val denominator = integer("denominator")
    override val id = text("id").clientDefault { UUID.randomUUID().toString() }.entityId()
    val numerator = integer("numerator")
    val symbolProfileId =
        reference(
            "symbolProfileId",
            SymbolProfileTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetProfileSplit_symbolProfileId_fkey\"",
        )
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AssetProfileSplit_pkey\"")

    init {
        index("\"AssetProfileSplit_date_idx\"", false, date)
        uniqueIndex("\"AssetProfileSplit_symbolProfileId_date_key\"", symbolProfileId, date)
    }
}
