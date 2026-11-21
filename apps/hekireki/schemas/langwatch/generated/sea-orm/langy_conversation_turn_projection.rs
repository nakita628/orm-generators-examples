use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Rebuildable per-turn operational render document.
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LangyConversationTurnProjection")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "projectId")]
    pub project_id: String,
    #[sea_orm(column_name = "conversationId")]
    pub conversation_id: String,
    #[sea_orm(column_name = "turnId")]
    pub turn_id: String,
    /// LangyConversationTurnStatus. Deliberately TEXT, not a Postgres enum: this
    /// is a PROJECTION — one writer (the fold), rebuildable from the event log,
    /// and nothing queries or indexes on it. An enum bought no selectivity here,
    /// only typo protection, and charged a migration for every new status that
    /// Postgres then could not remove. The one definition lives in
    /// `LANGY_CONVERSATION_TURN_STATUS`; the repository validates against it.
    pub status: String,
    /// LangyMessagePart[] — the user question that opened the turn.
    #[sea_orm(column_name = "questionParts")]
    pub question_parts: Json,
    /// LangyMessagePart[] — the agent's whole final answer.
    #[sea_orm(column_name = "answerParts")]
    pub answer_parts: Json,
    /// LangyTurnToolCall[] — tool-call lifecycle audit in initiation order.
    #[sea_orm(column_name = "toolCalls")]
    pub tool_calls: Json,
    /// LangyPlanItemData[] | null — null means the turn never kept a plan.
    pub plan: Option<Json>,
    pub error: Option<String>,
    #[sea_orm(column_name = "startedAt")]
    pub started_at: Option<f64>,
    #[sea_orm(column_name = "endedAt")]
    pub ended_at: Option<f64>,
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
