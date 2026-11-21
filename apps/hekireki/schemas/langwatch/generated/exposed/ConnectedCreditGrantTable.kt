package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A paid credit at the payment provider, which only metered usage draws down.
 * Kept here so a grant is created once per reason and the backoffice can show
 * what the customer holds without asking the provider.
 */
object ConnectedCreditGrantTable : IdTable<String>("\"ConnectedCreditGrant\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()

    /**
     * The `ConnectedBillingAccount` row. A plain column ref, like
     * `IssuedLicense.replacesId`: relationMode="prisma" carries no foreign keys.
     */
    val accountId = text("accountId")

    /** The grant at the payment provider. Unique, so a retry records one row. */
    val stripeCreditGrantId = text("stripeCreditGrantId").uniqueIndex("\"ConnectedCreditGrant_stripeCreditGrantId_key\"")
    val amountUsdCents = integer("amountUsdCents")
    val kind = pgEnum("kind", "\"ConnectedCreditGrantKind\"", ConnectedCreditGrantKind.entries, ConnectedCreditGrantKind::dbName)

    /**
     * The term this credit was granted for, and when the provider lets it lapse:
     * after the last invoice of that term has been finalized, never on the day
     * the term ends.
     */
    val termEndsAt = pgTimestamp("termEndsAt", 3)
    val expiresAt = pgTimestamp("expiresAt", 3)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"ConnectedCreditGrant_pkey\"")

    init {
        index("\"ConnectedCreditGrant_accountId_idx\"", false, accountId)
    }
}
