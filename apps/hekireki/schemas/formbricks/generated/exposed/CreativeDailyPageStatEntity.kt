package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * A creative's Impressions and Clicks on one UTC day on one page (host + path), written like
 * `CreativeDailyStat`. At most 1000 pages per creative per day; the rest are counted under `(other)`.
 */
class CreativeDailyPageStatEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<CreativeDailyPageStatEntity>(CreativeDailyPageStatTable)

    val creative by CreativeEntity referencedOn CreativeDailyPageStatTable.creativeId
    val creativeId by CreativeDailyPageStatTable.creativeId

    /** UTC day. */
    val date by CreativeDailyPageStatTable.date

    /** Host + path, without query string or hash, or `(other)`. */
    val page by CreativeDailyPageStatTable.page
    var impressions by CreativeDailyPageStatTable.impressions
    var clicks by CreativeDailyPageStatTable.clicks
}
