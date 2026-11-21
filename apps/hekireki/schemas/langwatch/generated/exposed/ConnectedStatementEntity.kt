package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * The monthly usage statement one connected customer was sent (ADR-139,
 * section 7). One row per account per month keeps the statement to one send,
 * whatever runs the tick.
 */
class ConnectedStatementEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ConnectedStatementEntity>(ConnectedStatementTable)

    /** The `ConnectedBillingAccount` row. A plain column ref. */
    var accountId by ConnectedStatementTable.accountId

    /** First day of the month the statement covers, in UTC. */
    var month by ConnectedStatementTable.month
    var sentAt by ConnectedStatementTable.sentAt
}
