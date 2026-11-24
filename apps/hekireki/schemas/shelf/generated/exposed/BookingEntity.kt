package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class BookingEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, BookingEntity>(BookingTable)

    var name by BookingTable.name
    var status by BookingTable.status
    var description by BookingTable.description
    var activeSchedulerReference by BookingTable.activeSchedulerReference
    var creator by UserEntity referencedOn BookingTable.creatorId
    var creatorId by BookingTable.creatorId
    var custodianUser by UserEntity optionalReferencedOn BookingTable.custodianUserId
    var custodianUserId by BookingTable.custodianUserId
    var custodianTeamMember by TeamMemberEntity optionalReferencedOn BookingTable.custodianTeamMemberId
    var custodianTeamMemberId by BookingTable.custodianTeamMemberId
    var organization by OrganizationEntity referencedOn BookingTable.organizationId
    var organizationId by BookingTable.organizationId
    val bookingAssets by BookingAssetEntity referrersOn BookingAssetTable.bookingId
    val consumptionLogs by ConsumptionLogEntity optionalReferrersOn ConsumptionLogTable.bookingId
    val notes by BookingNoteEntity referrersOn BookingNoteTable.bookingId
    val modelRequests by BookingModelRequestEntity referrersOn BookingModelRequestTable.bookingId
    var tags by TagEntity via BookingToTagTable
    val partialCheckins by PartialBookingCheckinEntity referrersOn PartialBookingCheckinTable.bookingId
    val partialCheckouts by PartialBookingCheckoutEntity referrersOn PartialBookingCheckoutTable.bookingId
    var createdAt by BookingTable.createdAt
    var updatedAt by BookingTable.updatedAt
    var from by BookingTable.from
    var to by BookingTable.to
    var originalFrom by BookingTable.originalFrom
    var originalTo by BookingTable.originalTo
    var autoArchivedAt by BookingTable.autoArchivedAt
    var archivedWithoutCheckin by BookingTable.archivedWithoutCheckin
    var cancellationReason by BookingTable.cancellationReason
    var approvedAt by BookingTable.approvedAt
    var approvedBy by UserEntity optionalReferencedOn BookingTable.approvedById
    var approvedById by BookingTable.approvedById
    var notificationRecipients by TeamMemberEntity via BookingNotificationRecipientsTable

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(BookingTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
