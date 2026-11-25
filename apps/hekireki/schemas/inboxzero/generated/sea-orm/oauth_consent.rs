use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OauthConsent")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "clientId", column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "referenceId", column_type = "Text")]
    pub reference_id: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub resources: Vec<String>,
    #[sea_orm(
        column_name = "requestedUserInfoClaims",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub requested_user_info_claims: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub scopes: Vec<String>,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::oauth_client::Entity",
        from = "Column::ClientId",
        to = "super::oauth_client::Column::ClientId",
        on_delete = "Cascade",
        fk_name = "OauthConsent_clientId_fkey"
    )]
    Client,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "OauthConsent_userId_fkey"
    )]
    User,
}

impl Related<super::oauth_client::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Client.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
