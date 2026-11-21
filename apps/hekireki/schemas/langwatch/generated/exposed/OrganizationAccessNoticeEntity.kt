package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * The mail log for seat requests and refused sign-ins: at most one row per organization,
 * kind, address and UTC day, each mailed at once.
 */
class OrganizationAccessNoticeEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OrganizationAccessNoticeEntity>(OrganizationAccessNoticeTable)

    var organizationId by OrganizationAccessNoticeTable.organizationId

    /** seat_request | idp_refused */
    var kind by OrganizationAccessNoticeTable.kind
    var email by OrganizationAccessNoticeTable.email
    var name by OrganizationAccessNoticeTable.name
    var idp by OrganizationAccessNoticeTable.idp

    /** UTC day the notice was recorded on, YYYY-MM-DD. */
    var dayKey by OrganizationAccessNoticeTable.dayKey

    /** Whether the seat request it announces was also refused by the identity provider. */
    var idpRefused by OrganizationAccessNoticeTable.idpRefused
    var createdAt by OrganizationAccessNoticeTable.createdAt

    /** Set once the notice was handed to the mailer; null if a send never started. */
    var mailedAt by OrganizationAccessNoticeTable.mailedAt
}
