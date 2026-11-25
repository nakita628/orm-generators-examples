use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RestrictedToken")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "hashedKey", column_type = "Text")]
    pub hashed_key: String,
    #[sea_orm(column_name = "partialKey", column_type = "Text")]
    pub partial_key: String,
    #[sea_orm(column_type = "Text")]
    pub scopes: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "live")]
    pub mode: String,
    #[sea_orm(column_type = "Text", default_value = "dashboard")]
    pub source: String,
    #[sea_orm(
        column_name = "subjectType",
        column_type = "Text",
        default_value = "user"
    )]
    pub subject_type: String,
    #[sea_orm(column_name = "clientId", column_type = "Text")]
    pub client_id: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub expires: Option<DateTime>,
    #[sea_orm(column_name = "lastUsed", column_type = "custom(\"timestamp(3)\")")]
    pub last_used: Option<DateTime>,
    #[sea_orm(column_name = "rateLimit", default_value = 60)]
    pub rate_limit: i32,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "teamId", column_type = "Text")]
    pub team_id: String,
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
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "RestrictedToken_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "RestrictedToken_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::oauth_client::Entity",
        from = "Column::ClientId",
        to = "super::oauth_client::Column::Id",
        on_delete = "SetNull",
        fk_name = "RestrictedToken_clientId_fkey"
    )]
    OauthClient,
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
