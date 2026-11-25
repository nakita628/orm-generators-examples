use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Used for the TaskRun number (e.g. #1,421)
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TaskRunNumberCounter")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "taskIdentifier", column_type = "Text")]
    pub task_identifier: String,
    #[sea_orm(column_name = "environmentId", column_type = "Text")]
    pub environment_id: String,
    #[sea_orm(column_name = "lastNumber", default_value = 0)]
    pub last_number: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::runtime_environment::Entity",
        from = "Column::EnvironmentId",
        to = "super::runtime_environment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "TaskRunNumberCounter_environmentId_fkey"
    )]
    Environment,
}

impl Related<super::runtime_environment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Environment.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
