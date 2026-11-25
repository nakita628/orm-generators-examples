use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Whether a target was already stopped before a stop first touched it, recorded once before the
/// stop's platform call, so a retried call after a success still knows the change is the stop's own.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AgentStopTargetPriorState")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "stopId", column_type = "Text")]
    pub stop_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "targetId", column_type = "Text")]
    pub target_id: String,
    #[sea_orm(column_name = "wasStopped")]
    pub was_stopped: bool,
    #[sea_orm(column_name = "recordedAt", column_type = "custom(\"timestamp(3)\")")]
    pub recorded_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
