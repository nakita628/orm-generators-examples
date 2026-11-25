use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OidcProvider")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// @zod.custom(z.coerce.date())
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    /// @zod.custom(z.coerce.date())
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub issuer: String,
    #[sea_orm(column_name = "clientId", column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_name = "clientSecret", column_type = "Text")]
    pub client_secret: String,
    #[sea_orm(column_name = "authorizationEndpoint", column_type = "Text")]
    pub authorization_endpoint: Option<String>,
    #[sea_orm(column_name = "tokenEndpoint", column_type = "Text")]
    pub token_endpoint: Option<String>,
    #[sea_orm(column_name = "userinfoEndpoint", column_type = "Text")]
    pub userinfo_endpoint: Option<String>,
    #[sea_orm(column_name = "jwksUri", column_type = "Text")]
    pub jwks_uri: Option<String>,
    #[sea_orm(column_name = "introspectionEndpoint", column_type = "Text")]
    pub introspection_endpoint: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub scopes: Vec<String>,
    #[sea_orm(column_type = "Text")]
    pub audience: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub prompt: Option<String>,
    #[sea_orm(
        column_name = "usernameClaim",
        column_type = "Text",
        default_value = "email"
    )]
    pub username_claim: String,
    #[sea_orm(
        column_name = "emailClaim",
        column_type = "Text",
        default_value = "email"
    )]
    pub email_claim: String,
    #[sea_orm(
        column_name = "nameClaim",
        column_type = "Text",
        default_value = "name"
    )]
    pub name_claim: String,
    #[sea_orm(
        column_name = "groupClaim",
        column_type = "Text",
        default_value = "groups"
    )]
    pub group_claim: String,
    #[sea_orm(default_value = true)]
    pub enabled: bool,
    #[sea_orm(column_name = "autoDiscovery", default_value = true)]
    pub auto_discovery: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::oidc_login_group::Entity")]
    OidcLoginGroups,
}

impl Related<super::oidc_login_group::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OidcLoginGroups.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
