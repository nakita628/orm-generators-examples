use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OAuthClient")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "clientId", column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_name = "clientName", column_type = "Text")]
    pub client_name: String,
    #[sea_orm(column_name = "clientUri", column_type = "Text")]
    pub client_uri: Option<String>,
    #[sea_orm(column_name = "logoUri", column_type = "Text")]
    pub logo_uri: Option<String>,
    #[sea_orm(
        column_name = "redirectUris",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub redirect_uris: Vec<String>,
    #[sea_orm(
        column_name = "grantTypes",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub grant_types: Vec<String>,
    #[sea_orm(column_type = "Text")]
    pub scopes: String,
    #[sea_orm(column_name = "tokenAuthMethod", column_type = "Text")]
    pub token_auth_method: String,
    #[sea_orm(column_name = "hashedSecret", column_type = "Text")]
    pub hashed_secret: Option<String>,
    #[sea_orm(column_name = "registrationType", column_type = "Text")]
    pub registration_type: String,
    #[sea_orm(column_name = "metadataUrl", column_type = "Text")]
    pub metadata_url: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::restricted_token::Entity")]
    Tokens,
}

impl Related<super::restricted_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Tokens.def()
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
