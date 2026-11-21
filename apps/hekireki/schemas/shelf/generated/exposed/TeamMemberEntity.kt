package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class TeamMemberEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TeamMemberEntity>(TeamMemberTable)

    var name by TeamMemberTable.name
    var organization by OrganizationEntity referencedOn TeamMemberTable.organizationId
    var organizationId by TeamMemberTable.organizationId
    val custodies by CustodyEntity referrersOn CustodyTable.teamMemberId
    val receivedInvites by InviteEntity referrersOn InviteTable.teamMemberId
    var user by UserEntity optionalReferencedOn TeamMemberTable.userId
    var userId by TeamMemberTable.userId
    val kitCustodies by KitCustodyEntity referrersOn KitCustodyTable.custodianId
    var createdAt by TeamMemberTable.createdAt
    var updatedAt by TeamMemberTable.updatedAt
    var deletedAt by TeamMemberTable.deletedAt
    val bookings by BookingEntity optionalReferrersOn BookingTable.custodianTeamMemberId
    var assetReminders by AssetReminderEntity via AssetReminderToTeamMemberTable
    val consumptionLogs by ConsumptionLogEntity optionalReferrersOn ConsumptionLogTable.custodianId
    val teamMemberNotes by TeamMemberNoteEntity referrersOn TeamMemberNoteTable.teamMemberId
    var alwaysNotifyBookingSettings by BookingSettingsEntity via BookingSettingsAlwaysNotifyTable
    var bookingNotifications by BookingEntity via BookingNotificationRecipientsTable

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(TeamMemberTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
