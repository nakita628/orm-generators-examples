package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * The join-request pipeline's projection (ADR-117, D12 - see
 * specs/identity/join-requests.feature). A pure event-truth head like
 * `Identifier` and `SsoConnection`: every column is fold-written, rows are
 * never deleted (a terminal state is a tombstone the admin panel stops
 * listing, not an absence), and a replay rebuilds it whole-row from the
 * event log.
 *
 * The fold cursor lives on the row rather than in a sibling table, because
 * one row IS one aggregate here - a request.
 *
 * No relation to Organization or User on purpose, matching the two heads
 * above: a projection is rebuilt from events, and a foreign key would make a
 * replay depend on relational state the log does not carry.
 *
 * Membership is NOT here. An approval dispatches an attach on the grants
 * ledger, exactly as accepting an invitation does, and that ledger fact
 * carries `source: "join-request"` as its provenance - which is what puts a
 * join approval on the customer's audit page beside every other membership
 * change rather than in a history of its own.
 *
 * The requester's address is not here either: the DOMAIN is the fact, and it
 * is the only part of an address a join decision may depend on.
 */
object JoinRequestTable : IdTable<String>("\"JoinRequest\"") {
    override val id = text("id").entityId()
    val userId = text("userId")
    val organizationId = text("organizationId")

    /** The normalized domain the match was made on. */
    val domain = text("domain")

    /** PENDING | APPROVED | REJECTED | EXPIRED | WITHDRAWN */
    val state = text("state")

    /** How the domain was matched: `verified-identifier-domain` or `sso-connection-domain`. */
    val matchedVia = text("matchedVia")

    /**
     * Where the request was made: `web` or `cli` (ADR-171 v6). A request from
     * the terminal lands a Developer seat whatever the joiner seat says.
     */
    val origin = text("origin").default("web")

    /** While PENDING: when the day-14 wake lapses it. Null once resolved. */
    val expiresAt = pgTimestamp("expiresAt", 3).nullable()
    val resolvedAt = pgTimestamp("resolvedAt", 3).nullable()

    /**
     * user | policy | invite - who ended it. `policy` is domain auto-join's
     * own principal, deliberately not the one SSO auto-join uses.
     */
    val resolvedByType = text("resolvedByType").nullable()

    /** The admin's user id, `domain-auto`, or the invitation id. */
    val resolvedById = text("resolvedById").nullable()

    /** user | invite-accepted, on a withdrawal. */
    val withdrawalCause = text("withdrawalCause").nullable()

    /** The SSO connection a single sign-on arrival was admitted through; null otherwise. */
    val connectionId = text("connectionId").nullable()
    val occurredAt = pgTimestamp("occurredAt", 3)
    val lastEventId = text("lastEventId")
    val acceptedAt = pgTimestamp("acceptedAt", 3)
    val projectionVersion = text("projectionVersion")

    /**
     * Event-derived, not database-managed: `@default(now())` / `@updatedAt`
     * would make a replayed row differ from the row it rebuilds, which is
     * exactly the parity this projection has to keep.
     */
    val createdAt = pgTimestamp("createdAt", 3)
    val updatedAt = pgTimestamp("updatedAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"JoinRequest_pkey\"")

    init {
        index("\"JoinRequest_organizationId_state_idx\"", false, organizationId, state)
        index("\"JoinRequest_userId_state_idx\"", false, userId, state)
        index("\"JoinRequest_userId_organizationId_idx\"", false, userId, organizationId)
        index("\"JoinRequest_domain_idx\"", false, domain)
    }
}
