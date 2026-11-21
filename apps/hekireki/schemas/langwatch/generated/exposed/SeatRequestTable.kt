package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A person asked for a full seat the organization had no room for (modules/organization/specs/
 * seat-requests.feature). One row per organization and address; reopening reuses the row.
 */
object SeatRequestTable : IdTable<String>("\"SeatRequest\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")

    /** Lowercased address of the person asking. */
    val email = text("email")

    /** Null when only the address is known (a sign-in the identity provider refused). */
    val userId = text("userId").nullable()
    val name = text("name").nullable()

    /** OPEN | APPROVED | DECLINED */
    val status = text("status")

    /** pending | lite | idp_refused: how the person was left when the request opened. */
    val source1 = text("source")

    /** The identity provider also refused this person's sign-in. */
    val idpRefused = bool("idpRefused").default(false)

    /** okta | entra | google | onelogin | jumpcloud | ping | auth0 | oidc | saml */
    val idpKind = text("idpKind").nullable()

    /** When sales were told about this request; reset when it reopens. */
    val salesNotifiedAt = pgTimestamp("salesNotifiedAt", 3).nullable()
    val openedAt = pgTimestamp("openedAt", 3)
    val resolvedAt = pgTimestamp("resolvedAt", 3).nullable()
    val resolvedByUserId = text("resolvedByUserId").nullable()

    /**
     * Set when an approval claims the request, cleared once its seat is given; an approval
     * still pending long after is one whose seat never landed, and is reopened.
     */
    val seatPendingSince = pgTimestamp("seatPendingSince", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"SeatRequest_pkey\"")

    init {
        index("\"SeatRequest_organizationId_status_idx\"", false, organizationId, status)
        index("\"SeatRequest_userId_idx\"", false, userId)
        uniqueIndex("\"SeatRequest_organizationId_email_key\"", organizationId, email)
    }
}
