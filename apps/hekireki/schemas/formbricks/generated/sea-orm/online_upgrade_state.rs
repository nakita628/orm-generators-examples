use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OnlineUpgradeState")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub release: String,
    #[sea_orm(column_name = "targetImage", column_type = "Text")]
    pub target_image: String,
    #[sea_orm(column_type = "Text")]
    pub phase: String,
    #[sea_orm(
        column_name = "completedPhases",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub completed_phases: Vec<String>,
    #[sea_orm(column_name = "surveyCursor", column_type = "Text")]
    pub survey_cursor: Option<String>,
    #[sea_orm(column_name = "surveyCount", default_value = "0")]
    pub survey_count: i64,
    #[sea_orm(
        column_name = "baselineCompletedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub baseline_completed_at: Option<DateTime>,
    #[sea_orm(
        column_name = "forwardOnlyAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub forward_only_at: Option<DateTime>,
    #[sea_orm(
        column_name = "queuePausedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub queue_paused_at: Option<DateTime>,
    #[sea_orm(
        column_name = "queuePauseDeadlineAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub queue_pause_deadline_at: Option<DateTime>,
    #[sea_orm(
        column_name = "queueResumedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub queue_resumed_at: Option<DateTime>,
    #[sea_orm(column_name = "activatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub activated_at: Option<DateTime>,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
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
