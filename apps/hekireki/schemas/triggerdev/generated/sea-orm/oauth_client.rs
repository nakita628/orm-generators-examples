use super::oauth_client_kind::OAuthClientKind;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OAuthClient")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub kind: OAuthClientKind,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "clientId")]
    pub client_id: String,
    #[sea_orm(column_name = "metadataUrl")]
    pub metadata_url: Option<String>,
    #[sea_orm(column_name = "displayName")]
    pub display_name: Option<String>,
    #[sea_orm(column_name = "selfReportedName")]
    pub self_reported_name: Option<String>,
    #[sea_orm(column_name = "redirectUris")]
    pub redirect_uris: Vec<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "dedupeKey")]
    pub dedupe_key: Option<String>,
    #[sea_orm(column_name = "metadataExpiresAt")]
    pub metadata_expires_at: Option<DateTime>,
    #[sea_orm(column_name = "lastUsedAt")]
    pub last_used_at: Option<DateTime>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::authorization_code::Entity")]
    AuthorizationCodes,
    #[sea_orm(has_many = "super::personal_access_token::Entity")]
    PersonalAccessTokens,
    #[sea_orm(has_many = "super::oauth_refresh_token::Entity")]
    RefreshTokens,
}

impl Related<super::authorization_code::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuthorizationCodes.def()
    }
}

impl Related<super::personal_access_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PersonalAccessTokens.def()
    }
}

impl Related<super::oauth_refresh_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RefreshTokens.def()
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
