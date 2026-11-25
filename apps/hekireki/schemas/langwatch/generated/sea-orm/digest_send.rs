use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One person's digest for one ISO week, accepted by our own mailer. The id is `userId:isoWeek`.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DigestSend")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "isoWeek", column_type = "Text")]
    pub iso_week: String,
    #[sea_orm(column_name = "templateId", column_type = "Text")]
    pub template_id: String,
    #[sea_orm(column_name = "sentAt", column_type = "custom(\"timestamp(3)\")")]
    pub sent_at: DateTime,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
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
