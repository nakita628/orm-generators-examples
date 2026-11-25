use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_datasets")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(column_name = "inputSchema", column_type = "JsonBinary")]
    pub input_schema: Option<Json>,
    #[sea_orm(column_name = "groundTruthSchema", column_type = "JsonBinary")]
    pub ground_truth_schema: Option<Json>,
    #[sea_orm(column_name = "requestContextSchema", column_type = "JsonBinary")]
    pub request_context_schema: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub tags: Option<Json>,
    #[sea_orm(column_name = "targetType", column_type = "Text")]
    pub target_type: Option<String>,
    #[sea_orm(column_name = "targetIds", column_type = "JsonBinary")]
    pub target_ids: Option<Json>,
    #[sea_orm(column_name = "scorerIds", column_type = "JsonBinary")]
    pub scorer_ids: Option<Json>,
    pub version: i32,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(
        column_name = "createdAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(
        column_name = "updatedAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_name = "candidateKey", column_type = "Text")]
    pub candidate_key: Option<String>,
    #[sea_orm(column_name = "candidateId", column_type = "Text")]
    pub candidate_id: Option<String>,
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
        if insert && self.updated_at_z.is_not_set() {
            self.updated_at_z = Set(Some(now.fixed_offset()));
        }
        Ok(self)
    }
}
