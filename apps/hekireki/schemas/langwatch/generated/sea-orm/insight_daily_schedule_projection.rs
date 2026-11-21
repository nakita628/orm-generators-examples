use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// One person's daily insights run on one board in a project, folded from its
/// `insight_daily_schedule` aggregate (modules/insight). A read model: rebuildable by
/// replaying the event log. Epochs in milliseconds.
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "InsightDailyScheduleProjection")]
pub struct Model {
    /// The schedule id, which is also its aggregate id: one per project, person and board.
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "projectId")]
    pub project_id: String,
    #[sea_orm(column_name = "userId")]
    pub user_id: String,
    /// dashboard | template. A pointer only, never a relation: a run checks the board is there.
    #[sea_orm(column_name = "boardKind")]
    pub board_kind: String,
    #[sea_orm(column_name = "boardId")]
    pub board_id: String,
    #[sea_orm(column_name = "boardName")]
    pub board_name: String,
    /// on | off | undecided.
    #[sea_orm(default_value = "undecided")]
    pub state: String,
    /// board_deleted | skipped_days | member_removed | project_archived. Why the system turned
    /// it off. NULL while it is on, and when the person turned it off.
    #[sea_orm(column_name = "offReason")]
    pub off_reason: Option<String>,
    pub hour: Option<i32>,
    pub timezone: Option<String>,
    #[sea_orm(column_name = "maxInsights")]
    pub max_insights: Option<i32>,
    /// The last run that recorded an outcome. NULL until one did.
    #[sea_orm(column_name = "lastRunId")]
    pub last_run_id: Option<String>,
    #[sea_orm(column_name = "lastRunAt")]
    pub last_run_at: Option<f64>,
    /// filed | nothing | failed | skipped
    #[sea_orm(column_name = "lastRunOutcome")]
    pub last_run_outcome: Option<String>,
    #[sea_orm(column_name = "lastRunReason")]
    pub last_run_reason: Option<String>,
    #[sea_orm(column_name = "lastRunFiled")]
    pub last_run_filed: Option<i32>,
    /// How many insights the run found still true and renewed.
    #[sea_orm(column_name = "lastRunRenewed")]
    pub last_run_renewed: Option<i32>,
    #[sea_orm(column_name = "lastRunConversationId")]
    pub last_run_conversation_id: Option<String>,
    /// The scheduled days in a row whose run ended skipped, and the slot of the newest counted.
    #[sea_orm(column_name = "consecutiveSkips", default_value = 0)]
    pub consecutive_skips: i32,
    #[sea_orm(column_name = "skipStreakSlot")]
    pub skip_streak_slot: Option<f64>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: f64,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: f64,
    #[sea_orm(column_name = "occurredAt")]
    pub occurred_at: f64,
    #[sea_orm(column_name = "acceptedAt")]
    pub accepted_at: f64,
    #[sea_orm(column_name = "lastEventId")]
    pub last_event_id: String,
    #[sea_orm(column_name = "projectionVersion")]
    pub projection_version: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
