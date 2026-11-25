use super::authentication_method::AuthenticationMethod;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "User")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub email: String,
    /// Address awaiting confirmation; becomes `email` only after a link sent to it is opened
    #[sea_orm(column_name = "pendingEmail", column_type = "Text")]
    pub pending_email: Option<String>,
    #[sea_orm(column_name = "authenticationMethod")]
    pub authentication_method: AuthenticationMethod,
    #[sea_orm(column_name = "authenticationProfile", column_type = "JsonBinary")]
    pub authentication_profile: Option<Json>,
    #[sea_orm(column_name = "authenticationExtraParams", column_type = "JsonBinary")]
    pub authentication_extra_params: Option<Json>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "authIdentifier", column_type = "Text")]
    pub auth_identifier: Option<String>,
    #[sea_orm(column_name = "displayName", column_type = "Text")]
    pub display_name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(column_name = "avatarUrl", column_type = "Text")]
    pub avatar_url: Option<String>,
    #[sea_orm(default_value = false)]
    pub admin: bool,
    /// Preferences for the dashboard
    #[sea_orm(column_name = "dashboardPreferences", column_type = "JsonBinary")]
    pub dashboard_preferences: Option<Json>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    /// @deprecated
    #[sea_orm(column_name = "isOnCloudWaitlist", default_value = false)]
    pub is_on_cloud_waitlist: bool,
    /// @deprecated
    #[sea_orm(column_name = "featureCloud", default_value = false)]
    pub feature_cloud: bool,
    /// @deprecated
    #[sea_orm(column_name = "isOnHostedRepoWaitlist", default_value = false)]
    pub is_on_hosted_repo_waitlist: bool,
    #[sea_orm(column_name = "marketingEmails", default_value = true)]
    pub marketing_emails: bool,
    #[sea_orm(column_name = "confirmedBasicDetails", default_value = false)]
    pub confirmed_basic_details: bool,
    #[sea_orm(column_name = "referralSource", column_type = "Text")]
    pub referral_source: Option<String>,
    #[sea_orm(column_name = "onboardingData", column_type = "JsonBinary")]
    pub onboarding_data: Option<Json>,
    #[sea_orm(column_name = "mfaEnabledAt", column_type = "custom(\"timestamp(3)\")")]
    pub mfa_enabled_at: Option<DateTime>,
    #[sea_orm(column_name = "mfaSecretReferenceId", column_type = "Text")]
    pub mfa_secret_reference_id: Option<String>,
    /// Hash of the last used code to prevent replay attacks
    #[sea_orm(column_name = "mfaLastUsedCode", column_type = "Text")]
    pub mfa_last_used_code: Option<String>,
    /// Maximum session lifetime in seconds for this user. Default = 1 year (Gregorian).
    /// May be further restricted by Organization.maxSessionDuration on any of the user's orgs.
    #[sea_orm(column_name = "sessionDuration", default_value = 31556952)]
    pub session_duration: i32,
    /// Absolute timestamp at which the user's current session must end. Written
    /// at login (and any time the effective duration is recomputed) and shortened
    /// when an admin lowers an org cap. Read on requireUser / getUser; null is
    /// treated as "no enforced deadline" (legacy sessions from before this
    /// shipped). Read-on-demand has zero per-request DB cost — the value
    /// piggybacks on the User fetch that requireUser already does.
    #[sea_orm(
        column_name = "nextSessionEnd",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub next_session_end: Option<DateTime>,
    #[sea_orm(column_name = "invitationCodeId", column_type = "Text")]
    pub invitation_code_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::secret_reference::Entity",
        from = "Column::MfaSecretReferenceId",
        to = "super::secret_reference::Column::Id",
        fk_name = "User_mfaSecretReferenceId_fkey"
    )]
    MfaSecretReference,
    #[sea_orm(
        belongs_to = "super::invitation_code::Entity",
        from = "Column::InvitationCodeId",
        to = "super::invitation_code::Column::Id",
        fk_name = "User_invitationCodeId_fkey"
    )]
    InvitationCode,
    #[sea_orm(has_many = "super::org_member::Entity")]
    OrgMemberships,
    #[sea_orm(has_many = "super::org_member_invite::Entity")]
    SentInvites,
    #[sea_orm(has_many = "super::personal_access_token::Entity")]
    PersonalAccessTokens,
    #[sea_orm(has_many = "super::api_key::Entity")]
    CreatedApiKeys,
    #[sea_orm(has_many = "super::worker_deployment::Entity")]
    Deployments,
    #[sea_orm(has_many = "super::mfa_backup_code::Entity")]
    BackupCodes,
    #[sea_orm(has_many = "super::bulk_action_group::Entity")]
    BulkActions,
    #[sea_orm(
        has_many = "super::impersonation_audit_log::Entity",
        from = "Column::Id",
        to = "super::impersonation_audit_log::Column::AdminId"
    )]
    ImpersonationsPerformed,
    #[sea_orm(
        has_many = "super::impersonation_audit_log::Entity",
        from = "Column::Id",
        to = "super::impersonation_audit_log::Column::TargetId"
    )]
    ImpersonationsReceived,
    #[sea_orm(
        has_many = "super::support_access_request::Entity",
        from = "Column::Id",
        to = "super::support_access_request::Column::RequestedById"
    )]
    SupportAccessRequestsMade,
    #[sea_orm(
        has_many = "super::support_access_request::Entity",
        from = "Column::Id",
        to = "super::support_access_request::Column::ApprovedById"
    )]
    SupportAccessRequestsApproved,
    #[sea_orm(has_many = "super::customer_query::Entity")]
    CustomerQueries,
    #[sea_orm(has_many = "super::metrics_dashboard::Entity")]
    MetricsDashboards,
    #[sea_orm(has_many = "super::platform_notification::Entity")]
    PlatformNotifications,
    #[sea_orm(has_many = "super::platform_notification_interaction::Entity")]
    PlatformNotificationInteractions,
    #[sea_orm(has_many = "super::environment_variable_ignore::Entity")]
    EnvironmentVariableIgnores,
}

impl Related<super::secret_reference::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::MfaSecretReference.def()
    }
}

impl Related<super::invitation_code::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::InvitationCode.def()
    }
}

impl Related<super::org_member::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OrgMemberships.def()
    }
}

impl Related<super::org_member_invite::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SentInvites.def()
    }
}

impl Related<super::personal_access_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PersonalAccessTokens.def()
    }
}

impl Related<super::api_key::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedApiKeys.def()
    }
}

impl Related<super::worker_deployment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Deployments.def()
    }
}

impl Related<super::mfa_backup_code::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BackupCodes.def()
    }
}

impl Related<super::bulk_action_group::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BulkActions.def()
    }
}

impl Related<super::impersonation_audit_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ImpersonationsPerformed.def()
    }
}

impl Related<super::support_access_request::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SupportAccessRequestsMade.def()
    }
}

impl Related<super::customer_query::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomerQueries.def()
    }
}

impl Related<super::metrics_dashboard::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::MetricsDashboards.def()
    }
}

impl Related<super::platform_notification::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PlatformNotifications.def()
    }
}

impl Related<super::platform_notification_interaction::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PlatformNotificationInteractions.def()
    }
}

impl Related<super::environment_variable_ignore::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EnvironmentVariableIgnores.def()
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
