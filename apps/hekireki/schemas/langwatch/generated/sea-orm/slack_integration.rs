use super::slack_integration_kind::SlackIntegrationKind;
use super::slack_integration_scope_type::SlackIntegrationScopeType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One named Slack connection (ADR-093 §5a); an organization holds as many as
/// it likes, each scoped to the organization or to one project. The secret is
/// AES-256-GCM ciphertext and never returned to a client. `secretFingerprint`
/// is an HMAC of the plaintext, so one secret is one connection per scope
/// without decrypting anything. No foreign key: `scopeId` is
/// polymorphic (ADR-021), as on RetentionPolicy and CustomLLMModelCost.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SlackIntegration")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    pub kind: SlackIntegrationKind,
    #[sea_orm(column_name = "scopeType")]
    pub scope_type: SlackIntegrationScopeType,
    #[sea_orm(column_name = "scopeId", column_type = "Text")]
    pub scope_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "botTokenEncrypted", column_type = "Text")]
    pub bot_token_encrypted: Option<String>,
    #[sea_orm(column_name = "webhookUrlEncrypted", column_type = "Text")]
    pub webhook_url_encrypted: Option<String>,
    #[sea_orm(column_name = "secretFingerprint", column_type = "Text")]
    pub secret_fingerprint: String,
    #[sea_orm(column_name = "secretHint", column_type = "Text")]
    pub secret_hint: String,
    #[sea_orm(column_name = "slackTeamId", column_type = "Text")]
    pub slack_team_id: Option<String>,
    #[sea_orm(column_name = "slackTeamName", column_type = "Text")]
    pub slack_team_name: Option<String>,
    #[sea_orm(column_name = "createdById", column_type = "Text")]
    pub created_by_id: String,
    #[sea_orm(column_name = "updatedById", column_type = "Text")]
    pub updated_by_id: String,
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
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

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
