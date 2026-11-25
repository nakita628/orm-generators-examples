use super::prisma_date_time::PrismaDateTime;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "document_sync_executions")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(column_name = "queueId")]
    pub queue_id: i32,
    #[sea_orm(column_type = "Text", default_value = "unknown")]
    pub status: String,
    #[sea_orm(column_type = "Text")]
    pub result: Option<String>,
    #[sea_orm(column_name = "createdAt", default_expr = "Expr::current_timestamp()")]
    pub created_at: PrismaDateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::document_sync_queues::Entity",
        from = "Column::QueueId",
        to = "super::document_sync_queues::Column::Id",
        on_delete = "Cascade",
        fk_name = "document_sync_executions_queueId_fkey"
    )]
    Queue,
}

impl Related<super::document_sync_queues::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Queue.def()
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
            self.created_at = Set(now.into());
        }
        Ok(self)
    }
}
