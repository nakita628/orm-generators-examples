use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// A project's ISO week, folded from its owners' events. The id is the fold's aggregate key.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DigestProjectWeek")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "isoWeek", column_type = "Text")]
    pub iso_week: String,
    /// The fold state: counts, sums, a fixed-bucket latency histogram, capped model and evaluator maps.
    #[sea_orm(column_type = "JsonBinary")]
    pub state: Json,
    #[sea_orm(column_name = "occurredAt", column_type = "custom(\"timestamp(3)\")")]
    pub occurred_at: DateTime,
    #[sea_orm(column_name = "lastEventId", column_type = "Text")]
    pub last_event_id: String,
    #[sea_orm(column_name = "acceptedAt", column_type = "custom(\"timestamp(3)\")")]
    pub accepted_at: DateTime,
    #[sea_orm(column_name = "projectionVersion", column_type = "Text")]
    pub projection_version: String,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
