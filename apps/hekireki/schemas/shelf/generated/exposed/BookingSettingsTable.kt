package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object BookingSettingsTable : IdTable<String>("\"BookingSettings\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val bufferStartTime = integer("bufferStartTime").default(0)
    val tagsRequired = bool("tagsRequired").default(false)
    val maxBookingLength = integer("maxBookingLength").nullable()
    val maxBookingLengthSkipClosedDays = bool("maxBookingLengthSkipClosedDays").default(false)
    val autoArchiveBookings = bool("autoArchiveBookings").default(false)
    val autoArchiveDays = integer("autoArchiveDays").default(2)
    val autoArchiveExpiredReservations = bool("autoArchiveExpiredReservations").default(false)
    val requireApprovalForBase = bool("requireApprovalForBase").default(false)
    val requireApprovalForSelfService = bool("requireApprovalForSelfService").default(false)
    val requireExplicitCheckinForAdmin = bool("requireExplicitCheckinForAdmin").default(false)
    val requireExplicitCheckinForSelfService = bool("requireExplicitCheckinForSelfService").default(false)
    val requireExplicitCheckoutForAdmin = bool("requireExplicitCheckoutForAdmin").default(false)
    val requireExplicitCheckoutForSelfService = bool("requireExplicitCheckoutForSelfService").default(false)
    val countKitsAsSingleUnit = bool("countKitsAsSingleUnit").default(false)
    val notifyBookingCreator = bool("notifyBookingCreator").default(true)
    val notifyAdminsOnNewBooking = bool("notifyAdminsOnNewBooking").default(true)
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"BookingSettings_organizationId_fkey\"",
        ).uniqueIndex("\"BookingSettings_organizationId_key\"")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"BookingSettings_pkey\"")

    init {
        index("\"BookingSettings_organizationId_idx\"", false, organizationId)
    }
}
