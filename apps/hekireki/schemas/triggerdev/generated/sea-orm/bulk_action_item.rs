use super::bulk_action_item_status::BulkActionItemStatus;
use super::bulk_action_type::BulkActionType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BulkActionItem")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// @deprecated not used in new BulkActions
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: Option<String>,
    #[sea_orm(column_name = "groupId", column_type = "Text")]
    pub group_id: String,
    pub r#type: BulkActionType,
    /// When the item is created it's pending. After we've processed the item it's completed. This does not mean the associated runs are completed.
    #[sea_orm(default_value = "PENDING")]
    pub status: BulkActionItemStatus,
    #[sea_orm(column_name = "sourceRunId", column_type = "Text")]
    pub source_run_id: String,
    #[sea_orm(column_name = "destinationRunId", column_type = "Text")]
    pub destination_run_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::task_run::Entity",
        from = "Column::SourceRunId",
        to = "super::task_run::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BulkActionItem_sourceRunId_fkey"
    )]
    SourceRun,
    #[sea_orm(
        belongs_to = "super::task_run::Entity",
        from = "Column::DestinationRunId",
        to = "super::task_run::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BulkActionItem_destinationRunId_fkey"
    )]
    DestinationRun,
}

impl Related<super::task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SourceRun.def()
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
