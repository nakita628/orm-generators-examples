package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A person asked for a full seat the organization had no room for (modules/organization/specs/
 * seat-requests.feature). One row per organization and address; reopening reuses the row.
 */
class SeatRequestEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SeatRequestEntity>(SeatRequestTable)

    var organizationId by SeatRequestTable.organizationId

    /** Lowercased address of the person asking. */
    var email by SeatRequestTable.email

    /** Null when only the address is known (a sign-in the identity provider refused). */
    var userId by SeatRequestTable.userId
    var name by SeatRequestTable.name

    /** OPEN | APPROVED | DECLINED */
    var status by SeatRequestTable.status

    /** pending | lite | idp_refused: how the person was left when the request opened. */
    var source by SeatRequestTable.source1

    /** The identity provider also refused this person's sign-in. */
    var idpRefused by SeatRequestTable.idpRefused

    /** okta | entra | google | onelogin | jumpcloud | ping | auth0 | oidc | saml */
    var idpKind by SeatRequestTable.idpKind

    /** When sales were told about this request; reset when it reopens. */
    var salesNotifiedAt by SeatRequestTable.salesNotifiedAt
    var openedAt by SeatRequestTable.openedAt
    var resolvedAt by SeatRequestTable.resolvedAt
    var resolvedByUserId by SeatRequestTable.resolvedByUserId

    /**
     * Set when an approval claims the request, cleared once its seat is given; an approval
     * still pending long after is one whose seat never landed, and is reopened.
     */
    var seatPendingSince by SeatRequestTable.seatPendingSince
    var createdAt by SeatRequestTable.createdAt
    var updatedAt by SeatRequestTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SeatRequestTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
