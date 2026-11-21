package formbricks.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * A creative's Impressions and Clicks on one UTC day on one page (host + path), written like
 * `CreativeDailyStat`. At most 1000 pages per creative per day; the rest are counted under `(other)`.
 */
object CreativeDailyPageStatTable : CompositeIdTable("\"CreativeDailyPageStat\"") {
    val creativeId =
        reference(
            "creativeId",
            CreativeTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"CreativeDailyPageStat_creativeId_fkey\"",
        )

    /** UTC day. */
    val date = pgDate("date").entityId()

    /** Host + path, without query string or hash, or `(other)`. */
    val page = text("page").entityId()
    val impressions = integer("impressions").default(0)
    val clicks = integer("clicks").default(0)
    override val primaryKey = PrimaryKey(creativeId, date, page, name = "\"CreativeDailyPageStat_pkey\"")

    init {
        addIdColumn(creativeId)
    }
}
