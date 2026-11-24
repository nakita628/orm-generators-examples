package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object BookingTable : IdTable<String>("\"Booking\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name")
    val status = pgEnum("status", "\"BookingStatus\"", BookingStatus.entries, BookingStatus::dbName).default(BookingStatus.DRAFT)
    val description = text("description").nullable().default("")
    val activeSchedulerReference = text("activeSchedulerReference").nullable()
    val creatorId =
        reference(
            "creatorId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Booking_creatorId_fkey\"",
        )
    val custodianUserId =
        optReference(
            "custodianUserId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Booking_custodianUserId_fkey\"",
        )
    val custodianTeamMemberId =
        optReference(
            "custodianTeamMemberId",
            TeamMemberTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Booking_custodianTeamMemberId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Booking_organizationId_fkey\"",
        )
    val createdAt = pgTimestamptz("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamptz("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val from = pgTimestamptz("from", 3)
    val to = pgTimestamptz("to", 3)
    val originalFrom = pgTimestamptz("originalFrom", 3).nullable()
    val originalTo = pgTimestamptz("originalTo", 3).nullable()
    val autoArchivedAt = pgTimestamptz("autoArchivedAt", 3).nullable()
    val archivedWithoutCheckin = bool("archivedWithoutCheckin").default(false)
    val cancellationReason = text("cancellationReason").nullable()
    val approvedAt = pgTimestamptz("approvedAt", 3).nullable()
    val approvedById =
        optReference(
            "approvedById",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Booking_approvedById_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"Booking_pkey\"")

    init {
        index("\"Booking_creatorId_idx\"", false, creatorId)
        index("\"Booking_custodianTeamMemberId_idx\"", false, custodianTeamMemberId)
        index("\"Booking_custodianUserId_idx\"", false, custodianUserId)
        index("\"Booking_approvedById_idx\"", false, approvedById)
        index("\"Booking_organizationId_idx\"", false, organizationId)
        index("\"Booking_status_idx\"", false, status)
        index("\"Booking_organizationId_from_id_idx\"", false, organizationId, from, id)
    }
}
