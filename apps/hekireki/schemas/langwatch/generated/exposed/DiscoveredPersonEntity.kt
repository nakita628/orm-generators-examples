package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A name-text seen on a provider's cost or audit rows: "the provider told us
 * someone called m.silva exists". One row per provider-scoped identity.
 *
 * This is the unit of GDPR erasure (ADR-128 §9), and deliberately so: most of
 * these people have no LangWatch login at all - contractors, seat holders who
 * never signed in, anyone whose email a provider put on a cost row - so there
 * is no platform user to key on and a user-deletion event can never fire for
 * them.
 */
class DiscoveredPersonEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DiscoveredPersonEntity>(DiscoveredPersonTable)

    var organizationId by DiscoveredPersonTable.organizationId

    /** "anthropic" | "openai" | "databricks" | ... - the provider that named them. */
    var provider by DiscoveredPersonTable.provider

    /**
     * What the provider calls them, verbatim. The join key onto the money rows'
     * RawActorId. Replaced by a deterministic pseudonym on erasure, which keeps
     * the row addressable for spend attribution without keeping the identifier.
     */
    var rawActorId by DiscoveredPersonTable.rawActorId

    /** Name or email text as seen. Pseudonymized alongside rawActorId on erasure. */
    var displayText by DiscoveredPersonTable.displayText

    /**
     * person | service_account - determined deterministically per ADR-128 §10
     * (a bare UUID in Databricks `executed_by` is a service account, never an
     * agent), never guessed from the shape of a name.
     */
    var kind by DiscoveredPersonTable.kind

    /**
     * The department the provider's directory filed this person under, verbatim.
     * An OBSERVATION, not the `Department` entity: it is free text a tenant owns
     * on their side, it exists for people who hold no LangWatch account, and
     * nothing rolls spend up by it. Spend attribution stays on `Department`, via
     * the membership the match engine proves (directoryDepartmentSync.service).
     * Null means no directory has named one - never "no department", because a
     * blank field in the tenant's directory and a person no directory lists are
     * the same null here.
     */
    var department by DiscoveredPersonTable.department
    var firstSeenAt by DiscoveredPersonTable.firstSeenAt
    var lastSeenAt by DiscoveredPersonTable.lastSeenAt

    /**
     * Set when this row's identity has been erased. The row survives so its
     * spend stays attributable to *someone*; the identifier does not.
     */
    var erasedAt by DiscoveredPersonTable.erasedAt

    /**
     * Set alongside `erasedAt` and cleared once the daily cost rows have been
     * removed and their rebuild requested. Non-null means the ClickHouse half
     * of the erasure did not finish — a crash between the delete and the
     * replay leaves days short by the erased amount, and without this marker a
     * re-run reads `erasedAt`, short-circuits, and reports a clean erasure that
     * did no work. The re-run resumes at that step instead.
     */
    var moneyRowsPendingAt by DiscoveredPersonTable.moneyRowsPendingAt

    /**
     * The earliest day to rebuild from, recorded BEFORE the rows are deleted.
     * Once they are gone there is nothing left to ask which days they were on,
     * so a resume with no record of this could only replay everything or
     * nothing. Null means nothing was rebuildable — either no day carried the
     * identifier, or every one of them is older than the event log keeps.
     */
    var moneyRebuildSince by DiscoveredPersonTable.moneyRebuildSince

    /**
     * ADR-128 §12 conflict rule: automatic linking has stopped for this
     * identity because evidence pointed at two candidates, or contradicted an
     * existing link. Stored HERE rather than on the suggestion rows, because a
     * halt on auto-linking that a recompute or a restart clears is not a halt.
     */
    var suspendedAt by DiscoveredPersonTable.suspendedAt

    /** What tripped the suspension, for the human who reviews it. */
    var suspendedReason by DiscoveredPersonTable.suspendedReason
    var createdAt by DiscoveredPersonTable.createdAt
    var updatedAt by DiscoveredPersonTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(DiscoveredPersonTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
