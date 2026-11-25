use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Scenario")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub situation: String,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub criteria: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub labels: Vec<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub parameters: Option<Json>,
    #[sea_orm(column_name = "simulatorModel", column_type = "Text")]
    pub simulator_model: Option<String>,
    #[sea_orm(column_name = "judgeModel", column_type = "Text")]
    pub judge_model: Option<String>,
    #[sea_orm(column_name = "traceQuietPeriodMs")]
    pub trace_quiet_period_ms: Option<i32>,
    #[sea_orm(column_name = "traceWaitTimeoutMs")]
    pub trace_wait_timeout_ms: Option<i32>,
    #[sea_orm(column_name = "maxTurns")]
    pub max_turns: Option<i32>,
    #[sea_orm(column_name = "minTurns")]
    pub min_turns: Option<i32>,
    #[sea_orm(column_type = "JsonBinary")]
    pub fields: Option<Json>,
    #[sea_orm(column_name = "callerVoice", column_type = "JsonBinary")]
    pub caller_voice: Option<Json>,
    #[sea_orm(column_name = "testSuiteId", column_type = "Text")]
    pub test_suite_id: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "standard")]
    pub kind: String,
    #[sea_orm(column_name = "redTeam", column_type = "JsonBinary")]
    pub red_team: Option<Json>,
    #[sea_orm(column_name = "sourceTemplateId", column_type = "Text")]
    pub source_template_id: Option<String>,
    #[sea_orm(default_value = 1)]
    pub version: i32,
    #[sea_orm(column_name = "lastUpdatedById", column_type = "Text")]
    pub last_updated_by_id: Option<String>,
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
        fk_name = "Scenario_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::simulation_suite::Entity",
        from = "Column::TestSuiteId",
        to = "super::simulation_suite::Column::Id",
        on_delete = "SetNull",
        fk_name = "Scenario_testSuiteId_fkey"
    )]
    TestSuite,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::LastUpdatedById",
        to = "super::user::Column::Id",
        fk_name = "Scenario_lastUpdatedById_fkey"
    )]
    LastUpdatedBy,
    #[sea_orm(has_many = "super::scenario_version::Entity")]
    Versions,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::simulation_suite::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TestSuite.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LastUpdatedBy.def()
    }
}

impl Related<super::scenario_version::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Versions.def()
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
