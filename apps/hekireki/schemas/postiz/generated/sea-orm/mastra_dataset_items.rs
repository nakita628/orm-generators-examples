use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_dataset_items")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "datasetId", column_type = "Text")]
    pub dataset_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "datasetVersion")]
    pub dataset_version: i32,
    #[sea_orm(column_name = "validTo")]
    pub valid_to: Option<i32>,
    #[sea_orm(column_name = "isDeleted")]
    pub is_deleted: bool,
    #[sea_orm(column_type = "JsonBinary")]
    pub input: Json,
    #[sea_orm(column_name = "groundTruth", column_type = "JsonBinary")]
    pub ground_truth: Option<Json>,
    #[sea_orm(column_name = "requestContext", column_type = "JsonBinary")]
    pub request_context: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub source: Option<Json>,
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
    #[sea_orm(column_name = "expectedTrajectory", column_type = "JsonBinary")]
    pub expected_trajectory: Option<Json>,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_name = "toolMocks", column_type = "JsonBinary")]
    pub tool_mocks: Option<Json>,
    #[sea_orm(column_name = "unmockedToolPolicy", column_type = "Text")]
    pub unmocked_tool_policy: Option<String>,
    #[sea_orm(column_name = "scorerIds", column_type = "JsonBinary")]
    pub scorer_ids: Option<Json>,
    #[sea_orm(column_name = "externalId", column_type = "Text")]
    pub external_id: Option<String>,
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
