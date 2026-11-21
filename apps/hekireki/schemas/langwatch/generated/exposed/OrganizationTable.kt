package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object OrganizationTable : IdTable<String>("\"Organization\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val name = text("name")

    /**
     * `customer` for every organization a person founds or an operator provisions; `system` for
     * the install's one hidden organization the Ops checkup canaries run in, never listed to a person.
     * Spec: modules/ops/specs/checkup-canary-project.feature.
     */
    val kind = text("kind").default("customer")
    val phoneNumber = text("phoneNumber").nullable()
    val slug = text("slug").uniqueIndex("\"Organization_slug_key\"")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val usageSpendingMaxLimit = integer("usageSpendingMaxLimit").nullable()

    /**
     * The largest file a dataset image or file cell accepts for this
     * organization, in mebibytes. Null means the default. Set by an operator
     * from the back office only; every other dataset size limit scales with it.
     */
    val datasetAttachmentMaxMb = integer("datasetAttachmentMaxMb").nullable()

    /**
     * The label of this organization's `CLICKHOUSE_URL__<label>__<id>` route, recorded by the
     * organization module the first time it is seen. Once set, a missing route refuses the
     * organization instead of serving it from the shared ClickHouse. Null: never private.
     * Spec: specs/private-dataplane/clickhouse-routing.feature.
     */
    val privateClickHouseLabel = text("privateClickHouseLabel").nullable()

    /** The same for `DATAPLANE_S3__<label>__<id>`. Spec: specs/private-dataplane/s3-routing.feature. */
    val privateS3Label = text("privateS3Label").nullable()

    /** When either label was first recorded. */
    val privateDataplaneSeenAt = pgTimestamp("privateDataplaneSeenAt", 3).nullable()

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
    val maxSessionDurationDays = integer("maxSessionDurationDays").default(0)

    /**
     * Every member of this organization can prove a second factor (D06). A
     * MEMBERSHIP CONDITION, not a session policy: it is asked when a member
     * reaches this organization's data, it holds the ones who cannot prove
     * one at an enrollment gate for THIS organization alone, and it ends no
     * session, ever. Turning it on signs nobody out.
     * Spec: specs/identity/mfa-and-session-shape.feature.
     */
    val mfaRequired = bool("mfaRequired").default(false)

    /**
     * Whether members may connect the documentation playground to their own access with a
     * short-lived docs session key (ADR-185, amendment of 2026-11-13). On by default; turned off,
     * no key is minted and live ones stop authenticating. The pasted-key path is untouched.
     * Spec: modules/api-key/specs/docs-session-keys.feature.
     */
    val docsPlaygroundSessionsEnabled = bool("docsPlaygroundSessionsEnabled").default(true)

    /**
     * When the docs playground was last turned off, written in the same update as the switch. A
     * docs session key minted at or before it never authenticates again, nor any key minted under
     * one, so turning it back on revives none of them (ADR-185, amendment of 2026-11-14).
     */
    val docsPlaygroundSessionsDisabledAt = pgTimestamp("docsPlaygroundSessionsDisabledAt", 3).nullable()

    /**
     * How many consecutive failed sign-ins lock an account (GAC-09).
     * 0 = never lock, which is the default and what every organization that
     * has not asked for this keeps. Counts EVERY credential together - a
     * password, a one-time code, a backup code, a passkey - because five per
     * factor hands an attacker one budget per factor.
     * Spec: specs/identity/org-account-lockout.feature.
     */
    val lockoutAfterFailedAttempts = integer("lockoutAfterFailedAttempts").default(0)

    /** How long a lock lasts, in minutes. Ignored while the threshold is 0. */
    val lockoutMinutes = integer("lockoutMinutes").default(30)

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
    val sessionIdleTimeoutMinutes = integer("sessionIdleTimeoutMinutes").default(1440)

    /**
     * The ceiling on a browser session from the moment it was minted, in
     * minutes, which no amount of activity extends. 0 = no ceiling, the
     * default. Offered beside the idle timeout because a session kept alive
     * forever by a polling page is not really bounded.
     */
    val sessionMaxLifetimeMinutes = integer("sessionMaxLifetimeMinutes").default(0)
    val signupData = jsonb<String>("signupData", { it }, { it }).nullable()
    val signedDPA = bool("signedDPA").default(false)
    val elasticsearchNodeUrl = text("elasticsearchNodeUrl").nullable()
    val elasticsearchApiKey = text("elasticsearchApiKey").nullable()
    val useCustomElasticsearch = bool("useCustomElasticsearch").default(false)
    val s3Endpoint = text("s3Endpoint").nullable()
    val s3AccessKeyId = text("s3AccessKeyId").nullable()
    val s3SecretAccessKey = text("s3SecretAccessKey").nullable()
    val s3Bucket = text("s3Bucket").nullable()
    val useCustomS3 = bool("useCustomS3").default(false)
    val sentPlanLimitAlert = pgTimestamp("sentPlanLimitAlert", 3).nullable()
    val ssoDomain = text("ssoDomain").nullable().uniqueIndex("\"Organization_ssoDomain_key\"")
    val ssoProvider = text("ssoProvider").nullable()

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
    val domainJoin = text("domainJoin").default("request")

    /**
     * The organisation role handed to a person admitted without an invitation:
     * a self-service domain join (`domainJoin`) or an SSO-admitted login.
     * `MEMBER` (default, a Full seat) or `DEVELOPER` (ADR-171). Set on the
     * organisation settings page only; the join flow never asks, and an
     * invitation always names its own role.
     */
    val joinerRole =
        pgEnum("joinerRole", "\"OrganizationUserRole\"", OrganizationUserRole.entries, OrganizationUserRole::dbName).default(
            OrganizationUserRole.MEMBER,
        )

    /**
     * The domains an administrator NAMED when turning automatic joining on.
     * Empty means automatic joining admits nobody whatever `domainJoin` says -
     * which is what keeps a setting left on from admitting a domain nobody
     * ever agreed to.
     */
    val joinDomains = array("joinDomains", TextColumnType()).nullable()
    val presenceEnabled = bool("presenceEnabled").default(true)

    /**
     * Org-level trace-sharing kill switch. Off = trace share links cannot be
     * created in ANY project of this org, and existing links stop resolving —
     * the org toggle is the global switch, the per-project
     * `Project.traceSharingEnabled` scopes the kill to a single project (mirrors
     * `presenceEnabled`). Effective sharing = org AND project. ADR-039.
     */
    val traceSharingEnabled = bool("traceSharingEnabled").default(true)

    /**
     * Org-configurable contact surfaced to users in CLI "contact your
     * admin" messages and the in-app budget-exceeded banner. Free text:
     * accepts an email, a URL (internal ticketing system), or any other
     * instruction. When null, callers fall back to the first ADMIN
     * member's email (`resolveSupportContact` helper).
     */
    val supportContact = text("supportContact").nullable()

    /**
     * Signup intent, editable later as the org "Primary use" setting.
     * AGENT_GOVERNANCE lands "/" on /me, LLM_OPS on the project home.
     * NULL means intent unset (every pre-fork org and any org created
     * outside onboarding) — the home resolver keeps legacy behavior.
     * ADR-038.
     */
    val primaryIntent = pgEnum("primaryIntent", "\"OrganizationIntent\"", OrganizationIntent.entries, OrganizationIntent::dbName).nullable()
    val promoCode = text("promoCode").nullable()
    val stripeCustomerId = text("stripeCustomerId").nullable().uniqueIndex("\"Organization_stripeCustomerId_key\"")
    val currency = pgEnum("currency", "\"Currency\"", Currency.entries, Currency::dbName).default(Currency.EUR)
    val pricingModel = pgEnum("pricingModel", "\"PricingModel\"", PricingModel.entries, PricingModel::dbName).default(PricingModel.TIERED)
    val license = text("license").nullable()
    val licenseExpiresAt = pgTimestamp("licenseExpiresAt", 3).nullable()
    val licenseLastValidatedAt = pgTimestamp("licenseLastValidatedAt", 3).nullable()

    /**
     * A customer that runs LangWatch self-hosted. The hosted usage, budgets and
     * invoices of its connected installs attach to this organization (ADR-139).
     */
    val selfHostedCustomer = bool("selfHostedCustomer").default(false)

    /**
     * LangWatch-hosted services this organization has switched OFF, named the
     * way its license names them ("instant_evals"). A service the license is
     * entitled to is on unless it is listed here, so a customer who bought
     * hosted judging gets it without finding a settings page first, and an
     * administrator who switched one off keeps it off across license changes
     * (ADR-139, section 2). A license that names no service reaches nothing,
     * whatever this column holds.
     */
    val connectServicesDisabled = array("connectServicesDisabled", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /**
     * When a member of this organization switched Instant Evals on from the
     * search bar, which is the organization's consent to have judged text sent
     * to the judge's provider. Null until then: a self-serve organization is
     * never opted in on its behalf, and an enterprise one is switched on by an
     * operator through the release flag after asking, never here.
     */
    val instantEvalsEnabledAt = pgTimestamp("instantEvalsEnabledAt", 3).nullable()

    /**
     * The member who switched them on, so the audit trail names a person and
     * not just a moment.
     */
    val instantEvalsEnabledByUserId = text("instantEvalsEnabledByUserId").nullable()

    /** When the last license sync succeeded. */
    val connectLastSyncAt = pgTimestamp("connectLastSyncAt", 3).nullable()

    /**
     * The code of the failure the last sync ended on, cleared on the next
     * success. Settings, Connect renders the copy for it.
     */
    val connectLastSyncError = text("connectLastSyncError").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"Organization_pkey\"")

    init {
        index("\"Organization_ssoProvider_idx\"", false, ssoProvider)
    }
}
