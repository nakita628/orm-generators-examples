use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Account")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_type = "Text")]
    pub provider: String,
    #[sea_orm(column_name = "providerAccountId", column_type = "Text")]
    pub provider_account_id: String,
    #[sea_orm(column_name = "providerRefreshToken", column_type = "Text")]
    pub provider_refresh_token: Option<String>,
    #[sea_orm(column_name = "providerAccessToken", column_type = "Text")]
    pub provider_access_token: Option<String>,
    #[sea_orm(column_name = "providerScope", column_type = "Text")]
    pub provider_scope: Option<String>,
    #[sea_orm(
        column_name = "loggedIn",
        column_type = "custom(\"timestamptz(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub logged_in: DateTimeWithTimeZone,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Uid",
        on_delete = "Cascade",
        fk_name = "Account_userId_fkey"
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
        if insert && self.logged_in.is_not_set() {
            self.logged_in = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
