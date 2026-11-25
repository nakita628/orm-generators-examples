use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_LensTraceSignal")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub trace_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text", default_value = "")]
    pub trace_ref: String,
    #[sea_orm(column_type = "Text")]
    pub config_key: String,
    pub span_count: i32,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub claimed_until: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub classified_at: Option<DateTime>,
    #[sea_orm(column_type = "JsonBinary")]
    pub data: Json,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
