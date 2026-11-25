use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Nurturing's read model per organization: its admin and its milestone counts.
/// `seeded` marks an organization learned from the project backfill: its first_* already happened.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingOrganization")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "adminUserId", column_type = "Text")]
    pub admin_user_id: Option<String>,
    #[sea_orm(default_value = false)]
    pub seeded: bool,
    #[sea_orm(column_name = "evaluationCount", default_value = 0)]
    pub evaluation_count: i32,
    #[sea_orm(column_name = "simulationRunCount", default_value = 0)]
    pub simulation_run_count: i32,
    #[sea_orm(
        column_name = "codingFirstSessionAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub coding_first_session_at: Option<DateTime>,
    #[sea_orm(
        column_name = "codingAgentTypes",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub coding_agent_types: Vec<String>,
    #[sea_orm(column_name = "codingActiveDays", default_value = 0)]
    pub coding_active_days: i32,
    #[sea_orm(column_name = "codingLastActiveDay", column_type = "Text")]
    pub coding_last_active_day: Option<String>,
    #[sea_orm(
        column_name = "codingTerminalTypes",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub coding_terminal_types: Vec<String>,
    #[sea_orm(
        column_name = "codingRepositoryKeys",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub coding_repository_keys: Vec<String>,
    #[sea_orm(column_name = "codingFirstDays")]
    pub coding_first_days: Vec<i32>,
    #[sea_orm(column_name = "lastTraceAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_trace_at: Option<DateTime>,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, _insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
