package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/** One calendar month's (UTC) seat peaks for an organization (ADR-002). */
object OrganizationSeatPeakMonthTable : CompositeIdTable("\"OrganizationSeatPeakMonth\"") {
    val organizationId = text("organizationId").entityId()

    /** `YYYY-MM`. */
    val month = text("month").entityId()
    val fullPeak = integer("fullPeak")
    val fullPeakAt = double("fullPeakAt")
    val litePeak = integer("litePeak")
    val litePeakAt = double("litePeakAt")
    val fullAtEnd = integer("fullAtEnd")
    val liteAtEnd = integer("liteAtEnd")
    override val primaryKey = PrimaryKey(organizationId, month, name = "\"OrganizationSeatPeakMonth_pkey\"")
}
