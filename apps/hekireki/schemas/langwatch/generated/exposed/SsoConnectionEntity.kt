package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * The SSO connection pipeline's projection (ADR-117 §5, D04 - see
 * specs/identity/sso-connection-lifecycle.feature). A pure event-truth head
 * like `Identifier`: every column is fold-written, rows are never deleted
 * (TORN_DOWN is a tombstone), and a replay rebuilds it whole-row from the
 * event log.
 *
 * The fold cursor lives on the row rather than in a sibling table, because
 * one row IS one aggregate here - a connection - so there is nothing for a
 * separate cursor table to key on that this row does not already key on.
 *
 * No relation to Organization on purpose, matching `Identifier`: a
 * projection is rebuilt from events, and a foreign key would make a replay
 * depend on relational state the log does not carry. `idpMetadata` holds
 * endpoints and credential REFERENCES; no secret is ever stored here or in
 * any event.
 */
class SsoConnectionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SsoConnectionEntity>(SsoConnectionTable)

    var organizationId by SsoConnectionTable.organizationId

    /** oidc | saml */
    var type by SsoConnectionTable.type

    /**
     * DRAFT | CLAIMED | APPROVED | REJECTED | DISCARDED |
     * VERIFICATION_PENDING | VERIFIED | ACTIVE | SUSPENDED |
     * TEARDOWN_PENDING | TORN_DOWN
     */
    var state by SsoConnectionTable.state
    var claimedDomains by SsoConnectionTable.claimedDomains

    /**
     * [{ domain, state, claimedAtMs, claimedByActorId, decidedAtMs,
     * decidedByActorId, authority, waitedMs, note }] - every claim this
     * connection has made and how long each waited to be decided. The tier-3
     * approval queue is a read of this column, and the epic's Open Q2 wants
     * that latency measured from the day the queue exists.
     */
    var domainClaims by SsoConnectionTable.domainClaims
    var approvedDomains by SsoConnectionTable.approvedDomains

    /** The only domains that ever route, and only while state is ACTIVE. */
    var verifiedDomains by SsoConnectionTable.verifiedDomains

    /**
     * The subset of `verifiedDomains` whose published record stayed missing
     * through its grace window (ADR-123). They still ROUTE - everyone already
     * here signs in unchanged - and they stop vouching for anybody NEW: no
     * provisioning on first sign-in, no walking in by domain. A column rather
     * than a read of the JSON below, because both questions are asked on a
     * sign-in path and have to be one indexed predicate.
     */
    var lapsedDomains by SsoConnectionTable.lapsedDomains

    /**
     * [{ domain, method, actorId, verifiedAtMs, proofState, firstAbsentAtMs,
     * graceEndsAtMs }] - what proved each verified domain, who, and whether
     * that evidence is still there. Kept on the row rather than only in the
     * log so that an operator-attested domain can never be read as one the
     * customer proved, and a wavering one never as one nothing is doubting.
     */
    var domainVerifications by SsoConnectionTable.domainVerifications

    /** { domain, method, tokenHash } while a ceremony is in flight. */
    var pendingVerification by SsoConnectionTable.pendingVerification
    var idpMetadata by SsoConnectionTable.idpMetadata

    /**
     * admit | request | refuse - what this connection does with somebody who
     * signs in through it and is not a member yet (ADR-117 3). Stated at
     * registration and changed by the setup journey, so it is never absent
     * and no reader has to decide what absence means. This was two fields
     * for a while - a boolean `allowsJit` beside it, kept in step by the
     * fold - and they disagreed about the middle answer.
     */
    var arrivalPolicy by SsoConnectionTable.arrivalPolicy

    /**
     * DEAD, AND STAYING FOR ONE RELEASE. Nothing reads or writes this; the
     * routing repository derives the same answer from `arrivalPolicy`. It is
     * declared so the schema still describes the table, and left in place so
     * the PREVIOUS release's pods — which select it by name on every
     * connection read — keep working through the rollout. Dropping it in the
     * release that stopped reading it would stop single sign-on routing and
     * the connection projection for the length of the deploy. Dropped by
     * `20260827120000_drop_sso_connection_allows_jit` in the next one.
     */
    var allowsJit by SsoConnectionTable.allowsJit

    /**
     * When somebody CHOSE the policy above, and null while the registration
     * default stands. A different fact from the policy rather than a second
     * copy of it: a connection always has a behaviour, and separately
     * somebody has or has not decided on it. Going live waits for the
     * deciding - "turn everybody away" is a decision too.
     */
    var arrivalPolicyDecidedAt by SsoConnectionTable.arrivalPolicyDecidedAt

    /** self-serve | legacy-grandfathered */
    var source by SsoConnectionTable.source1
    var testLoginAccountId by SsoConnectionTable.testLoginAccountId

    /**
     * The grandfathered connection this direct connection replaces. Null for
     * ordinary connections and for the grandfathered predecessor itself.
     *
     * Partial uniqueness for nonterminal replacements is installed by the
     * flattened migration; Prisma cannot express that predicate here.
     */
    var replacesConnectionId by SsoConnectionTable.replacesConnectionId

    /**
     * SETUP | GRACE_LEGACY | GRACE_DIRECT | FINALIZING | FINALIZED.
     * Null means this connection is not a legacy-to-direct replacement.
     */
    var migrationPhase by SsoConnectionTable.migrationPhase
    var graceStartedAt by SsoConnectionTable.graceStartedAt

    /**
     * When an administrator last selected the normal route. This is audit
     * evidence only: routing is derived from `migrationPhase`, never recency.
     */
    var routeChangedAt by SsoConnectionTable.routeChangedAt
    var finalizationRequestedAt by SsoConnectionTable.finalizationRequestedAt
    var finalizedAt by SsoConnectionTable.finalizedAt

    /** { domain, note } - the last ops rejection, read back on a re-claim. */
    var rejection by SsoConnectionTable.rejection
    var createdBy by SsoConnectionTable.createdBy

    /** While TEARDOWN_PENDING: when the process manager's wake completes it. */
    var tearDownAfter by SsoConnectionTable.tearDownAfter
    var occurredAt by SsoConnectionTable.occurredAt
    var lastEventId by SsoConnectionTable.lastEventId
    var acceptedAt by SsoConnectionTable.acceptedAt
    var projectionVersion by SsoConnectionTable.projectionVersion

    /**
     * Event-derived, not database-managed: `@default(now())` / `@updatedAt`
     * would make a replayed row differ from the row it rebuilds, which is
     * exactly the parity this projection has to keep.
     */
    var createdAt by SsoConnectionTable.createdAt
    var updatedAt by SsoConnectionTable.updatedAt
}
