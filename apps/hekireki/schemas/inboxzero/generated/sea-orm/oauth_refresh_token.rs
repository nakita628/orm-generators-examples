use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OauthRefreshToken")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub token: String,
    #[sea_orm(column_name = "clientId", column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_name = "sessionId", column_type = "Text")]
    pub session_id: Option<String>,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "referenceId", column_type = "Text")]
    pub reference_id: Option<String>,
    #[sea_orm(column_name = "authorizationCodeId", column_type = "Text")]
    pub authorization_code_id: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub resources: Vec<String>,
    #[sea_orm(
        column_name = "requestedUserInfoClaims",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub requested_user_info_claims: Vec<String>,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: DateTime,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub revoked: Option<DateTime>,
    #[sea_orm(column_name = "rotatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub rotated_at: Option<DateTime>,
    #[sea_orm(column_name = "rotationReplayResponse", column_type = "Text")]
    pub rotation_replay_response: Option<String>,
    #[sea_orm(
        column_name = "rotationReplayExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub rotation_replay_expires_at: Option<DateTime>,
    #[sea_orm(column_name = "authTime", column_type = "custom(\"timestamp(3)\")")]
    pub auth_time: Option<DateTime>,
    #[sea_orm(column_type = "JsonBinary")]
    pub confirmation: Option<Json>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub scopes: Vec<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::oauth_client::Entity",
        from = "Column::ClientId",
        to = "super::oauth_client::Column::ClientId",
        on_delete = "Cascade",
        fk_name = "OauthRefreshToken_clientId_fkey"
    )]
    Client,
    #[sea_orm(
        belongs_to = "super::session::Entity",
        from = "Column::SessionId",
        to = "super::session::Column::Id",
        on_delete = "SetNull",
        fk_name = "OauthRefreshToken_sessionId_fkey"
    )]
    Session,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "OauthRefreshToken_userId_fkey"
    )]
    User,
    #[sea_orm(has_many = "super::oauth_access_token::Entity")]
    Oauthaccesstokens,
}

impl Related<super::oauth_client::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Client.def()
    }
}

impl Related<super::session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Session.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::oauth_access_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Oauthaccesstokens.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
