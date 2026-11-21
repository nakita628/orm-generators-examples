use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Stores third-party authentication account information.
/// Enables OAuth and other external authentication methods.
///
/// @property id - Unique identifier for the account
/// @property user - The Formbricks user who owns this account
/// @property provider - The authentication provider (GitHub, Google, etc.)
/// @property providerAccountId - User ID from the provider
/// @property access_token - OAuth access token
/// @property refresh_token - OAuth refresh token
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Account")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created_at: DateTime,
    pub updated_at: DateTime,
    #[sea_orm(column_name = "userId")]
    pub user_id: String,
    pub r#type: Option<String>,
    pub provider: String,
    #[sea_orm(column_name = "providerAccountId")]
    pub provider_account_id: String,
    #[sea_orm(column_type = "Text")]
    pub access_token: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub refresh_token: Option<String>,
    pub expires_at: Option<i32>,
    pub ext_expires_in: Option<i32>,
    pub token_type: Option<String>,
    pub scope: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub id_token: Option<String>,
    pub session_state: Option<String>,
    pub password: Option<String>,
    #[sea_orm(column_name = "accessTokenExpiresAt")]
    pub access_token_expires_at: Option<DateTime>,
    #[sea_orm(column_name = "refreshTokenExpiresAt")]
    pub refresh_token_expires_at: Option<DateTime>,
    /// Better Auth 1.7.0 (ENG-2343): the account's identity namespace. 1.7.0 keyed accounts on
    /// (issuer, accountId) and filtered every lookup on it — findAccountByKey, findAccountOwnerByKey,
    /// findCredentialAccount and updatePassword — so on 1.7.0 a null here locks the user out.
    /// 1.7.3 dropped the field and 1.7.7 keys accounts on (provider, providerAccountId) again
    /// (ENG-3880). We keep writing it (see `account.create.before` in better-auth-hooks.ts) so a
    /// rollback to 1.7.0 still finds every row; drop it once that rollback is off the table.
    ///
    /// NULLABLE on purpose. A NOT NULL column would fail every INSERT from a 1.6 pod still serving
    /// during a rolling deploy. The migration backfills every existing row; see the migration SQL for
    /// the value derivation.
    pub issuer: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
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
