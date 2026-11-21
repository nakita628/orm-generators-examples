package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object UserTable : IdTable<String>("\"User\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val email = text("email").uniqueIndex("\"User_email_key\"")
    val username = text("username").clientDefault { CUID.randomCUID1().toString() }.uniqueIndex("\"User_username_key\"")
    val firstName = text("firstName").nullable()
    val lastName = text("lastName").nullable()
    val displayName = text("displayName").nullable()
    val profilePicture = text("profilePicture").nullable()
    val usedFreeTrial = bool("usedFreeTrial").default(false)
    val onboarded = bool("onboarded").default(false)
    val customerId = text("customerId").nullable().uniqueIndex("\"User_customerId_key\"")
    val sso = bool("sso").default(false)
    val createdWithInvite = bool("createdWithInvite").default(false)
    val skipSubscriptionCheck = bool("skipSubscriptionCheck").default(false)
    val hasUnpaidInvoice = bool("hasUnpaidInvoice").default(false)
    val warnForNoPaymentMethod = bool("warnForNoPaymentMethod").default(false)
    val lastSelectedOrganizationId =
        optReference(
            "lastSelectedOrganizationId",
            OrganizationTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"User_lastSelectedOrganizationId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val deletedAt = pgTimestamp("deletedAt", 3).nullable()
    val lastMobileActiveAt = pgTimestamp("lastMobileActiveAt", 3).nullable()
    val dateFormat = pgEnum("dateFormat", "\"DateFormatPreference\"", DateFormatPreference.entries, DateFormatPreference::dbName).nullable()
    val timeFormat = pgEnum("timeFormat", "\"TimeFormatPreference\"", TimeFormatPreference.entries, TimeFormatPreference::dbName).nullable()
    val weekStart = pgEnum("weekStart", "\"WeekStartPreference\"", WeekStartPreference.entries, WeekStartPreference::dbName).nullable()
    val timeZone = text("timeZone").nullable()
    val tierId =
        reference(
            "tierId",
            TierTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"User_tierId_fkey\"",
        ).default(EntityID(TierId.FREE, TierTable))
    val referralSource = text("referralSource").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"User_pkey\"")

    init {
        index("\"User_firstName_lastName_idx\"", false, firstName, lastName)
        index("\"User_tierId_idx\"", false, tierId)
        index("\"User_lastSelectedOrganizationId_idx\"", false, lastSelectedOrganizationId)
        index("\"User_lastMobileActiveAt_idx\"", false, lastMobileActiveAt)
        uniqueIndex("\"User_email_username_key\"", email, username)
    }
}
