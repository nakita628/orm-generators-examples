use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "scimManagedConnection")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "creationRequestId", column_type = "Text")]
    pub creation_request_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "connectionId", column_type = "Text")]
    pub connection_id: String,
    #[sea_orm(column_name = "provisioningDomainId", column_type = "Text")]
    pub provisioning_domain_id: String,
    #[sea_orm(column_type = "Text")]
    pub status: String,
    pub revision: i32,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "createdBy", column_type = "Text")]
    pub created_by: String,
    #[sea_orm(
        column_name = "decommissionStartedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub decommission_started_at: Option<DateTime>,
    #[sea_orm(column_name = "decommissionStartedBy", column_type = "Text")]
    pub decommission_started_by: Option<String>,
    #[sea_orm(
        column_name = "decommissionedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub decommissioned_at: Option<DateTime>,
    #[sea_orm(column_name = "decommissionedBy", column_type = "Text")]
    pub decommissioned_by: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::scim_managed_credential::Entity")]
    Scimmanagedcredentials,
    #[sea_orm(has_many = "super::scim_managed_connection_event::Entity")]
    Scimmanagedconnectionevents,
}

impl Related<super::scim_managed_credential::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scimmanagedcredentials.def()
    }
}

impl Related<super::scim_managed_connection_event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scimmanagedconnectionevents.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
