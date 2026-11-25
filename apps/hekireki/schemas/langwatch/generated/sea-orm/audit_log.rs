use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AuditLog")]
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
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub action: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub args: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(column_name = "ipAddress", column_type = "Text")]
    pub ip_address: Option<String>,
    #[sea_orm(column_name = "userAgent", column_type = "Text")]
    pub user_agent: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    /// WHO REALLY DID IT, when that is not `userId`.
    ///
    /// Under an impersonation the session's `user.id` is the SUBJECT — the
    /// customer whose access an operator is borrowing — so a row naming only
    /// `userId` filed the operator's act against the customer. An operator
    /// granting a permanent way into somebody's organization appeared, in the
    /// only durable record of it, to be that somebody.
    ///
    /// Null on an ordinary request, where the two would be the same person and
    /// saying so twice adds nothing. ADR-118 / D06: every decision names both.
    #[sea_orm(column_name = "actorUserId", column_type = "Text")]
    pub actor_user_id: Option<String>,
    #[sea_orm(column_name = "targetKind", column_type = "Text")]
    pub target_kind: Option<String>,
    #[sea_orm(column_name = "targetId", column_type = "Text")]
    pub target_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub before: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub after: Option<Json>,
    /// A producer's outbox key: a redelivered audit intent finds its row here (Alex, audit R2).
    #[sea_orm(unique)]
    #[sea_orm(column_name = "idempotencyKey", column_type = "Text")]
    pub idempotency_key: Option<String>,
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
        Ok(self)
    }
}
