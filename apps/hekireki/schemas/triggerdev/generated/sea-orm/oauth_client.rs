use super::oauth_client_kind::OAuthClientKind;
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
    pub kind: OAuthClientKind,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "clientId", column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_name = "metadataUrl", column_type = "Text")]
    pub metadata_url: Option<String>,
    #[sea_orm(column_name = "displayName", column_type = "Text")]
    pub display_name: Option<String>,
    #[sea_orm(column_name = "selfReportedName", column_type = "Text")]
    pub self_reported_name: Option<String>,
    #[sea_orm(
        column_name = "redirectUris",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub redirect_uris: Vec<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "dedupeKey", column_type = "Text")]
    pub dedupe_key: Option<String>,
    #[sea_orm(
        column_name = "metadataExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub metadata_expires_at: Option<DateTime>,
    #[sea_orm(column_name = "lastUsedAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_used_at: Option<DateTime>,
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
