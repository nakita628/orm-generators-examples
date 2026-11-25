use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "scimConnectionBinding")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "connectionId", column_type = "Text")]
    pub connection_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "connectionKey", column_type = "Text")]
    pub connection_key: String,
    #[sea_orm(column_name = "provisioningDomainId", column_type = "Text")]
    pub provisioning_domain_id: String,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "decommissionedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub decommissioned_at: Option<DateTime>,
    #[sea_orm(
        column_name = "decommissionStatus",
        column_type = "Text",
        default_value = "active"
    )]
    pub decommission_status: String,
    #[sea_orm(column_name = "decommissionCursorUserId", column_type = "Text")]
    pub decommission_cursor_user_id: Option<String>,
    #[sea_orm(column_name = "decommissionReconciledUserCount", default_value = 0)]
    pub decommission_reconciled_user_count: i32,
    #[sea_orm(column_name = "decommissionBatchCount", default_value = 0)]
    pub decommission_batch_count: i32,
    #[sea_orm(column_name = "decommissionRevision", default_value = 0)]
    pub decommission_revision: i32,
    #[sea_orm(
        column_name = "decommissionCompletedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub decommission_completed_at: Option<DateTime>,
    #[sea_orm(column_name = "decommissionLeaseId", column_type = "Text")]
    pub decommission_lease_id: Option<String>,
    #[sea_orm(
        column_name = "decommissionLeaseExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub decommission_lease_expires_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
