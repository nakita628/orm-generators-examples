use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_schedules")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub target: Json,
    #[sea_orm(column_type = "Text")]
    pub cron: String,
    #[sea_orm(column_type = "Text")]
    pub timezone: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub status: String,
    pub next_fire_at: i64,
    pub last_fire_at: Option<i64>,
    #[sea_orm(column_type = "Text")]
    pub last_run_id: Option<String>,
    pub created_at: i64,
    pub updated_at: i64,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub owner_type: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub owner_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
