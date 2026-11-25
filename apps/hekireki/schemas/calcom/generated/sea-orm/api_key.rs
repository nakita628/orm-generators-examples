use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ApiKey")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userId")]
    pub user_id: i32,
    #[sea_orm(column_name = "teamId")]
    pub team_id: Option<i32>,
    #[sea_orm(column_type = "Text")]
    pub note: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(column_name = "lastUsedAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_used_at: Option<DateTime>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "hashedKey", column_type = "Text")]
    pub hashed_key: String,
    #[sea_orm(column_name = "appId", column_type = "Text")]
    pub app_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "ApiKey_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "ApiKey_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::app::Entity",
        from = "Column::AppId",
        to = "super::app::Column::Slug",
        on_delete = "Cascade",
        fk_name = "ApiKey_appId_fkey"
    )]
    App,
    #[sea_orm(has_many = "super::rate_limit::Entity")]
    RateLimits,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::app::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::App.def()
    }
}

impl Related<super::rate_limit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RateLimits.def()
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
