use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "scimProjectionGrant")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "connectionId", column_type = "Text")]
    pub connection_id: String,
    #[sea_orm(column_name = "provisioningDomainId", column_type = "Text")]
    pub provisioning_domain_id: String,
    #[sea_orm(column_name = "scimUserId", column_type = "Text")]
    pub scim_user_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "sourceKind", column_type = "Text")]
    pub source_kind: String,
    #[sea_orm(column_name = "sourceId", column_type = "Text")]
    pub source_id: String,
    #[sea_orm(column_name = "sourceValue", column_type = "Text")]
    pub source_value: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub role: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "grantKey", column_type = "Text")]
    pub grant_key: String,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::scim_user::Entity",
        from = "Column::ScimUserId",
        to = "super::scim_user::Column::Id",
        on_delete = "Cascade",
        fk_name = "scimProjectionGrant_scimUserId_fkey"
    )]
    ScimUser,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "scimProjectionGrant_userId_fkey"
    )]
    User,
}

impl Related<super::scim_user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScimUser.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
