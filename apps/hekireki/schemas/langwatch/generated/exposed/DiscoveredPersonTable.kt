package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
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
object DiscoveredPersonTable : IdTable<String>("\"DiscoveredPerson\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")

    /** "anthropic" | "openai" | "databricks" | ... - the provider that named them. */
    val provider = text("provider")

    /**
     * What the provider calls them, verbatim. The join key onto the money rows'
     * RawActorId. Replaced by a deterministic pseudonym on erasure, which keeps
     * the row addressable for spend attribution without keeping the identifier.
     */
    val rawActorId = text("rawActorId")

    /** Name or email text as seen. Pseudonymized alongside rawActorId on erasure. */
    val displayText = text("displayText")

    /**
     * person | service_account - determined deterministically per ADR-128 §10
     * (a bare UUID in Databricks `executed_by` is a service account, never an
     * agent), never guessed from the shape of a name.
     */
    val kind = text("kind")

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
    val department = text("department").nullable()
    val firstSeenAt = pgTimestamp("firstSeenAt", 3)
    val lastSeenAt = pgTimestamp("lastSeenAt", 3)

    /**
     * Set when this row's identity has been erased. The row survives so its
     * spend stays attributable to *someone*; the identifier does not.
     */
    val erasedAt = pgTimestamp("erasedAt", 3).nullable()

    /**
     * Set alongside `erasedAt` and cleared once the daily cost rows have been
     * removed and their rebuild requested. Non-null means the ClickHouse half
     * of the erasure did not finish — a crash between the delete and the
     * replay leaves days short by the erased amount, and without this marker a
     * re-run reads `erasedAt`, short-circuits, and reports a clean erasure that
     * did no work. The re-run resumes at that step instead.
     */
    val moneyRowsPendingAt = pgTimestamp("moneyRowsPendingAt", 3).nullable()

    /**
     * The earliest day to rebuild from, recorded BEFORE the rows are deleted.
     * Once they are gone there is nothing left to ask which days they were on,
     * so a resume with no record of this could only replay everything or
     * nothing. Null means nothing was rebuildable — either no day carried the
     * identifier, or every one of them is older than the event log keeps.
     */
    val moneyRebuildSince = text("moneyRebuildSince").nullable()

    /**
     * ADR-128 §12 conflict rule: automatic linking has stopped for this
     * identity because evidence pointed at two candidates, or contradicted an
     * existing link. Stored HERE rather than on the suggestion rows, because a
     * halt on auto-linking that a recompute or a restart clears is not a halt.
     */
    val suspendedAt = pgTimestamp("suspendedAt", 3).nullable()

    /** What tripped the suspension, for the human who reviews it. */
    val suspendedReason = text("suspendedReason").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"DiscoveredPerson_pkey\"")

    init {
        index("\"DiscoveredPerson_organizationId_lastSeenAt_idx\"", false, organizationId, lastSeenAt)
        uniqueIndex("\"DiscoveredPerson_organizationId_provider_rawActorId_key\"", organizationId, provider, rawActorId)
    }
}
