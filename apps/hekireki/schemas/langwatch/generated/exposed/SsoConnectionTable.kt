package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

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
object SsoConnectionTable : IdTable<String>("\"SsoConnection\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")

    /** oidc | saml */
    val type = text("type")

    /**
     * DRAFT | CLAIMED | APPROVED | REJECTED | DISCARDED |
     * VERIFICATION_PENDING | VERIFIED | ACTIVE | SUSPENDED |
     * TEARDOWN_PENDING | TORN_DOWN
     */
    val state = text("state")
    val claimedDomains = array("claimedDomains", TextColumnType()).nullable()

    /**
     * [{ domain, state, claimedAtMs, claimedByActorId, decidedAtMs,
     * decidedByActorId, authority, waitedMs, note }] - every claim this
     * connection has made and how long each waited to be decided. The tier-3
     * approval queue is a read of this column, and the epic's Open Q2 wants
     * that latency measured from the day the queue exists.
     */
    val domainClaims = jsonb<String>("domainClaims", { it }, { it }).default("[]")
    val approvedDomains = array("approvedDomains", TextColumnType()).nullable()

    /** The only domains that ever route, and only while state is ACTIVE. */
    val verifiedDomains = array("verifiedDomains", TextColumnType()).nullable()

    /**
     * The subset of `verifiedDomains` whose published record stayed missing
     * through its grace window (ADR-123). They still ROUTE - everyone already
     * here signs in unchanged - and they stop vouching for anybody NEW: no
     * provisioning on first sign-in, no walking in by domain. A column rather
     * than a read of the JSON below, because both questions are asked on a
     * sign-in path and have to be one indexed predicate.
     */
    val lapsedDomains = array("lapsedDomains", TextColumnType()).nullable()

    /**
     * [{ domain, method, actorId, verifiedAtMs, proofState, firstAbsentAtMs,
     * graceEndsAtMs }] - what proved each verified domain, who, and whether
     * that evidence is still there. Kept on the row rather than only in the
     * log so that an operator-attested domain can never be read as one the
     * customer proved, and a wavering one never as one nothing is doubting.
     */
    val domainVerifications = jsonb<String>("domainVerifications", { it }, { it }).default("[]")

    /** { domain, method, tokenHash } while a ceremony is in flight. */
    val pendingVerification = jsonb<String>("pendingVerification", { it }, { it }).nullable()
    val idpMetadata = jsonb<String>("idpMetadata", { it }, { it })

    /**
     * admit | request | refuse - what this connection does with somebody who
     * signs in through it and is not a member yet (ADR-117 3). Stated at
     * registration and changed by the setup journey, so it is never absent
     * and no reader has to decide what absence means. This was two fields
     * for a while - a boolean `allowsJit` beside it, kept in step by the
     * fold - and they disagreed about the middle answer.
     */
    val arrivalPolicy = text("arrivalPolicy").default("refuse")

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
    val allowsJit = bool("allowsJit").default(false)

    /**
     * When somebody CHOSE the policy above, and null while the registration
     * default stands. A different fact from the policy rather than a second
     * copy of it: a connection always has a behaviour, and separately
     * somebody has or has not decided on it. Going live waits for the
     * deciding - "turn everybody away" is a decision too.
     */
    val arrivalPolicyDecidedAt = pgTimestamp("arrivalPolicyDecidedAt", 3).nullable()

    /** self-serve | legacy-grandfathered */
    val source1 = text("source")
    val testLoginAccountId = text("testLoginAccountId").nullable()

    /**
     * The grandfathered connection this direct connection replaces. Null for
     * ordinary connections and for the grandfathered predecessor itself.
     *
     * Partial uniqueness for nonterminal replacements is installed by the
     * flattened migration; Prisma cannot express that predicate here.
     */
    val replacesConnectionId = text("replacesConnectionId").nullable()

    /**
     * SETUP | GRACE_LEGACY | GRACE_DIRECT | FINALIZING | FINALIZED.
     * Null means this connection is not a legacy-to-direct replacement.
     */
    val migrationPhase =
        pgEnum(
            "migrationPhase",
            "\"SsoConnectionMigrationPhase\"",
            SsoConnectionMigrationPhase.entries,
            SsoConnectionMigrationPhase::dbName,
        ).nullable()
    val graceStartedAt = pgTimestamp("graceStartedAt", 3).nullable()

    /**
     * When an administrator last selected the normal route. This is audit
     * evidence only: routing is derived from `migrationPhase`, never recency.
     */
    val routeChangedAt = pgTimestamp("routeChangedAt", 3).nullable()
    val finalizationRequestedAt = pgTimestamp("finalizationRequestedAt", 3).nullable()
    val finalizedAt = pgTimestamp("finalizedAt", 3).nullable()

    /** { domain, note } - the last ops rejection, read back on a re-claim. */
    val rejection = jsonb<String>("rejection", { it }, { it }).nullable()
    val createdBy = text("createdBy").nullable()

    /** While TEARDOWN_PENDING: when the process manager's wake completes it. */
    val tearDownAfter = pgTimestamp("tearDownAfter", 3).nullable()
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
    override val primaryKey = PrimaryKey(id, name = "\"SsoConnection_pkey\"")

    init {
        index("\"SsoConnection_organizationId_idx\"", false, organizationId)
        index("\"SsoConnection_organizationId_migrationPhase_idx\"", false, organizationId, migrationPhase)
        index("\"SsoConnection_state_idx\"", false, state)
        index("\"SsoConnection_verifiedDomains_idx\"", false, verifiedDomains, indexType = "gin")
        index("\"SsoConnection_lapsedDomains_idx\"", false, lapsedDomains, indexType = "gin")
    }
}
