use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Durable state for a long-running event-sourced process.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ProcessManagerInstance")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "processName", column_type = "Text")]
    pub process_name: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "processKey", column_type = "Text")]
    pub process_key: String,
    #[sea_orm(column_name = "tenantId", column_type = "Text")]
    pub tenant_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub state: Json,
    pub revision: i32,
    #[sea_orm(column_name = "nextWakeAt", column_type = "custom(\"timestamptz(3)\")")]
    pub next_wake_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamptz(3)\")")]
    pub updated_at: DateTimeWithTimeZone,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
