use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OauthClientResource")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "clientId", column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_name = "resourceId", column_type = "Text")]
    pub resource_id: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::oauth_client::Entity",
        from = "Column::ClientId",
        to = "super::oauth_client::Column::ClientId",
        on_delete = "Cascade",
        fk_name = "OauthClientResource_clientId_fkey"
    )]
    Client,
    #[sea_orm(
        belongs_to = "super::oauth_resource::Entity",
        from = "Column::ResourceId",
        to = "super::oauth_resource::Column::Identifier",
        on_delete = "Cascade",
        fk_name = "OauthClientResource_resourceId_fkey"
    )]
    OauthResource,
}

impl Related<super::oauth_client::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Client.def()
    }
}

impl Related<super::oauth_resource::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OauthResource.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
