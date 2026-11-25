use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_knowledge_records")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub node: String,
    #[sea_orm(column_type = "Text")]
    pub text: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub scope: Json,
    #[sea_orm(column_name = "scopeKey", column_type = "Text")]
    pub scope_key: String,
    #[sea_orm(column_name = "sourceThreadId", column_type = "Text")]
    pub source_thread_id: String,
    #[sea_orm(column_name = "capturedAt")]
    pub captured_at: DateTime,
    pub when: Option<DateTime>,
    #[sea_orm(column_name = "maxScope", column_type = "Text")]
    pub max_scope: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(column_name = "deletedAt")]
    pub deleted_at: Option<DateTime>,
    #[sea_orm(column_name = "deletedBy", column_type = "Text")]
    pub deleted_by: Option<String>,
    #[sea_orm(
        column_name = "capturedAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub captured_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(
        column_name = "whenZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub when_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(
        column_name = "deletedAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub deleted_at_z: Option<DateTimeWithTimeZone>,
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
        if insert && self.captured_at_z.is_not_set() {
            self.captured_at_z = Set(Some(now.fixed_offset()));
        }
        if insert && self.when_z.is_not_set() {
            self.when_z = Set(Some(now.fixed_offset()));
        }
        if insert && self.deleted_at_z.is_not_set() {
            self.deleted_at_z = Set(Some(now.fixed_offset()));
        }
        Ok(self)
    }
}
