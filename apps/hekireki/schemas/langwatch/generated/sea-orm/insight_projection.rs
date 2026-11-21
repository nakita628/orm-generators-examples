use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// One insight in its owner's inbox in a project, folded from its `insight` aggregate
/// (modules/insight). A read model: rebuildable by replaying the event log. Epochs in
/// milliseconds.
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "InsightProjection")]
pub struct Model {
    /// The insight id, which is also its aggregate id.
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "projectId")]
    pub project_id: String,
    pub title: String,
    pub body: String,
    /// bad | watch | good
    pub tone: String,
    pub topic: Option<String>,
    #[sea_orm(column_name = "validDays")]
    pub valid_days: i32,
    pub lwql: Option<String>,
    #[sea_orm(column_name = "sourceConversationId")]
    pub source_conversation_id: Option<String>,
    #[sea_orm(column_name = "sourceMessageId")]
    pub source_message_id: Option<String>,
    #[sea_orm(column_name = "filedByUserId")]
    pub filed_by_user_id: Option<String>,
    #[sea_orm(column_name = "filedAt")]
    pub filed_at: f64,
    /// Set by a later run that found the insight still true.
    #[sea_orm(column_name = "renewedAt")]
    pub renewed_at: Option<f64>,
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
    /// The board and widget it came from. A pointer only, never a relation: the row reads the
    /// same once either is deleted, so the names are kept as filed.
    #[sea_orm(column_name = "boardId")]
    pub board_id: Option<String>,
    #[sea_orm(column_name = "boardName")]
    pub board_name: Option<String>,
    #[sea_orm(column_name = "widgetId")]
    pub widget_id: Option<String>,
    #[sea_orm(column_name = "widgetName")]
    pub widget_name: Option<String>,
    /// chat | run
    #[sea_orm(column_name = "filedVia", default_value = "chat")]
    pub filed_via: String,
    /// The fixed window `lwql` replays over (epoch milliseconds, half-open) and its step.
    #[sea_orm(column_name = "replayStart")]
    pub replay_start: Option<f64>,
    #[sea_orm(column_name = "replayEnd")]
    pub replay_end: Option<f64>,
    #[sea_orm(column_name = "replayGranularitySeconds")]
    pub replay_granularity_seconds: Option<i32>,
    /// The values in force when it was filed: `{ period, parameters }`.
    #[sea_orm(column_name = "replayContext")]
    pub replay_context: Option<Json>,
    /// Whose insight it is: the one person who reads it. NULL on a row folded before the
    /// column existed, which is owned by whoever filed it (`filedByUserId`).
    #[sea_orm(column_name = "ownerUserId")]
    pub owner_user_id: Option<String>,
    /// The stable key a daily run matches a repeated finding by, so it renews the insight
    /// and files no twin. NULL for an insight saved from a chat and on a row folded before
    /// the column existed.
    #[sea_orm(column_name = "findingKey")]
    pub finding_key: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
