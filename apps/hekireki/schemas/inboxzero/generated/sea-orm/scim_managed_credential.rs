use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "scimManagedCredential")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "connectionRecordId", column_type = "Text")]
    pub connection_record_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "credentialId", column_type = "Text")]
    pub credential_id: String,
    #[sea_orm(column_name = "tokenDigest", column_type = "Text")]
    pub token_digest: String,
    #[sea_orm(column_name = "hashVersion", column_type = "Text")]
    pub hash_version: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "activeSlotKey", column_type = "Text")]
    pub active_slot_key: String,
    #[sea_orm(column_type = "Text")]
    pub status: String,
    #[sea_orm(column_name = "serializedScopes", column_type = "Text")]
    pub serialized_scopes: String,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: DateTime,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "createdBy", column_type = "Text")]
    pub created_by: String,
    #[sea_orm(column_name = "lastUsedAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_used_at: Option<DateTime>,
    #[sea_orm(column_name = "revokedAt", column_type = "custom(\"timestamp(3)\")")]
    pub revoked_at: Option<DateTime>,
    #[sea_orm(column_name = "revokedBy", column_type = "Text")]
    pub revoked_by: Option<String>,
    #[sea_orm(
        column_name = "decommissionedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub decommissioned_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::scim_managed_connection::Entity",
        from = "Column::ConnectionRecordId",
        to = "super::scim_managed_connection::Column::Id",
        on_delete = "Cascade",
        fk_name = "scimManagedCredential_connectionRecordId_fkey"
    )]
    Scimmanagedconnection,
}

impl Related<super::scim_managed_connection::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scimmanagedconnection.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
