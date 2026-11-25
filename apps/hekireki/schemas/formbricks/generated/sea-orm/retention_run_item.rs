use super::retention_run_item_action::RetentionRunItemAction;
use super::retention_skip_reason::RetentionSkipReason;
use super::retention_target_type::RetentionTargetType;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// What a run did, one row per target: the History CSV. Ids, counts, names and the recipient's email,
/// never response content. `targetId` is not a foreign key on purpose: the target is often gone, and
/// `targetName` is a snapshot taken when the row is written so the CSV still says what it was.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RetentionRunItem")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "runId", column_type = "Text")]
    pub run_id: String,
    #[sea_orm(column_name = "targetType")]
    pub target_type: RetentionTargetType,
    #[sea_orm(column_name = "targetId", column_type = "Text")]
    pub target_id: String,
    #[sea_orm(column_name = "targetName", column_type = "Text")]
    pub target_name: Option<String>,
    pub action: RetentionRunItemAction,
    /// 1, except for grouped response deletions, where it is the number of responses deleted.
    #[sea_orm(default_value = 1)]
    pub count: i32,
    /// Who was emailed, for `notified` rows.
    #[sea_orm(column_type = "Text")]
    pub recipient: Option<String>,
    #[sea_orm(column_name = "skipReason")]
    pub skip_reason: Option<RetentionSkipReason>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::retention_run::Entity",
        from = "Column::RunId",
        to = "super::retention_run::Column::Id",
        on_delete = "Cascade",
        fk_name = "RetentionRunItem_runId_fkey"
    )]
    Run,
}

impl Related<super::retention_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Run.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
