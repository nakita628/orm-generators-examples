use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Rebuildable operational message fact. One source event produces one row.
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LangyMessageProjection")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "conversationId", column_type = "Text")]
    pub conversation_id: String,
    #[sea_orm(column_name = "messageId", column_type = "Text")]
    pub message_id: String,
    #[sea_orm(column_type = "Text")]
    pub role: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub parts: Json,
    #[sea_orm(column_name = "sourceEventId", column_type = "Text")]
    pub source_event_id: String,
    #[sea_orm(column_name = "occurredAt")]
    pub occurred_at: f64,
    #[sea_orm(column_name = "acceptedAt")]
    pub accepted_at: f64,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: f64,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: f64,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
