use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "scimManagedConnectionEvent")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "connectionRecordId", column_type = "Text")]
    pub connection_record_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "eventKey", column_type = "Text")]
    pub event_key: String,
    pub sequence: i32,
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(column_name = "actorId", column_type = "Text")]
    pub actor_id: String,
    #[sea_orm(column_name = "credentialId", column_type = "Text")]
    pub credential_id: Option<String>,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::scim_managed_connection::Entity",
        from = "Column::ConnectionRecordId",
        to = "super::scim_managed_connection::Column::Id",
        on_delete = "Cascade",
        fk_name = "scimManagedConnectionEvent_connectionRecordId_fkey"
    )]
    Scimmanagedconnection,
}

impl Related<super::scim_managed_connection::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scimmanagedconnection.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
