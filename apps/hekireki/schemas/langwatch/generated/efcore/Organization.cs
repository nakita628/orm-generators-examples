#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

public partial class Organization
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    /// <summary>
    /// `customer` for every organization a person founds or an operator provisions; `system` for
    /// the install's one hidden organization the Ops checkup canaries run in, never listed to a person.
    /// Spec: modules/ops/specs/checkup-canary-project.feature.
    /// </summary>
    public string Kind { get; set; } = "customer";

    public string? PhoneNumber { get; set; }

    public string Slug { get; set; } = null!;

    public virtual ICollection<OrganizationUser> Members { get; set; } = new List<OrganizationUser>();

    public virtual ICollection<Team> Teams { get; set; } = new List<Team>();

    public virtual ICollection<WebhookEndpoint> WebhookEndpoints { get; set; } = new List<WebhookEndpoint>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<OrganizationInvite> OrganizationInvite { get; set; } = new List<OrganizationInvite>();

    public int? UsageSpendingMaxLimit { get; set; }

    /// <summary>
    /// The largest file a dataset image or file cell accepts for this
    /// organization, in mebibytes. Null means the default. Set by an operator
    /// from the back office only; every other dataset size limit scales with it.
    /// </summary>
    public int? DatasetAttachmentMaxMb { get; set; }

    /// <summary>
    /// The label of this organization's `CLICKHOUSE_URL__&lt;label&gt;__&lt;id&gt;` route, recorded by the
    /// organization module the first time it is seen. Once set, a missing route refuses the
    /// organization instead of serving it from the shared ClickHouse. Null: never private.
    /// Spec: specs/private-dataplane/clickhouse-routing.feature.
    /// </summary>
    public string? PrivateClickHouseLabel { get; set; }

    /// <summary>
    /// The same for `DATAPLANE_S3__&lt;label&gt;__&lt;id&gt;`. Spec: specs/private-dataplane/s3-routing.feature.
    /// </summary>
    public string? PrivateS3Label { get; set; }

    /// <summary>
    /// When either label was first recorded.
    /// </summary>
    public DateTime? PrivateDataplaneSeenAt { get; set; }

    /// <summary>
    /// Phase 8 - admin-configurable max lifetime for CLI device-flow
    /// sessions, in days. 0 = unbounded (default, mirrors GitHub CLI /
    /// gh-style flows). Companies that need finite session lifetime
    /// (eg. SOC 2 / ISO 27001 customers with mandated session
    /// expiration) can set a positive value. The /api/auth/cli/refresh
    /// endpoint rejects refresh attempts where now - issued_at exceeds
    /// this value, regardless of the underlying refresh-token TTL.
    /// Spec: specs/ai-governance/sessions/session-ttl.feature.
    /// </summary>
    public int MaxSessionDurationDays { get; set; } = 0;

    /// <summary>
    /// Every member of this organization can prove a second factor (D06). A
    /// MEMBERSHIP CONDITION, not a session policy: it is asked when a member
    /// reaches this organization's data, it holds the ones who cannot prove
    /// one at an enrollment gate for THIS organization alone, and it ends no
    /// session, ever. Turning it on signs nobody out.
    /// Spec: specs/identity/mfa-and-session-shape.feature.
    /// </summary>
    public bool MfaRequired { get; set; } = false;

    /// <summary>
    /// Whether members may connect the documentation playground to their own access with a
    /// short-lived docs session key (ADR-185, amendment of 2026-11-13). On by default; turned off,
    /// no key is minted and live ones stop authenticating. The pasted-key path is untouched.
    /// Spec: modules/api-key/specs/docs-session-keys.feature.
    /// </summary>
    public bool DocsPlaygroundSessionsEnabled { get; set; } = true;

    /// <summary>
    /// When the docs playground was last turned off, written in the same update as the switch. A
    /// docs session key minted at or before it never authenticates again, nor any key minted under
    /// one, so turning it back on revives none of them (ADR-185, amendment of 2026-11-14).
    /// </summary>
    public DateTime? DocsPlaygroundSessionsDisabledAt { get; set; }

    /// <summary>
    /// How many consecutive failed sign-ins lock an account (GAC-09).
    /// 0 = never lock, which is the default and what every organization that
    /// has not asked for this keeps. Counts EVERY credential together - a
    /// password, a one-time code, a backup code, a passkey - because five per
    /// factor hands an attacker one budget per factor.
    /// Spec: specs/identity/org-account-lockout.feature.
    /// </summary>
    public int LockoutAfterFailedAttempts { get; set; } = 0;

    /// <summary>
    /// How long a lock lasts, in minutes. Ignored while the threshold is 0.
    /// </summary>
    public int LockoutMinutes { get; set; } = 30;

    /// <summary>
    /// How long a BROWSER session may sit idle before it ends, in minutes
    /// (GAC-10). 0 = no idle timeout. Activity rolls the window, which is the
    /// "refreshed every hour" half of the control.
    ///
    /// A day for organizations created from now on; organizations that predate
    /// the column keep 0 and are unchanged, because the migration adds the
    /// column at zero and moves the default afterwards.
    ///
    /// NOT `maxSessionDurationDays`, which bounds CLI and device sessions and
    /// is enforced when the CLI refreshes. Two surfaces, two lifetimes, two
    /// decisions an administrator makes separately.
    /// Spec: specs/identity/org-session-lifetime.feature.
    /// </summary>
    public int SessionIdleTimeoutMinutes { get; set; } = 1440;

    /// <summary>
    /// The ceiling on a browser session from the moment it was minted, in
    /// minutes, which no amount of activity extends. 0 = no ceiling, the
    /// default. Offered beside the idle timeout because a session kept alive
    /// forever by a polling page is not really bounded.
    /// </summary>
    public int SessionMaxLifetimeMinutes { get; set; } = 0;

    public string? SignupData { get; set; }

    public bool SignedDPA { get; set; } = false;

    public string? ElasticsearchNodeUrl { get; set; }

    public string? ElasticsearchApiKey { get; set; }

    public bool UseCustomElasticsearch { get; set; } = false;

    public string? S3Endpoint { get; set; }

    public string? S3AccessKeyId { get; set; }

    public string? S3SecretAccessKey { get; set; }

    public string? S3Bucket { get; set; }

    public bool UseCustomS3 { get; set; } = false;

    public DateTime? SentPlanLimitAlert { get; set; }

    public string? SsoDomain { get; set; }

    public string? SsoProvider { get; set; }

    /// <summary>
    /// How colleagues on a matching domain get in (D12, ADR-117):
    /// `off` (invisible to them entirely) | `request` (they ask, an admin
    /// approves) | `auto` (a verified colleague walks straight in).
    ///
    /// `request` is the default for a self-serve organization; `auto` is never
    /// a default and never inferred - an administrator turns it on, names the
    /// domain in `joinDomains` while doing it, and the licence gate holds that
    /// value while letting `request` through.
    /// </summary>
    public string DomainJoin { get; set; } = "request";

    /// <summary>
    /// The organisation role handed to a person admitted without an invitation:
    /// a self-service domain join (`domainJoin`) or an SSO-admitted login.
    /// `MEMBER` (default, a Full seat) or `DEVELOPER` (ADR-171). Set on the
    /// organisation settings page only; the join flow never asks, and an
    /// invitation always names its own role.
    /// </summary>
    public OrganizationUserRole JoinerRole { get; set; } = OrganizationUserRole.Member;

    /// <summary>
    /// The domains an administrator NAMED when turning automatic joining on.
    /// Empty means automatic joining admits nobody whatever `domainJoin` says -
    /// which is what keeps a setting left on from admitting a domain nobody
    /// ever agreed to.
    /// </summary>
    public List<string>? JoinDomains { get; set; }

    public bool PresenceEnabled { get; set; } = true;

    /// <summary>
    /// Org-level trace-sharing kill switch. Off = trace share links cannot be
    /// created in ANY project of this org, and existing links stop resolving —
    /// the org toggle is the global switch, the per-project
    /// `Project.traceSharingEnabled` scopes the kill to a single project (mirrors
    /// `presenceEnabled`). Effective sharing = org AND project. ADR-039.
    /// </summary>
    public bool TraceSharingEnabled { get; set; } = true;

    /// <summary>
    /// Org-configurable contact surfaced to users in CLI "contact your
    /// admin" messages and the in-app budget-exceeded banner. Free text:
    /// accepts an email, a URL (internal ticketing system), or any other
    /// instruction. When null, callers fall back to the first ADMIN
    /// member's email (`resolveSupportContact` helper).
    /// </summary>
    public string? SupportContact { get; set; }

    /// <summary>
    /// Signup intent, editable later as the org "Primary use" setting.
    /// AGENT_GOVERNANCE lands "/" on /me, LLM_OPS on the project home.
    /// NULL means intent unset (every pre-fork org and any org created
    /// outside onboarding) — the home resolver keeps legacy behavior.
    /// ADR-038.
    /// </summary>
    public OrganizationIntent? PrimaryIntent { get; set; }

    public string? PromoCode { get; set; }

    public string? StripeCustomerId { get; set; }

    public Currency Currency { get; set; } = Currency.Eur;

    public PricingModel PricingModel { get; set; } = PricingModel.Tiered;

    public virtual ICollection<CustomRole> CustomRoles { get; set; } = new List<CustomRole>();

    public virtual ICollection<Notification> Notifications { get; set; } = new List<Notification>();

    public virtual ICollection<Subscription> Subscriptions { get; set; } = new List<Subscription>();

    public virtual ICollection<ScimToken> ScimTokens { get; set; } = new List<ScimToken>();

    public virtual ICollection<ScimRequestLog> ScimRequestLog { get; set; } = new List<ScimRequestLog>();

    public virtual ICollection<PromptTag> PromptTags { get; set; } = new List<PromptTag>();

    public virtual ICollection<Group> Groups { get; set; } = new List<Group>();

    public virtual ICollection<RoleBinding> RoleBindings { get; set; } = new List<RoleBinding>();

    public virtual ICollection<ApiKey> ApiKeys { get; set; } = new List<ApiKey>();

    public string? License { get; set; }

    public DateTime? LicenseExpiresAt { get; set; }

    public DateTime? LicenseLastValidatedAt { get; set; }

    /// <summary>
    /// A customer that runs LangWatch self-hosted. The hosted usage, budgets and
    /// invoices of its connected installs attach to this organization (ADR-139).
    /// </summary>
    public bool SelfHostedCustomer { get; set; } = false;

    /// <summary>
    /// LangWatch-hosted services this organization has switched OFF, named the
    /// way its license names them ("instant_evals"). A service the license is
    /// entitled to is on unless it is listed here, so a customer who bought
    /// hosted judging gets it without finding a settings page first, and an
    /// administrator who switched one off keeps it off across license changes
    /// (ADR-139, section 2). A license that names no service reaches nothing,
    /// whatever this column holds.
    /// </summary>
    public List<string>? ConnectServicesDisabled { get; set; } = new List<string>();

    /// <summary>
    /// When a member of this organization switched Instant Evals on from the
    /// search bar, which is the organization's consent to have judged text sent
    /// to the judge's provider. Null until then: a self-serve organization is
    /// never opted in on its behalf, and an enterprise one is switched on by an
    /// operator through the release flag after asking, never here.
    /// </summary>
    public DateTime? InstantEvalsEnabledAt { get; set; }

    /// <summary>
    /// The member who switched them on, so the audit trail names a person and
    /// not just a moment.
    /// </summary>
    public string? InstantEvalsEnabledByUserId { get; set; }

    /// <summary>
    /// When the last license sync succeeded.
    /// </summary>
    public DateTime? ConnectLastSyncAt { get; set; }

    /// <summary>
    /// The code of the failure the last sync ended on, cleared on the next
    /// success. Settings, Connect renders the copy for it.
    /// </summary>
    public string? ConnectLastSyncError { get; set; }

    public virtual ICollection<IssuedLicense> IssuedLicenses { get; set; } = new List<IssuedLicense>();

    public virtual ConnectedBillingAccount? ConnectedBilling { get; set; }

    public virtual ICollection<GatewayChangeEvent> GatewayChangeEvents { get; set; } = new List<GatewayChangeEvent>();

    public virtual ICollection<RoutingPolicy> RoutingPolicies { get; set; } = new List<RoutingPolicy>();

    public virtual ICollection<IngestionSource> IngestionSources { get; set; } = new List<IngestionSource>();

    public virtual ICollection<AnomalyRule> AnomalyRules { get; set; } = new List<AnomalyRule>();

    public virtual ICollection<AnomalyAlert> AnomalyAlerts { get; set; } = new List<AnomalyAlert>();

    public virtual ICollection<AiToolEntry> AiToolEntries { get; set; } = new List<AiToolEntry>();

    public virtual ICollection<IngestionTemplate> IngestionTemplates { get; set; } = new List<IngestionTemplate>();

    public virtual ICollection<GithubInstallation> GithubInstallations { get; set; } = new List<GithubInstallation>();

    public virtual ICollection<GithubPullRequest> GithubPullRequests { get; set; } = new List<GithubPullRequest>();

    public virtual ICollection<GithubBranchPullRequestCheck> GithubBranchPrChecks { get; set; } = new List<GithubBranchPullRequestCheck>();
}
