use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// The registry's append-only change feed: one row per registered, versioned, removed or
/// tool-deprecated fact. `sequence` is the pull cursor; `deliveryRecordedAt` marks the webhook
/// fan-out as recorded in governance's outbox.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AgentRegistryEvent")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    pub sequence: i64,
    pub r#type: String,
    #[sea_orm(column_name = "entryId")]
    pub entry_id: String,
    #[sea_orm(column_name = "fromVersion")]
    pub from_version: Option<i32>,
    #[sea_orm(column_name = "toVersion")]
    pub to_version: Option<i32>,
    pub payload: Json,
    #[sea_orm(column_name = "occurredAt")]
    pub occurred_at: DateTime,
    #[sea_orm(column_name = "deliveryRecordedAt")]
    pub delivery_recorded_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
