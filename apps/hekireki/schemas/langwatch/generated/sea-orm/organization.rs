use super::currency::Currency;
use super::organization_intent::OrganizationIntent;
use super::organization_user_role::OrganizationUserRole;
use super::pricing_model::PricingModel;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Organization")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    /// `customer` for every organization a person founds or an operator provisions; `system` for
    /// the install's one hidden organization the Ops checkup canaries run in, never listed to a person.
    /// Spec: modules/ops/specs/checkup-canary-project.feature.
    #[sea_orm(column_type = "Text", default_value = "customer")]
    pub kind: String,
    #[sea_orm(column_name = "phoneNumber", column_type = "Text")]
    pub phone_number: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub slug: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "usageSpendingMaxLimit")]
    pub usage_spending_max_limit: Option<i32>,
    /// The largest file a dataset image or file cell accepts for this
    /// organization, in mebibytes. Null means the default. Set by an operator
    /// from the back office only; every other dataset size limit scales with it.
    #[sea_orm(column_name = "datasetAttachmentMaxMb")]
    pub dataset_attachment_max_mb: Option<i32>,
    /// The label of this organization's `CLICKHOUSE_URL__<label>__<id>` route, recorded by the
    /// organization module the first time it is seen. Once set, a missing route refuses the
    /// organization instead of serving it from the shared ClickHouse. Null: never private.
    /// Spec: specs/private-dataplane/clickhouse-routing.feature.
    #[sea_orm(column_name = "privateClickHouseLabel", column_type = "Text")]
    pub private_click_house_label: Option<String>,
    /// The same for `DATAPLANE_S3__<label>__<id>`. Spec: specs/private-dataplane/s3-routing.feature.
    #[sea_orm(column_name = "privateS3Label", column_type = "Text")]
    pub private_s3_label: Option<String>,
    /// When either label was first recorded.
    #[sea_orm(
        column_name = "privateDataplaneSeenAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub private_dataplane_seen_at: Option<DateTime>,
    /// Phase 8 - admin-configurable max lifetime for CLI device-flow
    /// sessions, in days. 0 = unbounded (default, mirrors GitHub CLI /
    /// gh-style flows). Companies that need finite session lifetime
    /// (eg. SOC 2 / ISO 27001 customers with mandated session
    /// expiration) can set a positive value. The /api/auth/cli/refresh
    /// endpoint rejects refresh attempts where now - issued_at exceeds
    /// this value, regardless of the underlying refresh-token TTL.
    /// Spec: specs/ai-governance/sessions/session-ttl.feature.
    #[sea_orm(column_name = "maxSessionDurationDays", default_value = 0)]
    pub max_session_duration_days: i32,
    /// Every member of this organization can prove a second factor (D06). A
    /// MEMBERSHIP CONDITION, not a session policy: it is asked when a member
    /// reaches this organization's data, it holds the ones who cannot prove
    /// one at an enrollment gate for THIS organization alone, and it ends no
    /// session, ever. Turning it on signs nobody out.
    /// Spec: specs/identity/mfa-and-session-shape.feature.
    #[sea_orm(column_name = "mfaRequired", default_value = false)]
    pub mfa_required: bool,
    /// Whether members may connect the documentation playground to their own access with a
    /// short-lived docs session key (ADR-185, amendment of 2026-11-13). On by default; turned off,
    /// no key is minted and live ones stop authenticating. The pasted-key path is untouched.
    /// Spec: modules/api-key/specs/docs-session-keys.feature.
    #[sea_orm(column_name = "docsPlaygroundSessionsEnabled", default_value = true)]
    pub docs_playground_sessions_enabled: bool,
    /// When the docs playground was last turned off, written in the same update as the switch. A
    /// docs session key minted at or before it never authenticates again, nor any key minted under
    /// one, so turning it back on revives none of them (ADR-185, amendment of 2026-11-14).
    #[sea_orm(
        column_name = "docsPlaygroundSessionsDisabledAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub docs_playground_sessions_disabled_at: Option<DateTime>,
    /// How many consecutive failed sign-ins lock an account (GAC-09).
    /// 0 = never lock, which is the default and what every organization that
    /// has not asked for this keeps. Counts EVERY credential together - a
    /// password, a one-time code, a backup code, a passkey - because five per
    /// factor hands an attacker one budget per factor.
    /// Spec: specs/identity/org-account-lockout.feature.
    #[sea_orm(column_name = "lockoutAfterFailedAttempts", default_value = 0)]
    pub lockout_after_failed_attempts: i32,
    /// How long a lock lasts, in minutes. Ignored while the threshold is 0.
    #[sea_orm(column_name = "lockoutMinutes", default_value = 30)]
    pub lockout_minutes: i32,
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
    #[sea_orm(column_name = "sessionIdleTimeoutMinutes", default_value = 1440)]
    pub session_idle_timeout_minutes: i32,
    /// The ceiling on a browser session from the moment it was minted, in
    /// minutes, which no amount of activity extends. 0 = no ceiling, the
    /// default. Offered beside the idle timeout because a session kept alive
    /// forever by a polling page is not really bounded.
    #[sea_orm(column_name = "sessionMaxLifetimeMinutes", default_value = 0)]
    pub session_max_lifetime_minutes: i32,
    #[sea_orm(column_name = "signupData", column_type = "JsonBinary")]
    pub signup_data: Option<Json>,
    #[sea_orm(column_name = "signedDPA", default_value = false)]
    pub signed_dpa: bool,
    #[sea_orm(column_name = "elasticsearchNodeUrl", column_type = "Text")]
    pub elasticsearch_node_url: Option<String>,
    #[sea_orm(column_name = "elasticsearchApiKey", column_type = "Text")]
    pub elasticsearch_api_key: Option<String>,
    #[sea_orm(column_name = "useCustomElasticsearch", default_value = false)]
    pub use_custom_elasticsearch: bool,
    #[sea_orm(column_name = "s3Endpoint", column_type = "Text")]
    pub s3_endpoint: Option<String>,
    #[sea_orm(column_name = "s3AccessKeyId", column_type = "Text")]
    pub s3_access_key_id: Option<String>,
    #[sea_orm(column_name = "s3SecretAccessKey", column_type = "Text")]
    pub s3_secret_access_key: Option<String>,
    #[sea_orm(column_name = "s3Bucket", column_type = "Text")]
    pub s3_bucket: Option<String>,
    #[sea_orm(column_name = "useCustomS3", default_value = false)]
    pub use_custom_s3: bool,
    #[sea_orm(
        column_name = "sentPlanLimitAlert",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub sent_plan_limit_alert: Option<DateTime>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "ssoDomain", column_type = "Text")]
    pub sso_domain: Option<String>,
    #[sea_orm(column_name = "ssoProvider", column_type = "Text")]
    pub sso_provider: Option<String>,
    /// How colleagues on a matching domain get in (D12, ADR-117):
    /// `off` (invisible to them entirely) | `request` (they ask, an admin
    /// approves) | `auto` (a verified colleague walks straight in).
    ///
    /// `request` is the default for a self-serve organization; `auto` is never
    /// a default and never inferred - an administrator turns it on, names the
    /// domain in `joinDomains` while doing it, and the licence gate holds that
    /// value while letting `request` through.
    #[sea_orm(
        column_name = "domainJoin",
        column_type = "Text",
        default_value = "request"
    )]
    pub domain_join: String,
    /// The organisation role handed to a person admitted without an invitation:
    /// a self-service domain join (`domainJoin`) or an SSO-admitted login.
    /// `MEMBER` (default, a Full seat) or `DEVELOPER` (ADR-171). Set on the
    /// organisation settings page only; the join flow never asks, and an
    /// invitation always names its own role.
    #[sea_orm(column_name = "joinerRole", default_value = "MEMBER")]
    pub joiner_role: OrganizationUserRole,
    /// The domains an administrator NAMED when turning automatic joining on.
    /// Empty means automatic joining admits nobody whatever `domainJoin` says -
    /// which is what keeps a setting left on from admitting a domain nobody
    /// ever agreed to.
    #[sea_orm(
        column_name = "joinDomains",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub join_domains: Vec<String>,
    #[sea_orm(column_name = "presenceEnabled", default_value = true)]
    pub presence_enabled: bool,
    /// Org-level trace-sharing kill switch. Off = trace share links cannot be
    /// created in ANY project of this org, and existing links stop resolving —
    /// the org toggle is the global switch, the per-project
    /// `Project.traceSharingEnabled` scopes the kill to a single project (mirrors
    /// `presenceEnabled`). Effective sharing = org AND project. ADR-039.
    #[sea_orm(column_name = "traceSharingEnabled", default_value = true)]
    pub trace_sharing_enabled: bool,
    /// Org-configurable contact surfaced to users in CLI "contact your
    /// admin" messages and the in-app budget-exceeded banner. Free text:
    /// accepts an email, a URL (internal ticketing system), or any other
    /// instruction. When null, callers fall back to the first ADMIN
    /// member's email (`resolveSupportContact` helper).
    #[sea_orm(column_name = "supportContact", column_type = "Text")]
    pub support_contact: Option<String>,
    /// Signup intent, editable later as the org "Primary use" setting.
    /// AGENT_GOVERNANCE lands "/" on /me, LLM_OPS on the project home.
    /// NULL means intent unset (every pre-fork org and any org created
    /// outside onboarding) — the home resolver keeps legacy behavior.
    /// ADR-038.
    #[sea_orm(column_name = "primaryIntent")]
    pub primary_intent: Option<OrganizationIntent>,
    #[sea_orm(column_name = "promoCode", column_type = "Text")]
    pub promo_code: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "stripeCustomerId", column_type = "Text")]
    pub stripe_customer_id: Option<String>,
    #[sea_orm(default_value = "EUR")]
    pub currency: Currency,
    #[sea_orm(column_name = "pricingModel", default_value = "TIERED")]
    pub pricing_model: PricingModel,
    #[sea_orm(column_type = "Text")]
    pub license: Option<String>,
    #[sea_orm(
        column_name = "licenseExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub license_expires_at: Option<DateTime>,
    #[sea_orm(
        column_name = "licenseLastValidatedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub license_last_validated_at: Option<DateTime>,
    /// A customer that runs LangWatch self-hosted. The hosted usage, budgets and
    /// invoices of its connected installs attach to this organization (ADR-139).
    #[sea_orm(column_name = "selfHostedCustomer", default_value = false)]
    pub self_hosted_customer: bool,
    /// LangWatch-hosted services this organization has switched OFF, named the
    /// way its license names them ("instant_evals"). A service the license is
    /// entitled to is on unless it is listed here, so a customer who bought
    /// hosted judging gets it without finding a settings page first, and an
    /// administrator who switched one off keeps it off across license changes
    /// (ADR-139, section 2). A license that names no service reaches nothing,
    /// whatever this column holds.
    #[sea_orm(
        column_name = "connectServicesDisabled",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub connect_services_disabled: Vec<String>,
    /// When a member of this organization switched Instant Evals on from the
    /// search bar, which is the organization's consent to have judged text sent
    /// to the judge's provider. Null until then: a self-serve organization is
    /// never opted in on its behalf, and an enterprise one is switched on by an
    /// operator through the release flag after asking, never here.
    #[sea_orm(
        column_name = "instantEvalsEnabledAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub instant_evals_enabled_at: Option<DateTime>,
    /// The member who switched them on, so the audit trail names a person and
    /// not just a moment.
    #[sea_orm(column_name = "instantEvalsEnabledByUserId", column_type = "Text")]
    pub instant_evals_enabled_by_user_id: Option<String>,
    /// When the last license sync succeeded.
    #[sea_orm(
        column_name = "connectLastSyncAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub connect_last_sync_at: Option<DateTime>,
    /// The code of the failure the last sync ended on, cleared on the next
    /// success. Settings, Connect renders the copy for it.
    #[sea_orm(column_name = "connectLastSyncError", column_type = "Text")]
    pub connect_last_sync_error: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::organization_user::Entity")]
    Members,
    #[sea_orm(has_many = "super::team::Entity")]
    Teams,
    #[sea_orm(has_many = "super::webhook_endpoint::Entity")]
    WebhookEndpoints,
    #[sea_orm(has_many = "super::organization_invite::Entity")]
    OrganizationInvite,
    #[sea_orm(has_many = "super::custom_role::Entity")]
    CustomRoles,
    #[sea_orm(has_many = "super::notification::Entity")]
    Notifications,
    #[sea_orm(has_many = "super::subscription::Entity")]
    Subscriptions,
    #[sea_orm(has_many = "super::scim_token::Entity")]
    ScimTokens,
    #[sea_orm(has_many = "super::scim_request_log::Entity")]
    ScimRequestLog,
    #[sea_orm(has_many = "super::prompt_tag::Entity")]
    PromptTags,
    #[sea_orm(has_many = "super::group::Entity")]
    Groups,
    #[sea_orm(has_many = "super::role_binding::Entity")]
    RoleBindings,
    #[sea_orm(has_many = "super::api_key::Entity")]
    ApiKeys,
    #[sea_orm(has_many = "super::issued_license::Entity")]
    IssuedLicenses,
    #[sea_orm(has_many = "super::gateway_change_event::Entity")]
    GatewayChangeEvents,
    #[sea_orm(has_many = "super::routing_policy::Entity")]
    RoutingPolicies,
    #[sea_orm(has_many = "super::ingestion_source::Entity")]
    IngestionSources,
    #[sea_orm(has_many = "super::anomaly_rule::Entity")]
    AnomalyRules,
    #[sea_orm(has_many = "super::anomaly_alert::Entity")]
    AnomalyAlerts,
    #[sea_orm(has_many = "super::ai_tool_entry::Entity")]
    AiToolEntries,
    #[sea_orm(has_many = "super::ingestion_template::Entity")]
    IngestionTemplates,
    #[sea_orm(has_many = "super::github_installation::Entity")]
    GithubInstallations,
    #[sea_orm(has_many = "super::github_pull_request::Entity")]
    GithubPullRequests,
    #[sea_orm(has_many = "super::github_branch_pull_request_check::Entity")]
    GithubBranchPrChecks,
    #[sea_orm(has_one = "super::connected_billing_account::Entity")]
    ConnectedBilling,
}

impl Related<super::organization_user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Members.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Teams.def()
    }
}

impl Related<super::webhook_endpoint::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WebhookEndpoints.def()
    }
}

impl Related<super::organization_invite::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OrganizationInvite.def()
    }
}

impl Related<super::custom_role::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomRoles.def()
    }
}

impl Related<super::notification::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Notifications.def()
    }
}

impl Related<super::subscription::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Subscriptions.def()
    }
}

impl Related<super::scim_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScimTokens.def()
    }
}

impl Related<super::scim_request_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScimRequestLog.def()
    }
}

impl Related<super::prompt_tag::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PromptTags.def()
    }
}

impl Related<super::group::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Groups.def()
    }
}

impl Related<super::role_binding::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RoleBindings.def()
    }
}

impl Related<super::api_key::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ApiKeys.def()
    }
}

impl Related<super::issued_license::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::IssuedLicenses.def()
    }
}

impl Related<super::gateway_change_event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::GatewayChangeEvents.def()
    }
}

impl Related<super::routing_policy::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RoutingPolicies.def()
    }
}

impl Related<super::ingestion_source::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::IngestionSources.def()
    }
}

impl Related<super::anomaly_rule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AnomalyRules.def()
    }
}

impl Related<super::anomaly_alert::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AnomalyAlerts.def()
    }
}

impl Related<super::ai_tool_entry::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AiToolEntries.def()
    }
}

impl Related<super::ingestion_template::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::IngestionTemplates.def()
    }
}

impl Related<super::github_installation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::GithubInstallations.def()
    }
}

impl Related<super::github_pull_request::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::GithubPullRequests.def()
    }
}

impl Related<super::github_branch_pull_request_check::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::GithubBranchPrChecks.def()
    }
}

impl Related<super::connected_billing_account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ConnectedBilling.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
