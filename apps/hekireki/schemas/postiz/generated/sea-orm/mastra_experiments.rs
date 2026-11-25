use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_experiments")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(column_name = "datasetId", column_type = "Text")]
    pub dataset_id: Option<String>,
    #[sea_orm(column_name = "datasetVersion")]
    pub dataset_version: Option<i32>,
    #[sea_orm(column_name = "targetType", column_type = "Text")]
    pub target_type: String,
    #[sea_orm(column_name = "targetId", column_type = "Text")]
    pub target_id: String,
    #[sea_orm(column_type = "Text")]
    pub status: String,
    #[sea_orm(column_name = "totalItems")]
    pub total_items: i32,
    #[sea_orm(column_name = "succeededCount")]
    pub succeeded_count: i32,
    #[sea_orm(column_name = "failedCount")]
    pub failed_count: i32,
    #[sea_orm(column_name = "skippedCount")]
    pub skipped_count: i32,
    #[sea_orm(column_name = "startedAt")]
    pub started_at: Option<DateTime>,
    #[sea_orm(column_name = "completedAt")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(column_name = "agentVersion", column_type = "Text")]
    pub agent_version: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
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
    #[sea_orm(column_type = "JsonBinary")]
    pub provenance: Option<Json>,
    #[sea_orm(column_name = "runnerAttestation", column_type = "JsonBinary")]
    pub runner_attestation: Option<Json>,
    #[sea_orm(column_name = "experimentSetId", column_type = "Text")]
    pub experiment_set_id: Option<String>,
    #[sea_orm(column_name = "comparisonId", column_type = "Text")]
    pub comparison_id: Option<String>,
    #[sea_orm(column_name = "variantId", column_type = "Text")]
    pub variant_id: Option<String>,
    #[sea_orm(column_name = "trialIndex")]
    pub trial_index: Option<i32>,
    #[sea_orm(column_name = "scorerIds", column_type = "JsonBinary")]
    pub scorer_ids: Option<Json>,
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
        if insert && self.updated_at_z.is_not_set() {
            self.updated_at_z = Set(Some(now.fixed_offset()));
        }
        Ok(self)
    }
}
