package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * A creative's Impressions and Clicks on one UTC day on one device type (ENG-4000). Written
 * only by the counter flush job as `GREATEST(stored, counted)`, from Redis counters that hold the day's
 * running totals, so a repeated flush changes nothing. At most 3 rows per creative per day. Deleting
 * the creative deletes its counts.
 */
class CreativeDailyStatEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<CreativeDailyStatEntity>(CreativeDailyStatTable)

    val creative by CreativeEntity referencedOn CreativeDailyStatTable.creativeId
    val creativeId by CreativeDailyStatTable.creativeId

    /** UTC day. */
    val date by CreativeDailyStatTable.date
    val deviceType by CreativeDailyStatTable.deviceType
    var impressions by CreativeDailyStatTable.impressions
    var clicks by CreativeDailyStatTable.clicks
}
