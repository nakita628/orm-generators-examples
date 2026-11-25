use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OAuthRefreshToken")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "hashedToken", column_type = "Text")]
    pub hashed_token: String,
    #[sea_orm(column_name = "personalAccessTokenId", column_type = "Text")]
    pub personal_access_token_id: String,
    #[sea_orm(column_name = "oauthClientId", column_type = "Text")]
    pub oauth_client_id: String,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub scope: Vec<String>,
    #[sea_orm(column_type = "Text")]
    pub resource: String,
    #[sea_orm(column_name = "usedAt", column_type = "custom(\"timestamp(3)\")")]
    pub used_at: Option<DateTime>,
    #[sea_orm(column_name = "replacedById", column_type = "Text")]
    pub replaced_by_id: Option<String>,
    #[sea_orm(
        column_name = "idleExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub idle_expires_at: DateTime,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::personal_access_token::Entity",
        from = "Column::PersonalAccessTokenId",
        to = "super::personal_access_token::Column::Id",
        on_delete = "Cascade",
        fk_name = "OAuthRefreshToken_personalAccessTokenId_fkey"
    )]
    PersonalAccessToken,
    #[sea_orm(
        belongs_to = "super::oauth_client::Entity",
        from = "Column::OauthClientId",
        to = "super::oauth_client::Column::Id",
        on_delete = "Cascade",
        fk_name = "OAuthRefreshToken_oauthClientId_fkey"
    )]
    OauthClient,
}

impl Related<super::personal_access_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PersonalAccessToken.def()
    }
}

impl Related<super::oauth_client::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OauthClient.def()
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
        Ok(self)
    }
}
