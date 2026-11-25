use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SimulationSuite")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub slug: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(
        column_name = "scenarioIds",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub scenario_ids: Vec<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub targets: Json,
    #[sea_orm(column_name = "repeatCount", default_value = 1)]
    pub repeat_count: i32,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub labels: Vec<String>,
    #[sea_orm(column_name = "simulatorModel", column_type = "Text")]
    pub simulator_model: Option<String>,
    #[sea_orm(column_name = "judgeModel", column_type = "Text")]
    pub judge_model: Option<String>,
    #[sea_orm(column_name = "traceQuietPeriodMs")]
    pub trace_quiet_period_ms: Option<i32>,
    #[sea_orm(column_name = "traceWaitTimeoutMs")]
    pub trace_wait_timeout_ms: Option<i32>,
    #[sea_orm(column_type = "Text", default_value = "run_plan")]
    pub kind: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub scope: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub fields: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub evaluators: Option<Json>,
    #[sea_orm(column_name = "templateId", column_type = "Text")]
    pub template_id: Option<String>,
    #[sea_orm(column_name = "templateVersion")]
    pub template_version: Option<i32>,
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamp(3)\")")]
    pub archived_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        fk_name = "SimulationSuite_projectId_fkey"
    )]
    Project,
    #[sea_orm(has_many = "super::scenario::Entity")]
    TestSuiteScenarios,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::scenario::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TestSuiteScenarios.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
