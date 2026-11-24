package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class BookingSettingsEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, BookingSettingsEntity>(BookingSettingsTable)

    var bufferStartTime by BookingSettingsTable.bufferStartTime
    var tagsRequired by BookingSettingsTable.tagsRequired
    var maxBookingLength by BookingSettingsTable.maxBookingLength
    var maxBookingLengthSkipClosedDays by BookingSettingsTable.maxBookingLengthSkipClosedDays
    var autoArchiveBookings by BookingSettingsTable.autoArchiveBookings
    var autoArchiveDays by BookingSettingsTable.autoArchiveDays
    var autoArchiveExpiredReservations by BookingSettingsTable.autoArchiveExpiredReservations
    var requireApprovalForBase by BookingSettingsTable.requireApprovalForBase
    var requireApprovalForSelfService by BookingSettingsTable.requireApprovalForSelfService
    var requireExplicitCheckinForAdmin by BookingSettingsTable.requireExplicitCheckinForAdmin
    var requireExplicitCheckinForSelfService by BookingSettingsTable.requireExplicitCheckinForSelfService
    var requireExplicitCheckoutForAdmin by BookingSettingsTable.requireExplicitCheckoutForAdmin
    var requireExplicitCheckoutForSelfService by BookingSettingsTable.requireExplicitCheckoutForSelfService
    var countKitsAsSingleUnit by BookingSettingsTable.countKitsAsSingleUnit
    var notifyBookingCreator by BookingSettingsTable.notifyBookingCreator
    var notifyAdminsOnNewBooking by BookingSettingsTable.notifyAdminsOnNewBooking
    var alwaysNotifyTeamMembers by TeamMemberEntity via BookingSettingsAlwaysNotifyTable
    var organization by OrganizationEntity referencedOn BookingSettingsTable.organizationId
    var organizationId by BookingSettingsTable.organizationId
    var createdAt by BookingSettingsTable.createdAt
    var updatedAt by BookingSettingsTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(BookingSettingsTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
