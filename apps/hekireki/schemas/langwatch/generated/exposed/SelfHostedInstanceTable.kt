package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One self-hosted install, as its daily usage report describes it (ADR-139,
 * section 10). The row is created by the first report that names an instance
 * id and is updated by every report after that, so the backoffice reads one
 * row per install rather than folding a history on every page load.
 *
 * The report carries no credential, so it identifies an install and never a
 * customer. `organizationId` and `issuedLicenseId` are resolved from the
 * license bound to this instance, which the install proved by presenting its
 * token to license sync. A field in the posted body never reaches them.
 */
object SelfHostedInstanceTable : IdTable<String>("\"SelfHostedInstance\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()

    /** The UUID the install minted for itself. The join key for everything else. */
    val instanceId = text("instanceId").uniqueIndex("\"SelfHostedInstance_instanceId_key\"")
    val firstSeenAt = pgTimestamp("firstSeenAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val lastSeenAt = pgTimestamp("lastSeenAt", 3)

    /** What the install says about itself. */
    val version = text("version").nullable()
    val installMethod = text("installMethod").nullable()
    val chartVersion = text("chartVersion").nullable()

    /** The host it answers on, when the operator left hostname reporting on. */
    val hostname = text("hostname").nullable()
    val environment = text("environment").nullable()

    /**
     * The day the install itself says it first came up, which predates the
     * first report on an install that switched reporting on later.
     */
    val installedAt = pgTimestamp("installedAt", 3).nullable()
    val reportSchemaVersion = integer("reportSchemaVersion").nullable()

    /**
     * The customer, and the license that named them. Plain column refs rather
     * than relations, like `IssuedLicense.replacesId`: relationMode="prisma"
     * carries no foreign keys, and the receiver resolves both from the license
     * registry rather than from the report.
     */
    val organizationId = text("organizationId").nullable()
    val issuedLicenseId = text("issuedLicenseId").nullable()

    /**
     * Domains with counts, as `{"acme.test": 14}`. Never an address. Null on an
     * install that switched the optional category off.
     */
    val userEmailDomains = jsonb<String>("userEmailDomains", { it }, { it }).nullable()

    /**
     * The same domains as the keys of `userEmailDomains`, flattened so an
     * operator can search for one. A JSON object's keys are the domains
     * themselves, and no query language addresses a key it does not already
     * know the name of.
     */
    val userDomains = array("userDomains", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /**
     * What the last report carried, so the list and the drawer read one row:
     * the onboarding ladder, the usage counts and the operational block, under
     * the same keys the dictionary declares.
     */
    val latestReport = jsonb<String>("latestReport", { it }, { it }).nullable()

    /**
     * What the last report said the operator had switched off, read from which
     * categories arrived rather than from a field the install could assert.
     */
    val optionalMetricsReported = bool("optionalMetricsReported").default(true)
    val hostnameReported = bool("hostnameReported").default(true)

    /**
     * How many reports have landed, and how many fields the last one carried
     * that this release has no name for. A count above zero says the senders
     * are ahead of the receiver.
     */
    val reportCount = integer("reportCount").default(0)
    val lastUnknownFields = integer("lastUnknownFields").default(0)

    /**
     * The lead signals already raised for this install. A signal on this list
     * is never raised again, which is what keeps a large install from
     * reporting that it is large every morning.
     */
    val raisedSignals = array("raisedSignals", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"SelfHostedInstance_pkey\"")

    init {
        index("\"SelfHostedInstance_lastSeenAt_idx\"", false, lastSeenAt)
        index("\"SelfHostedInstance_organizationId_idx\"", false, organizationId)
        index("\"SelfHostedInstance_issuedLicenseId_idx\"", false, issuedLicenseId)
    }
}
