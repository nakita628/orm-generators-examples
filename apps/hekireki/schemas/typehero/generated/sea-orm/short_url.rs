use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ShortURL")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    /// the original full url that we want to shorten e.g. https://www.url.com
    #[sea_orm(column_name = "originalUrl", column_type = "custom(\"mediumtext\")")]
    pub original_url: String,
    /// slug for the short url e.g. url.com/[shortUrlSlug]
    #[sea_orm(unique)]
    #[sea_orm(
        column_name = "shortUrlSlug",
        column_type = "String(StringLen::N(191))"
    )]
    pub short_url_slug: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub created_at: DateTime,
    /// the time the short url will expire
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"datetime(3)\")")]
    pub expires_at: Option<DateTime>,
    /// the number of times the short url has redirected to the original url
    #[sea_orm(default_value = 0)]
    pub clicks: i32,
    #[sea_orm(column_name = "userId", column_type = "String(StringLen::N(191))")]
    pub user_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        fk_name = "ShortURL_userId_fkey"
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
        Ok(self)
    }
}
