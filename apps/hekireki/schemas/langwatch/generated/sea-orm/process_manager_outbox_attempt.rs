use super::process_manager_outbox_attempt_outcome::ProcessManagerOutboxAttemptOutcome;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Why an outbox delivery failed, one row per FAILED attempt
/// (specs/ops/dead-letter-recovery.feature). Successes write nothing, so the
/// table grows with trouble, not with traffic. Rows die with their message:
/// the cascade delete is the retention story, riding the existing sweep.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ProcessManagerOutboxAttempt")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "outboxId", column_type = "Text")]
    pub outbox_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    /// 1-based attempt number, as the dispatcher counts it.
    pub attempt: i32,
    #[sea_orm(column_name = "occurredAt", column_type = "custom(\"timestamptz(3)\")")]
    pub occurred_at: DateTimeWithTimeZone,
    pub outcome: ProcessManagerOutboxAttemptOutcome,
    #[sea_orm(column_name = "errorType", column_type = "Text")]
    pub error_type: String,
    /// The safe failure diagnostic the dispatcher already computes for its
    /// span — never a raw provider body.
    #[sea_orm(column_name = "errorMessage", column_type = "Text")]
    pub error_message: String,
    #[sea_orm(column_name = "retryAfterMs")]
    pub retry_after_ms: Option<i32>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::process_manager_outbox::Entity",
        from = "Column::OutboxId",
        to = "super::process_manager_outbox::Column::Id",
        on_delete = "Cascade",
        fk_name = "ProcessManagerOutboxAttempt_outboxId_fkey"
    )]
    Outbox,
}

impl Related<super::process_manager_outbox::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Outbox.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
