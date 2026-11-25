use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_background_tasks")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub tool_call_id: String,
    #[sea_orm(column_type = "Text")]
    pub tool_name: String,
    #[sea_orm(column_type = "Text")]
    pub agent_id: String,
    #[sea_orm(column_type = "Text")]
    pub run_id: String,
    #[sea_orm(column_type = "Text")]
    pub thread_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub resource_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub status: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub args: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub result: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub error: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub suspend_payload: Option<Json>,
    pub retry_count: i32,
    pub max_retries: i32,
    pub timeout_ms: i32,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "startedAt")]
    pub started_at: Option<DateTime>,
    #[sea_orm(column_name = "suspendedAt")]
    pub suspended_at: Option<DateTime>,
    #[sea_orm(column_name = "completedAt")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(
        column_name = "startedAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub started_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(
        column_name = "suspendedAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub suspended_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(
        column_name = "completedAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub completed_at_z: Option<DateTimeWithTimeZone>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at_z.is_not_set() {
            self.created_at_z = Set(Some(now.fixed_offset()));
        }
        if insert && self.started_at_z.is_not_set() {
            self.started_at_z = Set(Some(now.fixed_offset()));
        }
        if insert && self.suspended_at_z.is_not_set() {
            self.suspended_at_z = Set(Some(now.fixed_offset()));
        }
        if insert && self.completed_at_z.is_not_set() {
            self.completed_at_z = Set(Some(now.fixed_offset()));
        }
        Ok(self)
    }
}
