package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * The monthly usage statement one connected customer was sent (ADR-139,
 * section 7). One row per account per month keeps the statement to one send,
 * whatever runs the tick.
 */
object ConnectedStatementTable : IdTable<String>("\"ConnectedStatement\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()

    /** The `ConnectedBillingAccount` row. A plain column ref. */
    val accountId = text("accountId")

    /** First day of the month the statement covers, in UTC. */
    val month = pgTimestamp("month", 3)
    val sentAt = pgTimestamp("sentAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"ConnectedStatement_pkey\"")

    init {
        index("\"ConnectedStatement_accountId_idx\"", false, accountId)
        uniqueIndex("\"ConnectedStatement_accountId_month_key\"", accountId, month)
    }
}
