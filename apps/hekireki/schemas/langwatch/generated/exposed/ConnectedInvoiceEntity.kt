package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * An invoice LangWatch raised for a connected customer, as the backoffice
 * reads it.
 */
class ConnectedInvoiceEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ConnectedInvoiceEntity>(ConnectedInvoiceTable)

    /** The `ConnectedBillingAccount` row, as a plain column ref. */
    var accountId by ConnectedInvoiceTable.accountId
    var stripeInvoiceId by ConnectedInvoiceTable.stripeInvoiceId
    var kind by ConnectedInvoiceTable.kind
    var currency by ConnectedInvoiceTable.currency
    var amountCents by ConnectedInvoiceTable.amountCents

    /** The provider's own status: `draft`, `open`, `paid`, `void` or `uncollectible`. */
    var status by ConnectedInvoiceTable.status
    var paidOutOfBandAt by ConnectedInvoiceTable.paidOutOfBandAt
    var termStartsAt by ConnectedInvoiceTable.termStartsAt
    var createdAt by ConnectedInvoiceTable.createdAt
    var updatedAt by ConnectedInvoiceTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ConnectedInvoiceTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
