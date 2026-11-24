use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Used to generate PersonalAccessTokens, they're one-time use
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AuthorizationCode")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(unique)]
    pub code: String,
    #[sea_orm(column_name = "personalAccessTokenId")]
    pub personal_access_token_id: Option<String>,
    /// Client IP of the request that minted the code, compared with the approving browser's IP. Deleted with the code once it expires.
    #[sea_orm(column_name = "requesterIp")]
    pub requester_ip: Option<String>,
    /// S256 PKCE challenge from the client that minted the code. Null for codes minted by CLIs that predate it.
    #[sea_orm(column_name = "codeChallenge")]
    pub code_challenge: Option<String>,
    /// Loopback port the approving browser is redirected to with the approval code. Null means the user pastes it.
    #[sea_orm(column_name = "callbackPort")]
    pub callback_port: Option<i32>,
    /// sha256 of the approval code, which only the approving browser ever sees.
    #[sea_orm(column_name = "approvalCodeHash")]
    pub approval_code_hash: Option<String>,
    #[sea_orm(column_name = "oauthClientId")]
    pub oauth_client_id: Option<String>,
    #[sea_orm(column_name = "redirectUri")]
    pub redirect_uri: Option<String>,
    pub scope: Vec<String>,
    pub resource: Option<String>,
    #[sea_orm(column_name = "approvedByUserId")]
    pub approved_by_user_id: Option<String>,
    /// The user who started the OAuth authorization. Only they can approve it.
    #[sea_orm(column_name = "requestedByUserId")]
    pub requested_by_user_id: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::personal_access_token::Entity",
        from = "Column::PersonalAccessTokenId",
        to = "super::personal_access_token::Column::Id",
        on_update = "Cascade",
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
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
