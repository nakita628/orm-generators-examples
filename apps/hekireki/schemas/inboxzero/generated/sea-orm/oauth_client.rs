use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OauthClient")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "clientId", column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_name = "clientSecret", column_type = "Text")]
    pub client_secret: Option<String>,
    #[sea_orm(column_name = "clientDiscoveryId", column_type = "Text")]
    pub client_discovery_id: Option<String>,
    #[sea_orm(default_value = false)]
    pub disabled: Option<bool>,
    #[sea_orm(column_name = "skipConsent")]
    pub skip_consent: Option<bool>,
    #[sea_orm(column_name = "enableEndSession")]
    pub enable_end_session: Option<bool>,
    #[sea_orm(column_name = "subjectType", column_type = "Text")]
    pub subject_type: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub scopes: Vec<String>,
    #[sea_orm(
        column_name = "clientCredentialsScopes",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub client_credentials_scopes: Vec<String>,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: Option<DateTime>,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub uri: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub icon: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub contacts: Vec<String>,
    #[sea_orm(column_type = "Text")]
    pub tos: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub policy: Option<String>,
    #[sea_orm(column_name = "softwareId", column_type = "Text")]
    pub software_id: Option<String>,
    #[sea_orm(column_name = "softwareVersion", column_type = "Text")]
    pub software_version: Option<String>,
    #[sea_orm(column_name = "softwareStatement", column_type = "Text")]
    pub software_statement: Option<String>,
    #[sea_orm(
        column_name = "redirectUris",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub redirect_uris: Vec<String>,
    #[sea_orm(
        column_name = "postLogoutRedirectUris",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub post_logout_redirect_uris: Vec<String>,
    #[sea_orm(column_name = "backchannelLogoutUri", column_type = "Text")]
    pub backchannel_logout_uri: Option<String>,
    #[sea_orm(column_name = "backchannelLogoutSessionRequired")]
    pub backchannel_logout_session_required: Option<bool>,
    #[sea_orm(column_name = "tokenEndpointAuthMethod", column_type = "Text")]
    pub token_endpoint_auth_method: Option<String>,
    #[sea_orm(column_name = "applicationType", column_type = "Text")]
    pub application_type: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub jwks: Option<String>,
    #[sea_orm(column_name = "jwksUri", column_type = "Text")]
    pub jwks_uri: Option<String>,
    #[sea_orm(
        column_name = "grantTypes",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub grant_types: Vec<String>,
    #[sea_orm(
        column_name = "responseTypes",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub response_types: Vec<String>,
    #[sea_orm(column_name = "requirePKCE")]
    pub require_pkce: Option<bool>,
    #[sea_orm(column_name = "dpopBoundAccessTokens", default_value = false)]
    pub dpop_bound_access_tokens: Option<bool>,
    #[sea_orm(column_name = "referenceId", column_type = "Text")]
    pub reference_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "OauthClient_userId_fkey"
    )]
    User,
    #[sea_orm(has_many = "super::oauth_client_resource::Entity")]
    OauthClientResources,
    #[sea_orm(has_many = "super::oauth_refresh_token::Entity")]
    Oauthrefreshtokens,
    #[sea_orm(has_many = "super::oauth_access_token::Entity")]
    Oauthaccesstokens,
    #[sea_orm(has_many = "super::oauth_consent::Entity")]
    Oauthconsents,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::oauth_client_resource::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OauthClientResources.def()
    }
}

impl Related<super::oauth_refresh_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Oauthrefreshtokens.def()
    }
}

impl Related<super::oauth_access_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Oauthaccesstokens.def()
    }
}

impl Related<super::oauth_consent::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Oauthconsents.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
