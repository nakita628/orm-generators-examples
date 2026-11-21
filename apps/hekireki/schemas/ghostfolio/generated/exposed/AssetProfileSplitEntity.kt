package ghostfolio.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

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
class AssetProfileSplitEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AssetProfileSplitEntity>(AssetProfileSplitTable)

    var createdAt by AssetProfileSplitTable.createdAt
    var date by AssetProfileSplitTable.date
    var denominator by AssetProfileSplitTable.denominator
    var numerator by AssetProfileSplitTable.numerator
    var symbolProfile by SymbolProfileEntity referencedOn AssetProfileSplitTable.symbolProfileId
    var symbolProfileId by AssetProfileSplitTable.symbolProfileId
    var updatedAt by AssetProfileSplitTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AssetProfileSplitTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
