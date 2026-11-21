package formbricks.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * A creative's Impressions and Clicks on one UTC day on one device type (ENG-4000). Written
 * only by the counter flush job as `GREATEST(stored, counted)`, from Redis counters that hold the day's
 * running totals, so a repeated flush changes nothing. At most 3 rows per creative per day. Deleting
 * the creative deletes its counts.
 */
object CreativeDailyStatTable : CompositeIdTable("\"CreativeDailyStat\"") {
    val creativeId =
        reference(
            "creativeId",
            CreativeTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"CreativeDailyStat_creativeId_fkey\"",
        )

    /** UTC day. */
    val date = pgDate("date").entityId()
    val deviceType = pgEnum("deviceType", "\"CreativeDeviceType\"", CreativeDeviceType.entries, CreativeDeviceType::dbName).entityId()
    val impressions = integer("impressions").default(0)
    val clicks = integer("clicks").default(0)
    override val primaryKey = PrimaryKey(creativeId, date, deviceType, name = "\"CreativeDailyStat_pkey\"")

    init {
        addIdColumn(creativeId)
    }
}
