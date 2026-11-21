package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The mail log for seat requests and refused sign-ins: at most one row per organization,
 * kind, address and UTC day, each mailed at once.
 */
object OrganizationAccessNoticeTable : IdTable<String>("\"OrganizationAccessNotice\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")

    /** seat_request | idp_refused */
    val kind = text("kind")
    val email = text("email")
    val name = text("name").nullable()
    val idp = text("idp").nullable()

    /** UTC day the notice was recorded on, YYYY-MM-DD. */
    val dayKey = text("dayKey")

    /** Whether the seat request it announces was also refused by the identity provider. */
    val idpRefused = bool("idpRefused").default(false)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /** Set once the notice was handed to the mailer; null if a send never started. */
    val mailedAt = pgTimestamp("mailedAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"OrganizationAccessNotice_pkey\"")

    init {
        uniqueIndex("\"OrganizationAccessNotice_organizationId_kind_email_dayKey_key\"", organizationId, kind, email, dayKey)
    }
}
