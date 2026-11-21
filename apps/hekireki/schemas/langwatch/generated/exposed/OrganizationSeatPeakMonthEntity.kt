package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/** One calendar month's (UTC) seat peaks for an organization (ADR-002). */
class OrganizationSeatPeakMonthEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<OrganizationSeatPeakMonthEntity>(OrganizationSeatPeakMonthTable)

    val organizationId by OrganizationSeatPeakMonthTable.organizationId

    /** `YYYY-MM`. */
    val month by OrganizationSeatPeakMonthTable.month
    var fullPeak by OrganizationSeatPeakMonthTable.fullPeak
    var fullPeakAt by OrganizationSeatPeakMonthTable.fullPeakAt
    var litePeak by OrganizationSeatPeakMonthTable.litePeak
    var litePeakAt by OrganizationSeatPeakMonthTable.litePeakAt
    var fullAtEnd by OrganizationSeatPeakMonthTable.fullAtEnd
    var liteAtEnd by OrganizationSeatPeakMonthTable.liteAtEnd
}
