package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * A paid credit at the payment provider, which only metered usage draws down.
 * Kept here so a grant is created once per reason and the backoffice can show
 * what the customer holds without asking the provider.
 */
class ConnectedCreditGrantEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ConnectedCreditGrantEntity>(ConnectedCreditGrantTable)

    /**
     * The `ConnectedBillingAccount` row. A plain column ref, like
     * `IssuedLicense.replacesId`: relationMode="prisma" carries no foreign keys.
     */
    var accountId by ConnectedCreditGrantTable.accountId

    /** The grant at the payment provider. Unique, so a retry records one row. */
    var stripeCreditGrantId by ConnectedCreditGrantTable.stripeCreditGrantId
    var amountUsdCents by ConnectedCreditGrantTable.amountUsdCents
    var kind by ConnectedCreditGrantTable.kind

    /**
     * The term this credit was granted for, and when the provider lets it lapse:
     * after the last invoice of that term has been finalized, never on the day
     * the term ends.
     */
    var termEndsAt by ConnectedCreditGrantTable.termEndsAt
    var expiresAt by ConnectedCreditGrantTable.expiresAt
    var createdAt by ConnectedCreditGrantTable.createdAt
}
