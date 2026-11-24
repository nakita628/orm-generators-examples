use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OAuthRefreshToken")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "hashedToken")]
    pub hashed_token: String,
    #[sea_orm(column_name = "personalAccessTokenId")]
    pub personal_access_token_id: String,
    #[sea_orm(column_name = "oauthClientId")]
    pub oauth_client_id: String,
    pub scope: Vec<String>,
    pub resource: String,
    #[sea_orm(column_name = "usedAt")]
    pub used_at: Option<DateTime>,
    #[sea_orm(column_name = "replacedById")]
    pub replaced_by_id: Option<String>,
    #[sea_orm(column_name = "idleExpiresAt")]
    pub idle_expires_at: DateTime,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::personal_access_token::Entity",
        from = "Column::PersonalAccessTokenId",
        to = "super::personal_access_token::Column::Id",
        on_delete = "Cascade"
    )]
    PersonalAccessToken,
    #[sea_orm(
        belongs_to = "super::oauth_client::Entity",
        from = "Column::OauthClientId",
        to = "super::oauth_client::Column::Id",
        on_delete = "Cascade"
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
