use super::retention_policy_scope_type::RetentionPolicyScopeType;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RetentionApplyStatus")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "scopeType")]
    pub scope_type: RetentionPolicyScopeType,
    #[sea_orm(column_name = "scopeId")]
    pub scope_id: String,
    #[sea_orm(column_name = "changedAt")]
    pub changed_at: Option<f64>,
    #[sea_orm(column_name = "appliedAt")]
    pub applied_at: Option<f64>,
    #[sea_orm(column_name = "raisedAt")]
    pub raised_at: Option<f64>,
    #[sea_orm(column_name = "shortenedAt")]
    pub shortened_at: Option<f64>,
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
