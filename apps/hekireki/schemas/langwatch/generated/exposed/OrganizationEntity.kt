package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class OrganizationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OrganizationEntity>(OrganizationTable)

    var name by OrganizationTable.name

    /**
     * `customer` for every organization a person founds or an operator provisions; `system` for
     * the install's one hidden organization the Ops checkup canaries run in, never listed to a person.
     * Spec: modules/ops/specs/checkup-canary-project.feature.
     */
    var kind by OrganizationTable.kind
    var phoneNumber by OrganizationTable.phoneNumber
    var slug by OrganizationTable.slug
    var createdAt by OrganizationTable.createdAt
    var updatedAt by OrganizationTable.updatedAt
    var usageSpendingMaxLimit by OrganizationTable.usageSpendingMaxLimit

    /**
     * The largest file a dataset image or file cell accepts for this
     * organization, in mebibytes. Null means the default. Set by an operator
     * from the back office only; every other dataset size limit scales with it.
     */
    var datasetAttachmentMaxMb by OrganizationTable.datasetAttachmentMaxMb

    /**
     * The label of this organization's `CLICKHOUSE_URL__<label>__<id>` route, recorded by the
     * organization module the first time it is seen. Once set, a missing route refuses the
     * organization instead of serving it from the shared ClickHouse. Null: never private.
     * Spec: specs/private-dataplane/clickhouse-routing.feature.
     */
    var privateClickHouseLabel by OrganizationTable.privateClickHouseLabel

    /** The same for `DATAPLANE_S3__<label>__<id>`. Spec: specs/private-dataplane/s3-routing.feature. */
    var privateS3Label by OrganizationTable.privateS3Label

    /** When either label was first recorded. */
    var privateDataplaneSeenAt by OrganizationTable.privateDataplaneSeenAt

    /**
     * Phase 8 - admin-configurable max lifetime for CLI device-flow
     * sessions, in days. 0 = unbounded (default, mirrors GitHub CLI /
     * gh-style flows). Companies that need finite session lifetime
     * (eg. SOC 2 / ISO 27001 customers with mandated session
     * expiration) can set a positive value. The /api/auth/cli/refresh
     * endpoint rejects refresh attempts where now - issued_at exceeds
     * this value, regardless of the underlying refresh-token TTL.
     * Spec: specs/ai-governance/sessions/session-ttl.feature.
     */
    var maxSessionDurationDays by OrganizationTable.maxSessionDurationDays

    /**
     * Every member of this organization can prove a second factor (D06). A
     * MEMBERSHIP CONDITION, not a session policy: it is asked when a member
     * reaches this organization's data, it holds the ones who cannot prove
     * one at an enrollment gate for THIS organization alone, and it ends no
     * session, ever. Turning it on signs nobody out.
     * Spec: specs/identity/mfa-and-session-shape.feature.
     */
    var mfaRequired by OrganizationTable.mfaRequired

    /**
     * Whether members may connect the documentation playground to their own access with a
     * short-lived docs session key (ADR-185, amendment of 2026-11-13). On by default; turned off,
     * no key is minted and live ones stop authenticating. The pasted-key path is untouched.
     * Spec: modules/api-key/specs/docs-session-keys.feature.
     */
    var docsPlaygroundSessionsEnabled by OrganizationTable.docsPlaygroundSessionsEnabled

    /**
     * When the docs playground was last turned off, written in the same update as the switch. A
     * docs session key minted at or before it never authenticates again, nor any key minted under
     * one, so turning it back on revives none of them (ADR-185, amendment of 2026-11-14).
     */
    var docsPlaygroundSessionsDisabledAt by OrganizationTable.docsPlaygroundSessionsDisabledAt

    /**
     * How many consecutive failed sign-ins lock an account (GAC-09).
     * 0 = never lock, which is the default and what every organization that
     * has not asked for this keeps. Counts EVERY credential together - a
     * password, a one-time code, a backup code, a passkey - because five per
     * factor hands an attacker one budget per factor.
     * Spec: specs/identity/org-account-lockout.feature.
     */
    var lockoutAfterFailedAttempts by OrganizationTable.lockoutAfterFailedAttempts

    /** How long a lock lasts, in minutes. Ignored while the threshold is 0. */
    var lockoutMinutes by OrganizationTable.lockoutMinutes

    /**
     * How long a BROWSER session may sit idle before it ends, in minutes
     * (GAC-10). 0 = no idle timeout. Activity rolls the window, which is the
     * "refreshed every hour" half of the control.
     *
     * A day for organizations created from now on; organizations that predate
     * the column keep 0 and are unchanged, because the migration adds the
     * column at zero and moves the default afterwards.
     *
     * NOT `maxSessionDurationDays`, which bounds CLI and device sessions and
     * is enforced when the CLI refreshes. Two surfaces, two lifetimes, two
     * decisions an administrator makes separately.
     * Spec: specs/identity/org-session-lifetime.feature.
     */
    var sessionIdleTimeoutMinutes by OrganizationTable.sessionIdleTimeoutMinutes

    /**
     * The ceiling on a browser session from the moment it was minted, in
     * minutes, which no amount of activity extends. 0 = no ceiling, the
     * default. Offered beside the idle timeout because a session kept alive
     * forever by a polling page is not really bounded.
     */
    var sessionMaxLifetimeMinutes by OrganizationTable.sessionMaxLifetimeMinutes
    var signupData by OrganizationTable.signupData
    var signedDPA by OrganizationTable.signedDPA
    var elasticsearchNodeUrl by OrganizationTable.elasticsearchNodeUrl
    var elasticsearchApiKey by OrganizationTable.elasticsearchApiKey
    var useCustomElasticsearch by OrganizationTable.useCustomElasticsearch
    var s3Endpoint by OrganizationTable.s3Endpoint
    var s3AccessKeyId by OrganizationTable.s3AccessKeyId
    var s3SecretAccessKey by OrganizationTable.s3SecretAccessKey
    var s3Bucket by OrganizationTable.s3Bucket
    var useCustomS3 by OrganizationTable.useCustomS3
    var sentPlanLimitAlert by OrganizationTable.sentPlanLimitAlert
    var ssoDomain by OrganizationTable.ssoDomain
    var ssoProvider by OrganizationTable.ssoProvider

    /**
     * How colleagues on a matching domain get in (D12, ADR-117):
     * `off` (invisible to them entirely) | `request` (they ask, an admin
     * approves) | `auto` (a verified colleague walks straight in).
     *
     * `request` is the default for a self-serve organization; `auto` is never
     * a default and never inferred - an administrator turns it on, names the
     * domain in `joinDomains` while doing it, and the licence gate holds that
     * value while letting `request` through.
     */
    var domainJoin by OrganizationTable.domainJoin

    /**
     * The organisation role handed to a person admitted without an invitation:
     * a self-service domain join (`domainJoin`) or an SSO-admitted login.
     * `MEMBER` (default, a Full seat) or `DEVELOPER` (ADR-171). Set on the
     * organisation settings page only; the join flow never asks, and an
     * invitation always names its own role.
     */
    var joinerRole by OrganizationTable.joinerRole

    /**
     * The domains an administrator NAMED when turning automatic joining on.
     * Empty means automatic joining admits nobody whatever `domainJoin` says -
     * which is what keeps a setting left on from admitting a domain nobody
     * ever agreed to.
     */
    var joinDomains by OrganizationTable.joinDomains
    var presenceEnabled by OrganizationTable.presenceEnabled

    /**
     * Org-level trace-sharing kill switch. Off = trace share links cannot be
     * created in ANY project of this org, and existing links stop resolving —
     * the org toggle is the global switch, the per-project
     * `Project.traceSharingEnabled` scopes the kill to a single project (mirrors
     * `presenceEnabled`). Effective sharing = org AND project. ADR-039.
     */
    var traceSharingEnabled by OrganizationTable.traceSharingEnabled

    /**
     * Org-configurable contact surfaced to users in CLI "contact your
     * admin" messages and the in-app budget-exceeded banner. Free text:
     * accepts an email, a URL (internal ticketing system), or any other
     * instruction. When null, callers fall back to the first ADMIN
     * member's email (`resolveSupportContact` helper).
     */
    var supportContact by OrganizationTable.supportContact

    /**
     * Signup intent, editable later as the org "Primary use" setting.
     * AGENT_GOVERNANCE lands "/" on /me, LLM_OPS on the project home.
     * NULL means intent unset (every pre-fork org and any org created
     * outside onboarding) — the home resolver keeps legacy behavior.
     * ADR-038.
     */
    var primaryIntent by OrganizationTable.primaryIntent
    var promoCode by OrganizationTable.promoCode
    var stripeCustomerId by OrganizationTable.stripeCustomerId
    var currency by OrganizationTable.currency
    var pricingModel by OrganizationTable.pricingModel
    var license by OrganizationTable.license
    var licenseExpiresAt by OrganizationTable.licenseExpiresAt
    var licenseLastValidatedAt by OrganizationTable.licenseLastValidatedAt

    /**
     * A customer that runs LangWatch self-hosted. The hosted usage, budgets and
     * invoices of its connected installs attach to this organization (ADR-139).
     */
    var selfHostedCustomer by OrganizationTable.selfHostedCustomer

    /**
     * LangWatch-hosted services this organization has switched OFF, named the
     * way its license names them ("instant_evals"). A service the license is
     * entitled to is on unless it is listed here, so a customer who bought
     * hosted judging gets it without finding a settings page first, and an
     * administrator who switched one off keeps it off across license changes
     * (ADR-139, section 2). A license that names no service reaches nothing,
     * whatever this column holds.
     */
    var connectServicesDisabled by OrganizationTable.connectServicesDisabled

    /**
     * When a member of this organization switched Instant Evals on from the
     * search bar, which is the organization's consent to have judged text sent
     * to the judge's provider. Null until then: a self-serve organization is
     * never opted in on its behalf, and an enterprise one is switched on by an
     * operator through the release flag after asking, never here.
     */
    var instantEvalsEnabledAt by OrganizationTable.instantEvalsEnabledAt

    /**
     * The member who switched them on, so the audit trail names a person and
     * not just a moment.
     */
    var instantEvalsEnabledByUserId by OrganizationTable.instantEvalsEnabledByUserId

    /** When the last license sync succeeded. */
    var connectLastSyncAt by OrganizationTable.connectLastSyncAt

    /**
     * The code of the failure the last sync ended on, cleared on the next
     * success. Settings, Connect renders the copy for it.
     */
    var connectLastSyncError by OrganizationTable.connectLastSyncError

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(OrganizationTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
