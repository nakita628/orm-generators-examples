use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Notification")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: Uuid,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_name = "actorId", column_type = "Text")]
    pub actor_id: String,
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(column_name = "tableName", column_type = "Text")]
    pub table_name: String,
    #[sea_orm(column_name = "channelId", column_type = "Text")]
    pub channel_id: String,
    #[sea_orm(column_name = "statusChangeId")]
    pub status_change_id: i64,
    /// @zod.custom(z.coerce.date())
    #[sea_orm(
        column_type = "custom(\"timestamptz(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub timestamp: DateTimeWithTimeZone,
    #[sea_orm(column_type = "Text")]
    pub status: String,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4()),
            ..ActiveModelTrait::default()
        }
    }

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
