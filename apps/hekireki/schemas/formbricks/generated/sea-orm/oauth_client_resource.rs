use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Better Auth 1.7 (ENG-2343): links a client to the resources it may request tokens for. With
/// `enforcePerClientResources` (the default) a missing row means `invalid_target` at the token
/// endpoint, which is why the migration backfills one row per existing client.
/// Both sides reference natural keys, matching the other oauth child tables.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "oauthClientResource")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "clientId")]
    pub client_id: String,
    #[sea_orm(column_name = "resourceId")]
    pub resource_id: String,
    pub metadata: Option<Json>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::oauth_client::Entity",
        from = "Column::ClientId",
        to = "super::oauth_client::Column::ClientId",
        on_delete = "Cascade"
    )]
    Client,
    #[sea_orm(
        belongs_to = "super::oauth_resource::Entity",
        from = "Column::ResourceId",
        to = "super::oauth_resource::Column::Identifier",
        on_delete = "Cascade"
    )]
    Resource,
}

impl Related<super::oauth_client::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Client.def()
    }
}

impl Related<super::oauth_resource::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Resource.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
