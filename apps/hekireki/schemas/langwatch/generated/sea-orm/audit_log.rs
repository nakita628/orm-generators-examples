use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AuditLog")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "projectId")]
    pub project_id: Option<String>,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: Option<String>,
    pub action: String,
    pub args: Option<Json>,
    pub error: Option<String>,
    #[sea_orm(column_name = "ipAddress")]
    pub ip_address: Option<String>,
    #[sea_orm(column_name = "userAgent")]
    pub user_agent: Option<String>,
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
    #[sea_orm(column_name = "actorUserId")]
    pub actor_user_id: Option<String>,
    #[sea_orm(column_name = "targetKind")]
    pub target_kind: Option<String>,
    #[sea_orm(column_name = "targetId")]
    pub target_id: Option<String>,
    pub before: Option<Json>,
    pub after: Option<Json>,
    /// A producer's outbox key: a redelivered audit intent finds its row here (Alex, audit R2).
    #[sea_orm(unique)]
    #[sea_orm(column_name = "idempotencyKey")]
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
