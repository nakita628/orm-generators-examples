use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OAuthAccessToken")]
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
    #[sea_orm(column_name = "lastUsed", column_type = "custom(\"timestamp(3)\")")]
    pub last_used: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub hash: String,
    /// @zod.custom(z.coerce.date())
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: DateTime,
    #[sea_orm(column_name = "refreshTokenId", column_type = "Text")]
    pub refresh_token_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user_api_token::Entity",
        from = "Column::RefreshTokenId",
        to = "super::user_api_token::Column::Id",
        fk_name = "OAuthAccessToken_refreshTokenId_fkey"
    )]
    RefreshToken,
}

impl Related<super::user_api_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RefreshToken.def()
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
