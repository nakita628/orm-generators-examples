package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

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
class JoinRequestEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, JoinRequestEntity>(JoinRequestTable)

    var userId by JoinRequestTable.userId
    var organizationId by JoinRequestTable.organizationId

    /** The normalized domain the match was made on. */
    var domain by JoinRequestTable.domain

    /** PENDING | APPROVED | REJECTED | EXPIRED | WITHDRAWN */
    var state by JoinRequestTable.state

    /** How the domain was matched: `verified-identifier-domain` or `sso-connection-domain`. */
    var matchedVia by JoinRequestTable.matchedVia

    /**
     * Where the request was made: `web` or `cli` (ADR-171 v6). A request from
     * the terminal lands a Developer seat whatever the joiner seat says.
     */
    var origin by JoinRequestTable.origin

    /** While PENDING: when the day-14 wake lapses it. Null once resolved. */
    var expiresAt by JoinRequestTable.expiresAt
    var resolvedAt by JoinRequestTable.resolvedAt

    /**
     * user | policy | invite - who ended it. `policy` is domain auto-join's
     * own principal, deliberately not the one SSO auto-join uses.
     */
    var resolvedByType by JoinRequestTable.resolvedByType

    /** The admin's user id, `domain-auto`, or the invitation id. */
    var resolvedById by JoinRequestTable.resolvedById

    /** user | invite-accepted, on a withdrawal. */
    var withdrawalCause by JoinRequestTable.withdrawalCause

    /** The SSO connection a single sign-on arrival was admitted through; null otherwise. */
    var connectionId by JoinRequestTable.connectionId
    var occurredAt by JoinRequestTable.occurredAt
    var lastEventId by JoinRequestTable.lastEventId
    var acceptedAt by JoinRequestTable.acceptedAt
    var projectionVersion by JoinRequestTable.projectionVersion

    /**
     * Event-derived, not database-managed: `@default(now())` / `@updatedAt`
     * would make a replayed row differ from the row it rebuilds, which is
     * exactly the parity this projection has to keep.
     */
    var createdAt by JoinRequestTable.createdAt
    var updatedAt by JoinRequestTable.updatedAt
}
