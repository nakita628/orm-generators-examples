use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_schedule_triggers")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub schedule_id: String,
    #[sea_orm(column_type = "Text")]
    pub run_id: Option<String>,
    pub scheduled_fire_at: i64,
    pub actual_fire_at: i64,
    #[sea_orm(column_type = "Text")]
    pub outcome: String,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub trigger_kind: String,
    #[sea_orm(column_type = "Text")]
    pub parent_trigger_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
