use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "StatusChange")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i64,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_name = "actorId", column_type = "Text")]
    pub actor_id: String,
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(column_name = "tableName", column_type = "Text")]
    pub table_name: String,
    /// @zod.custom(z.coerce.date())
    #[sea_orm(
        column_type = "custom(\"timestamptz(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub timestamp: DateTimeWithTimeZone,
    #[sea_orm(column_type = "Text")]
    pub status: String,
    #[sea_orm(column_name = "startedAt", column_type = "custom(\"timestamp(3)\")")]
    pub started_at: DateTime,
    #[sea_orm(default_value = 0)]
    pub counts: i32,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_name = "queueSize", default_value = 0)]
    pub queue_size: i32,
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
        if insert && self.timestamp.is_not_set() {
            self.timestamp = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
