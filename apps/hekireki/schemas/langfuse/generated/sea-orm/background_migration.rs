use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "background_migrations")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub script: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub args: Json,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub state: Json,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub finished_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub failed_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub failed_reason: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub worker_id: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub locked_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
