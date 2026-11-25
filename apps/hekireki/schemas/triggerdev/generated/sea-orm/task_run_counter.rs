use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// deprecated, we hadn't included the project id in the unique constraint
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TaskRunCounter")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "taskIdentifier", column_type = "Text")]
    pub task_identifier: String,
    #[sea_orm(column_name = "lastNumber", default_value = 0)]
    pub last_number: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
