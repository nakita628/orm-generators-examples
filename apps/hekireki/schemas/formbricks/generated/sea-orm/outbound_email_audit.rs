use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Metadata for an outbound email attempt. Content and recipient addresses are never stored here.
/// Identifiers are deliberately not foreign keys so erasure can scrub a row without retaining source data.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OutboundEmailAudit")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: DateTime,
    #[sea_orm(
        column_name = "contentExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub content_expires_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub environment: String,
    #[sea_orm(column_name = "emailType", column_type = "Text")]
    pub email_type: String,
    #[sea_orm(column_type = "Text")]
    pub trigger: String,
    #[sea_orm(column_name = "actorKind", column_type = "Text")]
    pub actor_kind: String,
    #[sea_orm(column_name = "actorUserId", column_type = "Text")]
    pub actor_user_id: Option<String>,
    #[sea_orm(column_name = "configuredByUserId", column_type = "Text")]
    pub configured_by_user_id: Option<String>,
    #[sea_orm(column_name = "actorUserHash", column_type = "Text")]
    pub actor_user_hash: Option<String>,
    #[sea_orm(column_name = "configuredByUserHash", column_type = "Text")]
    pub configured_by_user_hash: Option<String>,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_name = "surveyId", column_type = "Text")]
    pub survey_id: Option<String>,
    #[sea_orm(column_name = "responseId", column_type = "Text")]
    pub response_id: Option<String>,
    #[sea_orm(column_name = "followUpId", column_type = "Text")]
    pub follow_up_id: Option<String>,
    #[sea_orm(column_name = "workflowId", column_type = "Text")]
    pub workflow_id: Option<String>,
    #[sea_orm(column_name = "workflowRunId", column_type = "Text")]
    pub workflow_run_id: Option<String>,
    #[sea_orm(column_name = "inviteId", column_type = "Text")]
    pub invite_id: Option<String>,
    #[sea_orm(column_name = "recipientMac", column_type = "Text")]
    pub recipient_mac: Option<String>,
    #[sea_orm(column_name = "sesMessageHash", column_type = "Text")]
    pub ses_message_hash: Option<String>,
    #[sea_orm(column_name = "snapshotKey", column_type = "Text")]
    pub snapshot_key: Option<String>,
    #[sea_orm(column_name = "snapshotSha256", column_type = "Text")]
    pub snapshot_sha256: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "prepared")]
    pub status: String,
    #[sea_orm(
        column_name = "statusRecordedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub status_recorded_at: Option<DateTime>,
    #[sea_orm(column_name = "auditGap", default_value = false)]
    pub audit_gap: bool,
    #[sea_orm(
        column_name = "erasureRequestedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub erasure_requested_at: Option<DateTime>,
    #[sea_orm(column_name = "erasedAt", column_type = "custom(\"timestamp(3)\")")]
    pub erased_at: Option<DateTime>,
    #[sea_orm(
        column_name = "contentDeleteRetryAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub content_delete_retry_at: Option<DateTime>,
    #[sea_orm(column_name = "contentDeleteFailures", default_value = 0)]
    pub content_delete_failures: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

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
