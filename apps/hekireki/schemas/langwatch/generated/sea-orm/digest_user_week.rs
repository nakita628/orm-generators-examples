use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// A person's ISO week of coding-agent activity in one organization, folded from the agents' facts.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DigestUserWeek")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "isoWeek", column_type = "Text")]
    pub iso_week: String,
    /// Sessions (a capped set of ids) and the agents' metric series, converged last-write-wins.
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
