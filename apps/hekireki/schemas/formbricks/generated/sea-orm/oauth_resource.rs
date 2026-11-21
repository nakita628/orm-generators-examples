use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
/// its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
/// a token is now bound to the resource it was approved for instead of to a static allow-list.
/// `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "oauthResource")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(unique)]
    pub identifier: String,
    pub name: String,
    #[sea_orm(column_name = "accessTokenTtl")]
    pub access_token_ttl: Option<i32>,
    #[sea_orm(column_name = "refreshTokenTtl")]
    pub refresh_token_ttl: Option<i32>,
    #[sea_orm(column_name = "signingAlgorithm")]
    pub signing_algorithm: Option<String>,
    #[sea_orm(column_name = "signingKeyId")]
    pub signing_key_id: Option<String>,
    #[sea_orm(column_name = "allowedScopes")]
    pub allowed_scopes: Vec<String>,
    #[sea_orm(column_name = "customClaims")]
    pub custom_claims: Option<Json>,
    #[sea_orm(column_name = "dpopBoundAccessTokensRequired", default_value = false)]
    pub dpop_bound_access_tokens_required: Option<bool>,
    #[sea_orm(default_value = false)]
    pub disabled: Option<bool>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: Option<DateTime>,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: Option<DateTime>,
    #[sea_orm(column_name = "policyVersion", default_value = 1)]
    pub policy_version: Option<i32>,
    pub metadata: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::oauth_client_resource::Entity")]
    Clients,
}

impl Related<super::oauth_client_resource::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Clients.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
