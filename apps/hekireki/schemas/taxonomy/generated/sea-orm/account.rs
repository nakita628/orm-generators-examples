use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "accounts")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "String(StringLen::N(191))")]
    pub user_id: String,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub r#type: String,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub provider: String,
    #[sea_orm(
        column_name = "providerAccountId",
        column_type = "String(StringLen::N(191))"
    )]
    pub provider_account_id: String,
    #[sea_orm(column_type = "Text")]
    pub refresh_token: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub access_token: Option<String>,
    pub expires_at: Option<i32>,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub token_type: Option<String>,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub scope: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub id_token: Option<String>,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub session_state: Option<String>,
    #[sea_orm(
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "accounts_userId_fkey"
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
        if insert && self.updated_at.is_not_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
