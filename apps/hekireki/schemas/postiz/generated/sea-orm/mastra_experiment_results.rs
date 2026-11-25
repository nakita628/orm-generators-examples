use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_experiment_results")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "experimentId", column_type = "Text")]
    pub experiment_id: String,
    #[sea_orm(column_name = "itemId", column_type = "Text")]
    pub item_id: String,
    #[sea_orm(column_name = "itemDatasetVersion")]
    pub item_dataset_version: Option<i32>,
    #[sea_orm(column_type = "JsonBinary")]
    pub input: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub output: Option<Json>,
    #[sea_orm(column_name = "groundTruth", column_type = "JsonBinary")]
    pub ground_truth: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub error: Option<Json>,
    #[sea_orm(column_name = "startedAt")]
    pub started_at: DateTime,
    #[sea_orm(column_name = "completedAt")]
    pub completed_at: DateTime,
    #[sea_orm(column_name = "retryCount")]
    pub retry_count: i32,
    #[sea_orm(column_name = "traceId", column_type = "Text")]
    pub trace_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub status: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub tags: Option<Json>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "startedAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub started_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(
        column_name = "completedAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub completed_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(
        column_name = "createdAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_type = "Text")]
    pub comment: Option<String>,
    #[sea_orm(column_name = "toolMockReport", column_type = "JsonBinary")]
    pub tool_mock_report: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: Option<String>,
    pub attempt: Option<i32>,
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
        if insert && self.started_at_z.is_not_set() {
            self.started_at_z = Set(Some(now.fixed_offset()));
        }
        if insert && self.completed_at_z.is_not_set() {
            self.completed_at_z = Set(Some(now.fixed_offset()));
        }
        if insert && self.created_at_z.is_not_set() {
            self.created_at_z = Set(Some(now.fixed_offset()));
        }
        Ok(self)
    }
}
