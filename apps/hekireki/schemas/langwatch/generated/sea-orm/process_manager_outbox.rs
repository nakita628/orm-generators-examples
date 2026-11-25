use super::process_manager_outbox_status::ProcessManagerOutboxStatus;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Transactional intents emitted by a durable process transition.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ProcessManagerOutbox")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "processName", column_type = "Text")]
    pub process_name: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "processKey", column_type = "Text")]
    pub process_key: String,
    #[sea_orm(column_name = "tenantId", column_type = "Text")]
    pub tenant_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "messageKey", column_type = "Text")]
    pub message_key: String,
    #[sea_orm(column_name = "intentType", column_type = "Text")]
    pub intent_type: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub payload: Json,
    #[sea_orm(column_name = "traceCarrier", column_type = "JsonBinary")]
    pub trace_carrier: Json,
    #[sea_orm(column_name = "sourceEventId", column_type = "Text")]
    pub source_event_id: Option<String>,
    #[sea_orm(default_value = "pending")]
    pub status: ProcessManagerOutboxStatus,
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    #[sea_orm(
        column_name = "nextAttemptAt",
        column_type = "custom(\"timestamptz(3)\")"
    )]
    pub next_attempt_at: DateTimeWithTimeZone,
    #[sea_orm(
        column_name = "leasedUntil",
        column_type = "custom(\"timestamptz(3)\")"
    )]
    pub leased_until: Option<DateTimeWithTimeZone>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "leaseToken", column_type = "Text")]
    pub lease_token: Option<String>,
    #[sea_orm(
        column_name = "dispatchedAt",
        column_type = "custom(\"timestamptz(3)\")"
    )]
    pub dispatched_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamptz(3)\")")]
    pub created_at: DateTimeWithTimeZone,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamptz(3)\")")]
    pub updated_at: DateTimeWithTimeZone,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::process_manager_outbox_attempt::Entity")]
    AttemptLog,
}

impl Related<super::process_manager_outbox_attempt::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AttemptLog.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
