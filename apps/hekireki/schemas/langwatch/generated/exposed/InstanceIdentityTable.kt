package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The identity this install presents to LangWatch (ADR-139, section 2).
 *
 * Exactly one row, minted on first use and never rewritten. It survives
 * restarts, backups and hostname changes, and an install restored into
 * another database mints another one, which is what the registry refuses as
 * the wrong instance.
 *
 * It replaced an identity built from the organization name and id, which
 * leaked the customer's name in cleartext, named one organization rather than
 * one install, and could not be joined to a license row without splitting a
 * string.
 */
object InstanceIdentityTable : IdTable<String>("\"InstanceIdentity\"") {
    /**
     * Fixed, so this table holds one row. A second insert loses the race on the
     * primary key and re-reads the winner's id.
     */
    override val id = text("id").default("self").entityId()

    /** A random UUID. Carries nothing about the customer. */
    val instanceId = text("instanceId").uniqueIndex("\"InstanceIdentity_instanceId_key\"")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /**
     * When the usage report last reached LangWatch, and why the last attempt
     * did not. Both are what the checkup page reads to say whether reporting
     * works, which nothing recorded before: the sender logged success as soon
     * as `fetch` resolved and never looked at the response.
     */
    val lastReportAt = pgTimestamp("lastReportAt", 3).nullable()
    val lastReportError = text("lastReportError").nullable()

    /**
     * The optional category of the usage report, switched off from the checkup
     * page. With it off the report still says which release runs and how many
     * projects the install carries, and says nothing about what anyone does
     * with it (ADR-139, section 10).
     */
    val optionalMetricsOptOut = bool("optionalMetricsOptOut").default(false)

    /**
     * Hostname, which names the customer's own network and so has a switch of
     * its own rather than travelling with its category.
     */
    val hostnameOptOut = bool("hostnameOptOut").default(false)

    /**
     * Sending a failed upgrade step's diagnostics to LangWatch on its own,
     * switched off by the operator (upgrade-diagnostics-via-connect, decision 4).
     */
    val upgradeDiagnosticsOptOut = bool("upgradeDiagnosticsOptOut").default(false)
    override val primaryKey = PrimaryKey(id, name = "\"InstanceIdentity_pkey\"")
}
