use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "scimUser")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "connectionId", column_type = "Text")]
    pub connection_id: String,
    #[sea_orm(column_name = "provisioningDomainId", column_type = "Text")]
    pub provisioning_domain_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "connectionUserKey", column_type = "Text")]
    pub connection_user_key: String,
    #[sea_orm(column_name = "userName", column_type = "Text")]
    pub user_name: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "userNameKey", column_type = "Text")]
    pub user_name_key: String,
    #[sea_orm(column_name = "primaryEmail", column_type = "Text")]
    pub primary_email: String,
    #[sea_orm(column_name = "workEmailValueIndex", column_type = "Text")]
    pub work_email_value_index: String,
    #[sea_orm(column_name = "emailValueIndex", column_type = "Text")]
    pub email_value_index: String,
    #[sea_orm(column_name = "displayName", column_type = "Text")]
    pub display_name: String,
    #[sea_orm(column_name = "formattedName", column_type = "Text")]
    pub formatted_name: String,
    #[sea_orm(column_name = "givenName", column_type = "Text")]
    pub given_name: Option<String>,
    #[sea_orm(column_name = "familyName", column_type = "Text")]
    pub family_name: Option<String>,
    #[sea_orm(column_name = "serializedEmails", column_type = "Text")]
    pub serialized_emails: String,
    #[sea_orm(column_name = "serializedAttributes", column_type = "Text")]
    pub serialized_attributes: Option<String>,
    #[sea_orm(column_name = "externalId", column_type = "Text")]
    pub external_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "externalIdKey", column_type = "Text")]
    pub external_id_key: Option<String>,
    pub active: bool,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "orderKey", column_type = "Text")]
    pub order_key: String,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "scimUser_userId_fkey"
    )]
    User,
    #[sea_orm(has_many = "super::scim_projection_grant::Entity")]
    ScimProjectionGrants,
    #[sea_orm(has_many = "super::scim_group_member::Entity")]
    ScimGroupMembers,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::scim_projection_grant::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScimProjectionGrants.def()
    }
}

impl Related<super::scim_group_member::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScimGroupMembers.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
