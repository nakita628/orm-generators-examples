use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_traces")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "parentSpanId", column_type = "Text")]
    pub parent_span_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "traceId", column_type = "Text")]
    pub trace_id: String,
    #[sea_orm(column_type = "Text")]
    pub scope: String,
    pub kind: i32,
    #[sea_orm(column_type = "JsonBinary")]
    pub attributes: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub status: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub events: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub links: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub other: Option<String>,
    #[sea_orm(column_name = "startTime")]
    pub start_time: i64,
    #[sea_orm(column_name = "endTime")]
    pub end_time: i64,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "createdAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at_z: Option<DateTimeWithTimeZone>,
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
        Ok(self)
    }
}
