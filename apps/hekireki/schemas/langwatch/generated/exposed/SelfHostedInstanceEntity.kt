package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class SelfHostedInstanceEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SelfHostedInstanceEntity>(SelfHostedInstanceTable)

    /** The UUID the install minted for itself. The join key for everything else. */
    var instanceId by SelfHostedInstanceTable.instanceId
    var firstSeenAt by SelfHostedInstanceTable.firstSeenAt
    var lastSeenAt by SelfHostedInstanceTable.lastSeenAt

    /** What the install says about itself. */
    var version by SelfHostedInstanceTable.version
    var installMethod by SelfHostedInstanceTable.installMethod
    var chartVersion by SelfHostedInstanceTable.chartVersion

    /** The host it answers on, when the operator left hostname reporting on. */
    var hostname by SelfHostedInstanceTable.hostname
    var environment by SelfHostedInstanceTable.environment

    /**
     * The day the install itself says it first came up, which predates the
     * first report on an install that switched reporting on later.
     */
    var installedAt by SelfHostedInstanceTable.installedAt
    var reportSchemaVersion by SelfHostedInstanceTable.reportSchemaVersion

    /**
     * The customer, and the license that named them. Plain column refs rather
     * than relations, like `IssuedLicense.replacesId`: relationMode="prisma"
     * carries no foreign keys, and the receiver resolves both from the license
     * registry rather than from the report.
     */
    var organizationId by SelfHostedInstanceTable.organizationId
    var issuedLicenseId by SelfHostedInstanceTable.issuedLicenseId

    /**
     * Domains with counts, as `{"acme.test": 14}`. Never an address. Null on an
     * install that switched the optional category off.
     */
    var userEmailDomains by SelfHostedInstanceTable.userEmailDomains

    /**
     * The same domains as the keys of `userEmailDomains`, flattened so an
     * operator can search for one. A JSON object's keys are the domains
     * themselves, and no query language addresses a key it does not already
     * know the name of.
     */
    var userDomains by SelfHostedInstanceTable.userDomains

    /**
     * What the last report carried, so the list and the drawer read one row:
     * the onboarding ladder, the usage counts and the operational block, under
     * the same keys the dictionary declares.
     */
    var latestReport by SelfHostedInstanceTable.latestReport

    /**
     * What the last report said the operator had switched off, read from which
     * categories arrived rather than from a field the install could assert.
     */
    var optionalMetricsReported by SelfHostedInstanceTable.optionalMetricsReported
    var hostnameReported by SelfHostedInstanceTable.hostnameReported

    /**
     * How many reports have landed, and how many fields the last one carried
     * that this release has no name for. A count above zero says the senders
     * are ahead of the receiver.
     */
    var reportCount by SelfHostedInstanceTable.reportCount
    var lastUnknownFields by SelfHostedInstanceTable.lastUnknownFields

    /**
     * The lead signals already raised for this install. A signal on this list
     * is never raised again, which is what keeps a large install from
     * reporting that it is large every morning.
     */
    var raisedSignals by SelfHostedInstanceTable.raisedSignals
    var createdAt by SelfHostedInstanceTable.createdAt
    var updatedAt by SelfHostedInstanceTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SelfHostedInstanceTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
